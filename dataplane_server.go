// Code references : https://github.com/netsec-ethz/scion-homeworks/blob/master/latency/timestamp_server.go and reference https://github.com/perrig/scionlab/blob/master/sensorapp/sensorserver/sensorserver.go

package  main



import (

 "fmt"             //importing fmt package for printing

 "log"

 "flag"

 "encoding/binary"

 "time"

 "github.com/scionproto/scion/go/lib/snet" //importing snet packages for the scion connections

 "github.com/scionproto/scion/go/lib/sciond"

 )

 func geterror(e error){    //Error function

if e!=nil{

 log.Println(e)

}

}

func main(){ //main function



var ( // variable declarations

 serveradd string

 e error

 Server *snet.Addr

 connectUDP *snet.Conn

)
flag.StringVar(&serveradd, "s", "", "adress of scion server")  // fetch  server address from command line

 flag.Parse()

 Server, e = snet.AddrFromString(serveradd)      // AddrFromString converts an address string of format isd-as,[ipaddr]:port

 geterror(e)



 daddr := "/run/shm/dispatcher/default.sock"

	snet.Init(Server.IA, sciond.GetDefaultSCIONDPath(nil), daddr)  //initialises scion network

  connectUDP, e = snet.ListenSCION("udp4", Server) // server will listen for client connections

 geterror(e)



packetreceived := make([]byte, 3000)  //making packet array of size 2500 for receiving


for {

  		k, clientAddr, e := connectUDP.ReadFrom(packetreceived)  //Reads the receiver buffer

  	 geterror(e)



  		// send back the same packet with timestamp

  		m := binary.PutVarint(packetreceived[k:], time.Now().UnixNano())  // adding time to received array

  		_, e = connectUDP.WriteTo(packetreceived[:k+m], clientAddr)  //sending back the response to client

  		 geterror(e)

  		fmt.Println("received packets from client", clientAddr)

  	}

}
