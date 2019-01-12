package master

import(
	"log"
	"github.com/justjack555/elevator/pkg/common"
)

/**
	Request a new elevator for pickup from the master
 */
func (m *Master) AssignElevator(req common.MasterAssignRequest, reply *common.MasterAssignReply) error {
	log.Println("Calling Assign Elevator")

	return nil
}