package common

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"math/rand"
	"os"
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

/**
	Read the server configuration for the
	master from config file
 */
func LoadConfig(path string, c interface{}) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("ERR: Unable to read config file. Terminating...")
	}
	defer f.Close()

	fileinfo, err := f.Stat()
	if err != nil {
		log.Fatalln("ERR: Unable to stat config file info: ", err)
	}

	filesize := fileinfo.Size()
	b := make([]byte, filesize)

	br := bufio.NewReader(f)
	log.Println("LOAD_CONFIG(): Started new buffered reader...")
	for {
		n, err := br.Read(b)
		if err != nil {
			if err != io.EOF {
				log.Fatalln("ERR: While reading from file: ", err)
			}
			log.Println("Reached EOF. Breaking...")
			break;
		}

		log.Println("LOAD_CONFIG(): Number of bytes read: ", n)
		log.Println("LOAD_CONFIG(): Byte slice read from file is: ", string(b[:n]))
	}

	log.Println("LOAD_CONFIG(): Byte slice read from file is: ", string(b))

	err = yaml.Unmarshal(b, c)
	if err != nil {
		log.Fatal("ERR: ", err)
	}

	fmt.Println("LOAD_CONFIG(): MasterConfig is: ", c)
}