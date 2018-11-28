package master

import(
	"log"
	"net/rpc"
	"net/http"
	"net"
	"sync"
//	"errors"
	"github.com/justjack555/elevator/pkg/worker"
)

type Master struct {
	workers []worker
}

// Struct master returns once it completes serving
type masterResponse struct {
	masterIndx int
	masterErr error
}

// Structure that maintains state of masters
type masterState struct {
	allMasters []*Master
	mux sync.Mutex
}

// This will be moved to properties file or to front - end
const PORT = ":123"

/**
 Defacto constructor to return new master
* */
func new() *Master{
	return &Master{
		workers : nil,
	}
}

/**
	Helper to start an HTTP server and send its termination error
	to a waiting channel
**/
func serveAndReturnErr(l net.Listener, serverErrChan chan error){
	serverErrChan <- http.Serve(l, nil)
}
/**
	Setup and Launch each RPC Server that functions as a master
**/
func doMasterSetup(indx int, ch chan *masterResponse){
	serverErrChan := make(chan error)
	log.Println("Registering the ", indx, "th master...")
	m := new()

	// Register new master as RPC server
	err := rpc.Register(m)
	if err != nil {
		log.Fatal("ERR: Registering master ", indx, " - ", err, ". Terminating...")
	}

	// The masters will listen for HTTP Requests
	rpc.HandleHttp()
	l, err := net.Listen("tcp", PORT + string(indx))
	if err != nil {
		log.Fatal("ERR: Listen error:", err)
	}

	// Start the server
	go serveAndReturnErr(l, serverErrChan)

	// Server has returned an error. Handle this and return
	res := <- serverErrChan
	ch <- &masterResponse{
		masterIndx: indx,
		masterErr: res,
	}
}

/**
	Start designated number of masters
**/
func Start(numMasters int) []error {
	var errorList []error
	ch := make(chan *masterResponse)


	// Setup and Launch each RPC Server
	for i := 0; i < numMasters; i++ {
		go doMasterSetup(i, ch)
	}

	// Collect each RPC Server's response
	for i := 0; i < numMasters; i++ {
		res := <- ch
		errorList[res.masterIndx] = res.masterErr
	}

	return errorList
}
