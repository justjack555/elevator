package selection

import (
	"log"
	"net"
	"net/rpc"
	"net/http"

	"github.com/justjack555/elevator/pkg/common"
)

// Struct master returns once it completes serving
type selectorResponse struct {
	selectorIndx int
	selectorErr error
}

/**
	SortedElevators implements sort.Interface for []*Elevator based on
	the currentFloor field.
**/
type SortedElevators []*common.Elevator

type Selection struct{
	elevators SortedElevators
}

// This will be moved to config file
const PORT = "8081"

func (se SortedElevators) Len() int {
	return len(se)
}
func (se SortedElevators) Swap(i, j int) {
	se[i], se[j] = se[j], se[i]
}

func (se SortedElevators) Less(i, j int) bool {
	return se[i].GetCurrentFloor() < se[j].GetCurrentFloor()
}

/**
	Helper to start an RPC server and send its termination error
	to a waiting channel
**/
func serveAndReturnErr(l net.Listener, serverErrChan chan error){
	serverErrChan <- http.Serve(l, nil)
}

/**
	Return a selection structure
	with an empty list of elevators
 */
func createSelector() *Selection {
	return &Selection {
		elevators: make(SortedElevators, 0, 0),
	}
}

/**
	Setup and Launch each RPC Server that functions as a selector

	Using DefaultServeMux since we will only have one Selection server
	to start

	If they return an error, this is handled by passing the value into
	the provided channel
**/
func launchSelector(indx int, ch chan *selectorResponse){
	serverErrChan := make(chan error)
	log.Println("Registering the ", indx, "th selector...")

	s := createSelector()

	err := rpc.Register(s)
	if err != nil {
		log.Println("ERR: Registering selector ", indx, " - ", err, ". Terminating this server...")
		ch <- &selectorResponse{
			selectorIndx: indx,
			selectorErr: err,
		}
		return
	}

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", common.ConstructPort(PORT))
	if err != nil {
		log.Println("ERR: Listen error:", err)
		ch <- &selectorResponse{
			selectorIndx: indx,
			selectorErr: err,
		}
		return
	}

	go serveAndReturnErr(l, serverErrChan)

	res := <- serverErrChan
	ch <- &selectorResponse{
		selectorIndx: indx,
		selectorErr: res,
	}
}

/**
	Start designated number of masters
	and wait for any errors
**/
func Start(numSelectors int) []error {
	errorList := make([]error, numSelectors, numSelectors)
	ch := make(chan *selectorResponse)


	for i := 0; i < numSelectors; i++ {
		go launchSelector(i, ch)
	}

	for i := 0; i < numSelectors; i++ {
		res := <- ch
		log.Println("ERR: Master #", res.selectorIndx, " failed with error: ", res.selectorErr)
		errorList[res.selectorIndx] = res.selectorErr
	}

	return errorList
}