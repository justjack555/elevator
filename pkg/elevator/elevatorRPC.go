package elevator

import (
	"log"
	"github.com/justjack555/elevator/pkg/common"
)

func (e *Elevator) Status(req common.ElevatorStatusRequest, reply *common.ElevatorStatusReply) error {
	log.Println("ElevatorRPC.Status(): ")
	return nil
}

func (e *Elevator) Locate(req common.ElevatorLocationRequest, reply *common.ElevatorLocationReply) error {
	log.Println("ElevatorRPC.Locate():")
	return nil
}