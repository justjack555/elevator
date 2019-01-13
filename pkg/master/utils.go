package master

import (
	"strings"
	"strconv"
)

func constructPort(indx int) string {
	// This will be moved to config file
	const PORT = ":123"
	var portStr strings.Builder

	portStr.WriteString(PORT)
	portStr.WriteString(strconv.Itoa(indx))

	return portStr.String()
}