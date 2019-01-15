package common

import (
	"strings"
	"strconv"
)

func ConstructPort(portBase string, indx int) string {
	var portStr strings.Builder

	portStr.WriteString(portBase)
	portStr.WriteString(strconv.Itoa(indx))

	return portStr.String()
}