package common

import (
	"github.com/justjack555/elevator/pkg/elevator"
	"time"
)

type Direction int

const (
	UP Direction = 1
	NEUTRAL Direction = 0
	DOWN Direction = -1
)

type MasterAssignRequest struct {
	numPeople int
	floor int
	direction int
}

type MasterAssignReply struct {
	elevator *elevator.Elevator
}

type ElevatorStatusRequest nil

type ElevatorStatusReply struct {
	currentFloor int
	direction Direction
}

type LocationRequest struct {
	lastFloor int
	direction Direction
	lastTimestamp time.Time
}

type LocationResponse struct {
	currentFloor int
	newTimestamp time.Time
}
