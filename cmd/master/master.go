/**

Initializes our master by starting
an RPC service

The master receives input from the sensors
and translates this into work for some number
of worker elevators

Number of masters should be a configurable parameter - we hard code this for now

**/
package main

import(
	"log"
	"github.com/justjack555/elevator/pkg/master"
)

// Move to config file/UI
const numMasters = 3


// Main just starts masters
func main(){
	errList := master.Start(numMasters)
	for i, err := range errList {
		log.Println("ERR: ", i, "th master terminated with error: ", err)
	}

}

