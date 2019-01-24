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

func geterror(e error) {   //Error function



if e!=nil{
log.Println(e)

}

}


func main() { // main function for finding latency



var ( // necessary variable declarations

 clientadd string

 serveradd string

 e error

 client *snet.Addr

 server *snet.Addr

 connectUDP *snet.Conn

)

flag.StringVar(&clientadd, "c", "", "address of client") // fetch address values from command line

flag.StringVar(&serveradd, "s", "", "address of server")// fetch server address from command line

flag.Parse()


client, e = snet.AddrFromString(clientadd)  // AddrFromString converts an address string of format isd-as,[ipaddr]:port


geterror(e)

server, e = snet.AddrFromString(serveradd)

geterror(e)

daddr := "/run/shm/dispatcher/default.sock"


snet.Init(client.IA, sciond.GetDefaultSCIONDPath(nil), daddr) //initialises scion network


connectUDP, e = snet.DialSCION("udp4", client, server) // client connects to server through UDP

geterror(e)

packetreceived := make([]byte, 3000) //Creating a  array of 2500 size for receiving

packetsent := make([]byte, 40)   //Creating a  array of 16 size for sending



seed := rand.NewSource(time.Now().UnixNano()) 


id := rand.New(seed).Uint64() // generating random id for sending packect(inorder to check wheather the server is responding to the same packect sent)
n := binary.PutUvarint(packetsent, id)
packetsent[n] = 0

Ts := time.Now() // sending the time now
_, e = connectUDP.Write(packetsent) //sending message to server
geterror(e)

_, _, e = connectUDP.ReadFrom(packetreceived) // receiving message from server
geterror(e)

	Tr, _ := binary.Varint(packetreceived[n:]) // taking the time recived from received packet
	diff := (Tr - Ts.UnixNano())// finding the difference



var difference float64 = float64(diff)

fmt.Printf("\tlatency is - %.3fs\n", difference/2e9) // result will be printed in seconds
}
