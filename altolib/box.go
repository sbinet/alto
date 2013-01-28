package altolib

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"os"
	// "os/exec"
	"path/filepath"
	// "regexp"
	// "strconv"
	// "strings"
)

type Box struct {
	Id   string
	Tag  string
	Vm   Vm
	Disk Disk
	Cpus int
	Ram  int
}

func BoxList() ([]Box, error) {
	boxes := make([]Box, 0)
	var err error

	fname := os.ExpandEnv("${HOME}/.config/alto/boxes.json")
	if !path_exists(filepath.Dir(fname)) {
		err = os.MkdirAll(filepath.Dir(fname), 0755)
		return nil, err
	}
	if !path_exists(fname) {
		err = fmt.Errorf("altolib.box: no such file [%s]", fname)
		return nil, err
	}
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(&boxes)
	return boxes, err
}

// EOF
