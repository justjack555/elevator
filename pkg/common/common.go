package common

import (
	"time"
)

type Direction int

const (
	UP Direction = 1
	NEUTRAL Direction = 0
	DOWN Direction = -1
)

/**
	Define Elevator structure here
	as it is the sole common structure
	that must be shared across services
**/
type Elevator struct {
	isActive bool
	currentFloor int
	lastChecked time.Time
	direction Direction
}

type MasterAssignRequest struct {
	numPeople int
	floor int
	direction int
}

type MasterAssignReply struct {
	elevator *Elevator
}

type ElevatorStatusRequest struct{}

type ElevatorStatusReply struct {
	currentFloor int
	direction Direction
}

/**
	Used by the selection service in the
	background to maintain an updated
	list of elevator locations
 */
type ElevatorLocationRequest struct {
	lastFloor int
	direction Direction
	lastTimestamp time.Time
}

type ElevatorLocationReply struct {
	currentFloor int
	newTimestamp time.Time
}

func (e *Elevator) GetCurrentFloor() int{
	return e.currentFloor
}
