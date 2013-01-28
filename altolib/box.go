package altolib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrNoBoxDb = errors.New("No box db")

	BoxDbFileName = filepath.Join(ConfigDirName, "boxes.json")
)

type Box struct {
	Id   string
	Vm   Vm
	Disk Disk
	Cpus int
	Ram  int
}

func BoxList() ([]Box, error) {
	boxes := make([]Box, 0)
	var err error

	fname := BoxDbFileName
	if !path_exists(filepath.Dir(fname)) {
		err = os.MkdirAll(filepath.Dir(fname), 0755)
		return nil, err
	}
	if !path_exists(fname) {
		err = ErrNoBoxDb
		return nil, err
	}
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(&boxes)
	if err != nil {
		err = fmt.Errorf("altolib.box: empty file [%s] ? (got: %v)", fname, err)
		return nil, err
	}
	return boxes, err
}

// EOF
