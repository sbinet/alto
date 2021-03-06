package altolib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrNoVmDb = errors.New("No VM db")

	VmDbFileName = filepath.Join(ConfigDirName, "vms.json")
)

type Vm struct {
	Id  string
	Tag string
}

func (vm Vm) String() string {
	return fmt.Sprintf("Vm{Id=%s Tag=%q}", vm.Id, vm.Tag)
}

func VmList() ([]Vm, error) {
	vms := make([]Vm, 0)
	var err error

	if !path_exists(VmDbFileName) {
		return nil, ErrNoVmDb
	}

	f, err := os.Open(VmDbFileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&vms)
	if err != nil {
		return nil, err
	}
	return vms, err
}

// EOF
