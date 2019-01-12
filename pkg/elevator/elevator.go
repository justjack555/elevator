package elevator

import(
	"log"
	"time"
	"github.com/justjack555/elevator/pkg/common"
)

func (e *Elevator) getCurrentFloor() int{
	return e.currentFloor
}
