package altolib

import (
// "bytes"
// "fmt"
// "os/exec"
// "regexp"
// "strconv"
// "strings"
)

type Vm struct {
	Id  string
	Tag string
}

func VmList() ([]Vm, error) {
	vms := make([]Vm, 0)
	var err error

	return vms, err
}

// EOF
