package elevator

import (
	"log"
	"github.com/justjack555/elevator/pkg/common"
)

func (e *Elevator) Status(req common.ElevatorStatusRequest, reply *common.ElevatorStatusReply) error {
	log.Println("elevatorRPC.Status(): ")
	return nil
}