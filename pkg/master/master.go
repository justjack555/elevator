package master

import(
	"log"
	"net/rpc"
	"net/http"
	"net"
	"sync"
	"github.com/justjack555/elevator/pkg/elevator"
)

type Master struct {
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

// This will be moved to config file
const PORT = ":123"

/**
 Defacto constructor to return new master
* */
func createMaster() *Master{
	return &Master{
		elevators : nil,
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
	Masters are registered, established as HTTP Request listeners,
	and started.

	If they return an error, this is handled by passing the value into
	the provided channel
**/
func launchMaster(indx int, ch chan *masterResponse){
	serverErrChan := make(chan error)
	log.Println("Registering the ", indx, "th master...")
	m := createMaster()

	err := rpc.Register(m)
	if err != nil {
		log.Println("ERR: Registering master ", indx, " - ", err, ". Terminating this server...")
		ch <- err
		return
	}

	rpc.HandleHttp()
	l, err := net.Listen("tcp", PORT + string(indx))
	if err != nil {
		log.Println("ERR: Listen error:", err)
		ch <- err
		return
	}

	go serveAndReturnErr(l, serverErrChan)

	res := <- serverErrChan
	ch <- &masterResponse{
		masterIndx: indx,
		masterErr: res,
	}
}

/**
	Start designated number of masters
	and wait for any errors
**/
func Start(numMasters int) []error {
	var errorList []error
	ch := make(chan *masterResponse)


	for i := 0; i < numMasters; i++ {
		go launchMaster(i, ch)
	}

	for i := 0; i < numMasters; i++ {
		res := <- ch
		errorList[res.masterIndx] = res.masterErr
	}

	return errorList
}
