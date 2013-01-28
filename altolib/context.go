package altolib

import (
	"encoding/json"
	//"fmt"
	"os"
)

type Context struct {
	boxdb  map[string]Box
	diskdb map[string]Disk
	vmdb   map[string]Vm
}

func NewContext() (*Context, error) {
	var ctx *Context
	var err error

	ctx = &Context{
		boxdb:  make(map[string]Box),
		diskdb: make(map[string]Disk),
		vmdb:   make(map[string]Vm),
	}
	err = ctx.init()
	if err != nil {
		return nil, err
	}
	return ctx, err
}

func (ctx *Context) init() error {
	var err error

	// load pdisks
	disks, err := DiskList()
	if err != nil {
		return err
	}
	for _, disk := range disks {
		ctx.diskdb[disk.Guid] = disk
	}

	// load vms
	vms, err := VmList()
	if err != nil {
		return err
	}
	for _, vm := range vms {
		ctx.vmdb[vm.Id] = vm
	}

	// load boxes
	if !path_exists(BoxDbFileName) {
		var f *os.File
		f, err = os.Create(BoxDbFileName)
		defer f.Close()
		if err != nil {
			return err
		}
		boxes := make([]Box, 0)
		err = json.NewEncoder(f).Encode(&boxes)
		if err != nil {
			return err
		}
		err = f.Sync()
		if err != nil {
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
	}
	boxes, err := BoxList()
	if err != nil {
		if err != nil {
			return err
		}
	}
	for _, box := range boxes {
		ctx.boxdb[box.Id] = box
	}
	return err
}

func (ctx *Context) Boxes() []Box {
	boxes := make([]Box, 0, len(ctx.boxdb))
	for _, box := range ctx.boxdb {
		boxes = append(boxes, box)
	}
	return boxes
}

func (ctx *Context) VMs() []Vm {
	vms := make([]Vm, 0, len(ctx.vmdb))
	for _, vm := range ctx.vmdb {
		vms = append(vms, vm)
	}
	return vms
}

func (ctx *Context) Disks() []Disk {
	disks := make([]Disk, 0, len(ctx.diskdb))
	for _, disk := range ctx.diskdb {
		disks = append(disks, disk)
	}
	return disks
}

// EOF
