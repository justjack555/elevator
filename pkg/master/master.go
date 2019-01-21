package master

import(
	"log"
	"net/http"
	"sync"
	"github.com/justjack555/elevator/pkg/common"
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
//const PORT = ":123"

/**
 Defacto constructor to return new master
* */
func createMaster() *Master{
	log.Println("Master.CreateMaster(): Invoked")
	return new(Master)
}

/**
	Load all of the HTTP handler functions
	for requests
**/
func loadHandlers(selConfig *SelectionConfig) {
	routes := map[string] http.Handler {
		"/elevator/" : elevatorHandler(selConfig),
	}

	for pattern, handler := range routes {
		http.Handle(pattern, handler)
	}
}

/**
	Helper to start an HTTP server and send its termination error
	to a waiting channel
**/
func serveAndReturnErr(port string, serverErrChan chan error){
	serverErrChan <- http.ListenAndServe(port, nil)
}

/**
	Setup and Launch each HTTP Server that functions as a master
	Masters have their route handling registered and then setup to
	listen

	If they return an error, this is handled by passing the value into
	the provided channel
**/
func launchMaster(indx int, selConfig *SelectionConfig, port string, ch chan *masterResponse){
	serverErrChan := make(chan error)
	log.Println("Registering the ", indx, "th master...")

	loadHandlers(selConfig)

	go serveAndReturnErr(common.ConstructPort(port), serverErrChan)

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
func Start(numMasters int, selConfig *SelectionConfig, ports []string) []error {
	errorList := make([]error, numMasters, numMasters)
	ch := make(chan *masterResponse)


	for i := 0; i < numMasters; i++ {
		go launchMaster(i, selConfig, ports[i], ch)
	}

	for i := 0; i < numMasters; i++ {
		res := <- ch
		log.Println("ERR: Master #", res.masterIndx, " failed with error: ", res.masterErr)
		errorList[res.masterIndx] = res.masterErr
	}

	return errorList
}
