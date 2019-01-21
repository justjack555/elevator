package common

import (
	"bufio"
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
	Read all bytes from file of length size into
	a byte array
 */
func getBytesFromReader(br *bufio.Reader, size int64) []byte {
	b := make([]byte, size)
	for {
		_, err := br.Read(b)
		if err != nil {
			if err != io.EOF {
				log.Fatalln("ERR: While reading from file: ", err)
			}
			log.Println("Reached EOF. Breaking...")
			break;
		}
	}

//	log.Println("LOAD_CONFIG(): Byte slice read from file is: ", string(b))
	return b
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
	br := bufio.NewReader(f)
	b := getBytesFromReader(br, filesize)

	err = yaml.Unmarshal(b, c)
	if err != nil {
		log.Fatal("ERR: ", err)
	}
}