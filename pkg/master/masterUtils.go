package master

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

type SelectionConfig struct {
	Address string
	Ports []string
}

type MasterConfig struct {
	Num_instances int
	Ports []string
	Selection_config SelectionConfig
}

const MaxBYTES int = 4096
/**
	Read the server configuration for the
	master from config file
 */
func LoadConfig() *MasterConfig {
	mc := new(MasterConfig)


	f, err := os.Open("cmd/master/master.yaml")
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

	err = yaml.Unmarshal(b, mc)
	if err != nil {
		log.Fatal("ERR: ", err)
	}

	fmt.Println("LOAD_CONFIG(): MasterConfig is: ", *mc)
	return mc
}