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
	"net/rpc"
	"errors"
)

// This will be moved to properties file or to front - end
const numMasters = 3
var allMasters master.Master[]

// Do all of master rpc setup
func doSetup(i int) error {
	log.Println("Registering the ", i, "th master...")
	err := rpc.Register(master.New())
	if err != nil {
		log.Fatal("ERR: Registering master ", i, " - ", err, ". Terminating...")
	}

	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":123" + string(i))
	if e != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)

	return nil
}

// Main just starts masters
func main(){
	for i := 0; i < numMasters; i++ {
		go doSetup(i)
	}
}

