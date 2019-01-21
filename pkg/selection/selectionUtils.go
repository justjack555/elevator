package selection

import (
	"github.com/justjack555/elevator/pkg/common"
)

type ElevatorConfig struct {
	Address string
	Ports []string
}

type SelectionConfig struct {
	Num_instances int
	Ports []string
	Elevator_config *ElevatorConfig
}

const configPATH string = "cmd/selection/selection.yaml"

/**
	Read the server configuration for the
	master from config file
 */
func LoadConfig() *SelectionConfig {
	sc := new(SelectionConfig)
	common.LoadConfig(configPATH, sc)

//	fmt.Println("LOAD_CONFIG(): SelectionConfig is: ", *sc)
	return sc
}