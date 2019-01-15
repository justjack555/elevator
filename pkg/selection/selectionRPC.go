package selection

import (
	"log"
	"time"
	"github.com/justjack555/elevator/pkg/common"
)

/**
	Request a new elevator for pickup from the selector
 */
func (s *Selection) AssignElevator(req common.MasterAssignRequest, reply *common.MasterAssignReply) error {
	log.Println("Selection.AssignElevator():")
//	log.Println("Selection.AssignElevator(): Request: NumPeople: ", req.NumPeople, ", Floor: ", req.Floor)

	e :=  &common.Elevator {
		IsActive : true,
		CurrentFloor : 5,
		LastChecked : time.Now(),
		Direction : common.DOWN,
	}

	reply.Elevator = e

//	log.Println("Selection.AssignElevator(): Reply: ", e)

	return nil
}

func (s *Selection) Select(req common.ElevatorLocationRequest, reply *common.ElevatorLocationReply) error {
	log.Println("Selection.Select():")

	return nil
}