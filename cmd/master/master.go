/**

Initializes our master by starting
an RPC service

The master receives input from the sensors
and translates this into work for some number
of worker elevators

Number of masters should be a configurable parameter - we hard code this for now

**/
package main

import (
	"github.com/justjack555/elevator/pkg/master"
	"log"
)

// Move to config file/UI - currently only support one master
//const numMasters = 1


// Main just starts masters
func main(){
	mc := master.LoadConfig()
	errList := master.Start(mc.Num_instances, mc.Selection_config, mc.Ports)
	for i, err := range errList {
		log.Println("ERR: ", i, "th master terminated with error: ", err)
	}

}

