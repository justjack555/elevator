package master

import (
	"log"
	"fmt"
	"net/http"
//	"github.com/justjack555/elevator/pkg/common"
)

/**
	Handler for calls to elevator/*

	Parameters passed in as part of the request will
	be utilized to form the RPC request to the selection
	service.
 */
func elevatorHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		Here we'll make the RPC call to the Selection Service
		log.Println("ElevatorHandler():")

		fmt.Fprintf(w, "Elevator handler...\n")
	})
}