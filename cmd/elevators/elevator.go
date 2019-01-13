/**

Initializes our elevators by starting
an RPC service for each

Number of masters should be a configurable parameter - we hard code this for now

**/
package main

import(
	"log"
	"github.com/justjack555/elevator/pkg/elevator"
)

// Move to config file/UI
const numElevators = 5


// Main just starts masters
func main(){
	errList := elevator.Start(numElevators)
	for i, err := range errList {
		log.Println("ERR: ", i, "th elevator terminated with error: ", err)
	}

}

