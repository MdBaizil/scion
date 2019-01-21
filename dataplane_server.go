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

 func logerror(e error){    //Error function

if e!=nil{

 log.Println(ef)

}

}

func main(){ //main function



var (

 serveradd string

 e error

 Server *snet.Addr

 connectUDP *snet.Conn

)
flag.StringVar(&serveradd, "s", "", "address of scion server")  // fetch values from command line

 flag.Parse()

 Server, e = snet.AddrFromString(serveradd)      // AddrFromString converts an address string of format isd-as,[ipaddr]:port

 logerror(e)



 daddr := "/run/shm/dispatcher/default.sock"

	snet.Init(ser.IA, sciond.GetDefaultSCIONDPath(nil), daddr)  //initialises scion network

  scionconnection, e = snet.ListenSCION("udp4", Server) // server will listen for client connections

 logerror(e)

 logerror(ef)

receivePacketBuffer := make([]byte, 2500)  //Creating a buffer array of specified size



for {

  		n, clientAddr, ef := scionconnection.ReadFrom(receivePacketBuffer) // receives packet from client

  	 logerror(e)



  		// Packet received, send back response to same client 



  		_, e = scionconnection.WriteTo(receivePacketBuffer[:n], clientAddr) //Hint given in tutorial (sending back the response to client)

  		 logerror(ef)

  		fmt.Println("Scion connection from", clientAddr)

  	}
}
