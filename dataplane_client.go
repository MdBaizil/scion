// Code references : https://github.com/netsec-ethz/scion-homeworks/blob/master/latency/timestamp_client.go and reference https://github.com/perrig/scionlab/blob/master/sensorapp/sensorfetcher/sensorfetcher.go

package main



import (

	"flag"

	"fmt"  //importing fmt package for printing

	"log"

	"math/rand" //importing for mathematical operations

	"time"

  "encoding/binary"

	"github.com/scionproto/scion/go/lib/snet" //importing snet packages for the scion connections

	"github.com/scionproto/scion/go/lib/sciond"

)

func logerror(e error) {   //Error function



if e!=nil{
log.Println(e)

}

}


func main() {



var (

 clientadd string

 serveradd string

 e error

 client *snet.Addr

 server *snet.Addr

 connectUDP *snet.Conn

)

flag.StringVar(&clientadd, "c", "", "address of client") // fetch values from command line

flag.StringVar(&serveradd, "s", "", "address of server")

flag.Parse()


client, e = snet.AddrFromString(clientadd)  // AddrFromString converts an address string of format isd-as,[ipaddr]:port


logerror(e)

server, e = snet.AddrFromString(serveradd)

logerror(e)

daddr := "/run/shm/dispatcher/default.sock"


snet.Init(client.IA, sciond.GetDefaultSCIONDPath(nil), daddr) //initialises scion network


connectUDP, e = snet.DialSCION("udp4", client, server) // client connects to server through UDP

logerror(e)

receivePacketBuffer := make([]byte, 2500) //Creating a buffer array of specified size

sendPacketBuffer := make([]byte, 16)   //Creating a buffer array of specified size


seed := rand.NewSource(time.Now().UnixNano())
var D int64 = 0

id := rand.New(seed).Uint64() // id for the send packet
n := binary.PutUvarint(sendPacketBuffer, id)
sendPacketBuffer[n] = 0

time_sent := time.Now() // sending the time now
_, e = connectUDP.Write(sendPacketBuffer) //sending message to server
logerror(e)

_, _, e = connectUDP.ReadFrom(receivePacketBuffer) // receiving message from server
logerror(e)

ret_id, n := binary.Uvarint(receivePacketBuffer)
if ret_id == id {// checking the id received from the server

	time_received, _ := binary.Varint(receivePacketBuffer[n:]) // taking the time recived from received packet
	diff := (time_received - time_sent.UnixNano())// finding the difference
	D += diff
}

var difference float64 = float64(D)

fmt.Printf("\nSource: %s\nDestination: %s\n", clientadd, serveradd);
fmt.Println("Time estimates:")

fmt.Printf("\tRTT - %.3fms\n", difference/1e6)
	fmt.Printf("\tLatency - %.3fms\n", difference/2e6)
}
