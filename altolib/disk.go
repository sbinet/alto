package altolib

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrNoDiskDb = errors.New("No disk db")

	DiskDbFileName = filepath.Join(ConfigDirName, "disks.json")
)

type Disk struct {
	Guid  string
	Id    string
	Count int
	Owner string
	Size  int
	Tag   string
}

func (d Disk) String() string {
	return fmt.Sprintf(
		"Disk{Guid=%s Id=%s Count=%d Owner=%q Size=%dGb Tag=%q}",
		d.Guid,
		d.Id,
		d.Count,
		d.Owner,
		d.Size,
		d.Tag,
	)
}

func DiskList() ([]Disk, error) {
	pdisks := make([]Disk, 0)
	var err error

	slab := exec.Command("stratus-describe-volumes")
	out, err := slab.Output()
	if err != nil {
		return nil, fmt.Errorf("altolib.disk: failed to run stratus-describe-volumes: %v", err)
	}
	pdisks_data := bytes.Split(out, []byte(":: DISK "))
	pdisks = make([]Disk, 0, len(pdisks_data))
	re_pat := regexp.MustCompile(`(?P<tag>.*?): (?P<value>.*)`)
	// if err != nil {
	// 	return nil, err
	// }
	for _, data := range pdisks_data {
		if bytes.Equal(data, []byte("")) {
			continue
		}
		//fmt.Printf("--- %q\n", string(data))
		lines := bytes.Split(data, []byte("\n"))
		pdisk := Disk{Guid: strings.Trim(string(lines[0]), " \t\r\n")}
		for _, line := range lines[1:] {
			//sline := strings.Trim(string(line), "\t\r\n")
			sline := string(line)
			//fmt.Printf("+++ %q\n", sline)
			m := re_pat.FindStringSubmatch(sline)
			if m != nil {
				tag := strings.Trim(m[1], " \t\r\n")
				val := strings.Trim(m[2], " \t\r\n")
				//fmt.Printf(">>> (%v) %v %q %q\n", len(m), m, tag, val)
				switch tag {
				case "owner":
					pdisk.Owner = val
				case "count":
					count, err := strconv.Atoi(val)
					if err != nil {
						return nil, err
					}
					pdisk.Count = count

				case "identifier":
					pdisk.Id = val

				case "size":
					sz, err := strconv.Atoi(val)
					if err != nil {
						return nil, err
					}
					pdisk.Size = sz
				case "tag":
					pdisk.Tag = val

				default:
					err = fmt.Errorf("altolib.pdisk: unknown field [%s] (val=%q)", tag, val)
					return nil, err
				}
			}
		}
		pdisks = append(pdisks, pdisk)
	}

	return pdisks, err
}

// EOF
