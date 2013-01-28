package altolib

import (
// "bytes"
// "fmt"
// "os/exec"
// "regexp"
// "strconv"
// "strings"
)

type Box struct {
	Id  string
	Tag string

	Vm   Vm
	Disk Disk
	Cpus int
	Ram  int
}

func BoxList() ([]Box, error) {
	boxes := make([]Box, 0)
	var err error

	return boxes, err
}

// EOF
