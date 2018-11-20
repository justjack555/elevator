package master

import(
	"fmt"
	"github.com/justjack555/elevator/pkg/worker"
)

type Master struct {
	workers []worker
}

/**
 Defacto constructor to return new master
* */
func New() *Master{
	return &Master{
		workers : nil,
	}
}
