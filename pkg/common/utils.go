package common

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

/**
	Obtain port string in proper format
	for http/rpc API
 */
func ConstructPort(port string) string {
	var portStr strings.Builder

	portStr.WriteString(":")
	portStr.WriteString(port)

	return portStr.String()
}

/**
	Random selection of available services so that we distribute workload
 */
func ChoosePort(ports []string) string {
	numPorts := len(ports)
	if numPorts <= 0 {
		log.Fatalln("ChoosePort(): ERR: No ports to listen on. Terminating...")
	}

	rand.Seed(time.Now().UnixNano())
	port := ports[rand.Intn(numPorts)]
	return ConstructPort(port)
}