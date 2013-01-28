package altolib

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
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
	runtime.SetFinalizer(ctx, (*Context).sync_fs)
	return ctx, err
}

func (ctx *Context) init() error {
	var err error

	// load pdisks
	if !path_exists(DiskDbFileName) {
		var f *os.File
		f, err = os.Create(DiskDbFileName)
		defer f.Close()
		if err != nil {
			return err
		}
		disks := make([]Disk, 0)
		err = json.NewEncoder(f).Encode(&disks)
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
	disks, err := DiskList()
	if err != nil {
		return err
	}
	for _, disk := range disks {
		ctx.diskdb[disk.Guid] = disk
	}

	// load vms
	if !path_exists(VmDbFileName) {
		var f *os.File
		f, err = os.Create(VmDbFileName)
		defer f.Close()
		if err != nil {
			return err
		}
		vms := make([]Vm, 0)
		err = json.NewEncoder(f).Encode(&vms)
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

func (ctx *Context) Sync() error {
	// sync boxes
	{
		f, err := os.Create(BoxDbFileName)
		if err != nil {
			return err
		}
		defer f.Close()
		lst := ctx.Boxes()
		err = json.NewEncoder(f).Encode(&lst)
		if err != nil {
			return err
		}
		err = f.Sync()
		if err != nil {
			return err
		}
	}
	// sync vms
	{
		f, err := os.Create(VmDbFileName)
		if err != nil {
			return err
		}
		defer f.Close()
		lst := ctx.Vms()
		err = json.NewEncoder(f).Encode(&lst)
		if err != nil {
			return err
		}
		err = f.Sync()
		if err != nil {
			return err
		}
	}
	// sync pdisks
	{
		f, err := os.Create(DiskDbFileName)
		if err != nil {
			return err
		}
		defer f.Close()
		lst := ctx.Disks()
		err = json.NewEncoder(f).Encode(&lst)
		if err != nil {
			return err
		}
		err = f.Sync()
		if err != nil {
			return err
		}
	}
	return nil
}

// sync_fs synchronizes this context's data on disk
func (ctx *Context) sync_fs() {
	//fmt.Printf(">>> sync-fs...\n")
	err := ctx.Sync()
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf(">>> sync-fs... [done]\n")
}

func (ctx *Context) Boxes() []Box {
	boxes := make([]Box, 0, len(ctx.boxdb))
	for _, box := range ctx.boxdb {
		boxes = append(boxes, box)
	}
	return boxes
}

func (ctx *Context) Vms() []Vm {
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

func (ctx *Context) AddVm(name, id string) error {
	var err error
	// TODO: check id exists on the market!
	//  --> http://mp.stratuslab.eu/metadata/<id>
	_, ok := ctx.vmdb[id]
	if ok {
		return fmt.Errorf("altolib.context: VM with id=%s already in db", id)
	}
	ctx.vmdb[id] = Vm{Id: id, Tag: name}
	return err
}

func (ctx *Context) RemoveVm(id string) error {
	_, ok := ctx.vmdb[id]
	if !ok {
		return fmt.Errorf("altolib.context: no such VM [id=%s] in db", id)
	}
	delete(ctx.vmdb, id)
	return nil
}

// EOF
