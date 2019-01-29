package master

import (
	"log"
	"fmt"
	"net/http"
	"net/rpc"
	"net/url"
	"strconv"

	"github.com/justjack555/elevator/pkg/common"
)
const DEFAULT_NUM_PEOPLE string = "1"
const DEFAULT_FLOOR string = "1"
const DEFAULT_DIR string = "0"

func addDirection(args *common.MasterAssignRequest, q url.Values) error {
	d := q.Get("direction")
	if d == "" {
		log.Println("Direction not specified - assuming neutral...")
		d = DEFAULT_DIR
	}

	log.Println("Hanlder.buildArgs(): Direction is: ", d)
	numD, err := strconv.Atoi(d)
	if err != nil {
		log.Println("Invalid direction passed in. Throwing error and exiting...")
		return err
	}

	args.Direction = numD
	return nil
}

func addFloor(args *common.MasterAssignRequest, q url.Values) error {
	f := q.Get("floor")
	if f == "" {
		log.Println("Floor not specified - assuming ground floor...")
		f = DEFAULT_FLOOR
	}

	log.Println("Hanlder.buildArgs(): Num floor is: ", f)
	numF, err := strconv.Atoi(f)
	if err != nil {
		log.Println("Invalid number of floor passed in. Throwing error and exiting...")
		return err
	}

	args.Floor = numF
	return nil
}

func addPeople(args *common.MasterAssignRequest, q url.Values) error {
	ppl := q.Get("people")
	if ppl == "" {
		log.Println("Number of people not specified - assuming 1 person...")
		ppl = DEFAULT_NUM_PEOPLE
	}

	log.Println("Hanlder.buildArgs(): Num people is: ", ppl)
	numPpl, err := strconv.Atoi(ppl)
	if err != nil {
		log.Println("Invalid number of people passed in. Throwing error and exiting...")
		return err
	}

	args.NumPeople = numPpl
	return nil
}

func buildArgs(u *url.URL) (*common.MasterAssignRequest, error) {
	args := new(common.MasterAssignRequest)
	log.Println("Hanlder.buildArgs(): Request URI is: ", u.String())
	q := u.Query()

	if 	err := addPeople(args, q); err != nil {
		log.Println("Hanlder.buildArgs(): Unable to obtain people parameter")
		return nil, err
	}

	if 	err := addFloor(args, q); err != nil {
		log.Println("Hanlder.buildArgs(): Unable to obtain floor parameter")
		return nil, err
	}

	if 	err := addDirection(args, q); err != nil {
		log.Println("Hanlder.buildArgs(): Unable to obtain direction parameter")
		return nil, err
	}


	return args, nil
}

func requestElevator(client *rpc.Client, req *http.Request) *common.Elevator{
	var reply common.MasterAssignReply

	args, err := buildArgs(req.URL)
	if err != nil {
		log.Fatal("Handler.RequestElevator:", err)
	}

	err = client.Call("Selection.AssignElevator", args, &reply)
	if err != nil {
		log.Fatal("Handler.RequestElevator:", err)
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
func elevatorHandler(selConfig *SelectionConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Println("ElevatorHandler():")

		client, err := rpc.DialHTTP("tcp", selConfig.Address + common.ChoosePort(selConfig.Ports))
		if err != nil {
			log.Fatal("ERR: Dialing error: ", err)
		}

		e := requestElevator(client, r)

		fmt.Fprintf(w, "Elevator details: %v ...\n", e)
	})
}