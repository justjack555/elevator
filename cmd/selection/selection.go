/**

Initializes our selector by starting
an RPC service

The selector receives requests from the master
for the optimal elevator and returns the elevator
to be used

Number of selectors should be a configurable parameter - we hard code this for now

**/
package main

import(
	"log"
	"github.com/justjack555/elevator/pkg/selection"
)

// Move to config file/UI - currently only support one master
const numSelectors = 1


// Main just starts masters
func main(){
	sc := selection.LoadConfig()
	errList := selection.Start(sc.Num_instances, sc.Ports, sc.Elevator_config)
	for i, err := range errList {
		log.Println("ERR: ", i, "th selector terminated with error: ", err)
	}

}

