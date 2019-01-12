package selection

import (
	"log"
	"container/heap"
	"github.com/justjack555/elevator/pkg/common"
)

/**
	SortedElevators implements sort.Interface for []*Elevator based on
	the currentFloor field.
**/
type SortedElevators []*common.Elevator

type Selection struct{
	elevators SortedElevators
}

func (se SortedElevators) Len() int {
	return len(se)
}
func (se SortedElevators) Swap(i, j int) {
	se[i], se[j] = se[j], se[i]
}

func (se SortedElevators) Less(i, j int) bool {
	return se[i].getCurrentFloor() < se[j].getCurrentFloor()
}