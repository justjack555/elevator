package location

import (
	"log"
	"github.com/justjack555/elevator/pkg/common"
)

func (l *Location) Compute(req common.LocationRequest, reply *common.LocationReply) error {
	log.Println("Location.Compute():")

	return nil
}