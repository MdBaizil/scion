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
	data_SIZE int = 2000
	total_NUM int = 5
)

type Checkpoint struct {
	S, R int64
}

var (

	recvHash map[uint64]*Checkpoint // intialising datatype map
	udpConnection *snet.Conn

)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func AvBottleneckBW() (float64, float64) { // finding average bottleneck bandwidth for sent and received packets

sorted := make([]*Checkpoint, total_NUM) // sorting checkpoints
	i := 0
	for _, c := range recvHash {
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
  receivedbandwidth := float64(data_SIZE*8*1e3) / (float64(R_int) / float64(total_NUM-1))

  return sentbandwidth, receivedbandwidth
  }

  func sendPackets() {
  var e error

  sendPacketBuffer := make([]byte, data_SIZE + 1)

  seed := rand.NewSource(time.Now().UnixNano()) //
	iters := 0
	for iters < (total_NUM) {
		iters += 1

		id := rand.New(seed).Uint64()
		_ = binary.PutUvarint(sendPacketBuffer, id)

		recvHash[id] = &Checkpoint{time.Now().UnixNano(), 0}
		_, err = udpConnection.Write(sendPacketBuffer)  // writing packets to server
		check(e)

	}
}
func recvPackets() int {

	var e error
	receivePacketBuffer := make([]byte, data_SIZE + 1)

  for count < total_NUM {

  _, _, e = udpConnection.ReadFrom(receivePacketBuffer) // receiving message from server
  check(e)

  ret_id, n := binary.Uvarint(receivePacketBuffer)
  if ret_id == id { // checking the id received from the server
  	time_received, _ := binary.Varint(receivePacketBuffer[n:]) // taking the time recived from received packet
 count += 1
 }
 return count
 }

 func main() {
	var (
		sourceAddr string
		destinationAddr string

    e    error
		source  *snet.Addr
		destination *snet.Addr
  )

  flag.StringVar(&sourceAddr, "s", "", "Source SCION Address")
	flag.StringVar(&destinationAddr, "d", "", "Destination SCION Address")
	flag.Parse()


		source, err = snet.AddrFromString(sourceAddr)// creating udp connection
		check(e)

    destination, err = snet.AddrFromString(destinationAddr)
		check(e)

    dAddr := "/run/shm/dispatcher/default.sock"
	snet.Init(source.IA, sciond.GetDefaultSCIONDPath(nil), dAddr)



	udpConnection, err = snet.DialSCION("udp4", source, destination)
	check(e)

  recvHash = make(map[uint64]*Checkpoint) // creating hash table of checkpoints

     sendPackets()
	count := recvPackets()

  fmt.Println("# packets:", count)
	if count == 0 {
		check(fmt.Errorf("No packets received from server"))
	}

  sentbandwidth, receivedbandwidth := AvBottleneckBW()

  fmt.Printf("\nSource: %s\nDestination: %s\n", sourceAddress, destinationAddress);
	fmt.Println("Rate sent:")
	fmt.Printf("\tBW - %.3fMbps\n", sentbandwidth)
	fmt.Println("Bottleneck Bandwidth estimate:")
	fmt.Printf("\tBW - %.3fMbps\n", receivedbandwidth)
}
