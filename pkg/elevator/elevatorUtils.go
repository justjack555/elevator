package master

import (
	"github.com/justjack555/elevator/pkg/common"
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

const configPATH string = "cmd/master/master.yaml"

/**
	Read the server configuration for the
	master from config file
 */
func LoadConfig() *MasterConfig {
	mc := new(MasterConfig)
	common.LoadConfig(configPATH, mc)

//	fmt.Println("LOAD_CONFIG(): MasterConfig is: ", *mc)
	return mc
}