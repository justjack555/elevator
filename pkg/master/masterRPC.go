package master

import(
	"log"
	"github.com/justjack555/elevator/pkg/common"
)

func (m *Master) AssignWorker(req common.MasterAssignRequest, reply *common.MasterAssignReply) error {
	log.Println("Calling Assign Worker")

	return nil
}