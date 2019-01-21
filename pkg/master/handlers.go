package master

import (
	"log"
	"fmt"
	"net/http"
	"net/rpc"

	"github.com/justjack555/elevator/pkg/common"
)

func requestElevator(client *rpc.Client) *common.Elevator{
	var reply common.MasterAssignReply

	args := &common.MasterAssignRequest {
		NumPeople : 2,
		Floor : 1,
		Direction : -1,
	}

	err := client.Call("Selection.AssignElevator", args, &reply)
	if err != nil {
		log.Fatal("Selection.AssignElevator:", err)
	}

//	log.Println("Handler.RequestElevator(): Reply: ", reply.Elevator)

	return reply.Elevator
}
/**
	Handler for calls to elevator/*

	Parameters passed in as part of the request will
	be utilized to form the RPC request to the selConfig
	service.
 */
func elevatorHandler(selConfig SelectionConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		Here we'll make the RPC call to the Selection Service
		log.Println("ElevatorHandler():")

		client, err := rpc.DialHTTP("tcp", selConfig.Address + common.ChoosePort(selConfig.Ports))
		if err != nil {
			log.Fatal("ERR: Dialing error: ", err)
		}

		e := requestElevator(client)

		fmt.Fprintf(w, "Elevator details: %v ...\n", e)
	})
}