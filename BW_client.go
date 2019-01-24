// Code references : https://github.com/netsec-ethz/scion-homeworks/blob/master/bottleneck_bw_est/v1_bw_est_client.go and reference https://github.com/perrig/scionlab/blob/master/sensorapp/sensorserver/sensorserver.go

package main

import (
	"flag"
	"encoding/binary"
	"fmt" //importing fmt package for printing
	"log"
	"math/rand" //importing for mathematical operations

	"sort"  //importing package for sorting
	"time"

	"github.com/scionproto/scion/go/lib/sciond"
	"github.com/scionproto/scion/go/lib/snet"  //importing snet packages for the scion connections

	"github.com/scionproto/scion/go/lib/spath" //importing packages for finding path
	"github.com/scionproto/scion/go/lib/spath/spathmeta"
)

const (
	data_SIZE int = 2000 // the data size for the packets
	total_NUM int = 5 // total no of time the packets needs to sent or received
)

type Checkpoint struct {
	S, R int64
}

var (

	recvHash map[uint64]*Checkpoint // intialising datatype map
	connectUDP *snet.Conn

)

func geterror(e error) { // error fuction
	if e != nil {
		log.Fatal(e)
	}
}

func AvBottleneckBW() (float64, float64) { // finding average bottleneck bandwidth for sent and received packets

sorted := make([]*Checkpoint, total_NUM) // sorting checkpoints
	i := 0
	for i < total_NUM {
		if c.R != 0 {
			sorted[i] = c
			i += 1
		}
	}
  sort.Slice(sorted, func(i, j int) bool { return sorted[i].S < sorted[j].S })

  var S_int, R_int int64 //intialising variables

  for i := 1; i < total_NUM; i+=1 {
		S_int += (sorted[i].S_int - sorted[i-1].S_int)// finding difference intervals between sent packets
		R_int += (sorted[i].R_int - sorted[i-1].R_int)// finding difference intervals between received packets
	}

  sentbandwidth := float64(data_SIZE*8*1e3) / (float64(S_int) / float64(total_NUM-1)) // finding bandwidth using equation size/time
  receivedbandwidth := float64(data_SIZE*8*1e3) / (float64(R_int) / float64(total_NUM-1)) // finding the received bandwidth

  return sentbandwidth, receivedbandwidth
  }

func packets() { // function for sending and receiving packets
  var e error

  sentpacket := make([]byte, 3000)

  seed := rand.NewSource(time.Now().UnixNano()) // generating new random seed for sending
	iters := 0
	for iters < (total_NUM) {
		iters += 1

		id := rand.New(seed).Uint64() // generating random id for sending packect(which should be returned by server)
		_ = binary.PutUvarint(sentpacket, id)

		recvHash[id] = &Checkpoint{time.Now().UnixNano(), 0} // creating hash table for sentpackets
		_, err = connectUDP.Write(sentpacket)  // writing packets to server
		geterror(e)


   receivedpackets := make([]byte, 3000)

  for count < total_NUM {

  _, _, e =connectUDP.ReadFrom(receivedpackets) // receiving message from server
  geterror(e)

  ret_id, n := binary.Uvarint(receivedpackets)
  if ret_id == id { // checking the id received from the server is same that sent
  	Tr, _ := binary.Varint(receivedpackets[n:]) // taking the time recived from received packet
 count += 1
 count := receivedpackets
  }

 return sentpacket, count
  }

func main() {
var (
     sourceAddr string
     destinationAddr string

     e    error
     client  *snet.Addr
     server *snet.Addr
  )

	flag.StringVar(&clientadd, "c", "", "address of client") // fetch address of the client from command line

	flag.StringVar(&serveradd, "s", "", "address of server")// fetching address of the server from command line

	flag.Parse()


		source, err = snet.AddrFromString(clientadd)// AddrFromString converts an address string of format isd-as,[ipaddr]:port
		geterror(e)

    destination, err = snet.AddrFromString(serveradd)
		geterror(e)

    dAddr := "/run/shm/dispatcher/default.sock"
	snet.Init(source.IA, sciond.GetDefaultSCIONDPath(nil), dAddr) // initializing scion connection



	connectUDP, err = snet.DialSCION("udp4", client, server) // connecting to server
	geterror(e)

  recvHash = make(map[uint64]*Checkpoint) // creating hash table of checkpoints

  packets()


  sentbandwidth, receivedbandwidth := AvBottleneckBW()

	fmt.Printf("\tBW - %.3fMbps\n", sentbandwidth)
	fmt.Println("Bottleneck bandwidth is:")
	fmt.Printf("\tBW - %.3fMbps\n", receivedbandwidth) 
}
