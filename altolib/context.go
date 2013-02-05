package altolib

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

var (
	ConfigDirName = os.ExpandEnv("${HOME}/.config/alto")
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
		return nil, fmt.Errorf("altolib.init: %v", err)
	}
	runtime.SetFinalizer(ctx, (*Context).sync_fs)
	return ctx, err
}

func (ctx *Context) init() error {
	var err error

	if !path_exists(ConfigDirName) {
		err = os.MkdirAll(ConfigDirName, 0755)
		if err != nil {
			return fmt.Errorf(
				"altolib: could not make directory [%s] (%v)",
				ConfigDirName, err,
			)
		}
	}

	// load pdisks
	if !path_exists(DiskDbFileName) {
		var f *os.File
		f, err = os.Create(DiskDbFileName)
		if err != nil {
			return fmt.Errorf(
				"altolib: could not create file [%s] (%v)",
				DiskDbFileName, err,
			)
		}
		//defer f.Close()
		disks := make([]Disk, 0)
		err = json.NewEncoder(f).Encode(&disks)
		if err != nil {
			return fmt.Errorf(
				"altolib: could not encode into file [%s] (%v)",
				DiskDbFileName, err,
			)
		}
		err = f.Sync()
		if err != nil {
			return fmt.Errorf(
				"altolib: could not sync file [%s] (%v)",
				DiskDbFileName, err,
			)
		}
		err = f.Close()
		if err != nil {
			return fmt.Errorf(
				"altolib: could not close file [%s] (%v)",
				DiskDbFileName, err,
			)
		}
	}
	disks, err := DiskList()
	if err != nil {
		return fmt.Errorf(
			"altolib: could not fetch disk-list (%v)",
			err,
		)
	}
	for _, disk := range disks {
		ctx.diskdb[disk.Guid] = disk
	}

	// load vms
	if !path_exists(VmDbFileName) {
		var f *os.File
		f, err = os.Create(VmDbFileName)
		if err != nil {
			return fmt.Errorf(
				"altolib: could not create file [%s] (%v)",
				VmDbFileName, err,
			)
		}
		//defer f.Close()
		vms := make([]Vm, 0)
		err = json.NewEncoder(f).Encode(&vms)
		if err != nil {
			return fmt.Errorf(
				"altolib: could not encode into file [%s] (%v)",
				VmDbFileName, err,
			)
		}
		err = f.Sync()
		if err != nil {
			return fmt.Errorf(
				"altolib: could not sync file [%s] (%v)",
				VmDbFileName, err,
			)
		}
		err = f.Close()
		if err != nil {
			return fmt.Errorf(
				"altolib: could not close file [%s] (%v)",
				VmDbFileName, err,
			)
		}
	}
	vms, err := VmList()
	if err != nil {
		return fmt.Errorf(
			"altolib: could not fetch vm-list (%v)",
			err,
		)
	}
	for _, vm := range vms {
		ctx.vmdb[vm.Id] = vm
	}

	// load boxes
	if !path_exists(BoxDbFileName) {
		var f *os.File
		f, err = os.Create(BoxDbFileName)
		if err != nil {
			return fmt.Errorf(
				"altolib: could not create file [%s] (%v)",
				BoxDbFileName, err,
			)
		}
		//defer f.Close()
		boxes := make([]Box, 0)
		err = json.NewEncoder(f).Encode(&boxes)
		if err != nil {
			return fmt.Errorf(
				"altolib: could not encode into file [%s] (%v)",
				BoxDbFileName, err,
			)
		}
		err = f.Sync()
		if err != nil {
			return fmt.Errorf(
				"altolib: could not sync file [%s] (%v)",
				BoxDbFileName, err,
			)
		}
		err = f.Close()
		if err != nil {
			return fmt.Errorf(
				"altolib: could not close file [%s] (%v)",
				BoxDbFileName, err,
			)
		}
	}
	boxes, err := BoxList()
	if err != nil {
		if err != nil {
			return fmt.Errorf(
				"altolib: could not fetch box-list (%v)",
				err,
			)
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

func (ctx *Context) AddVm(vm Vm) error {
	var err error
	id := vm.Id
	// TODO: check id exists on the market!
	//  --> http://mp.stratuslab.eu/metadata/<id>
	_, ok := ctx.vmdb[id]
	if ok {
		return fmt.Errorf("altolib.context: VM with id=%s already in db", id)
	}
	ctx.vmdb[id] = vm
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

func (ctx *Context) GetVm(id string) (Vm, error) {
	var err error
	vm, ok := ctx.vmdb[id]
	if !ok {
		err = fmt.Errorf("altolib.context: no such VM [id=%s] in db", id)
		return Vm{}, err
	}
	return vm, nil
}

// func (ctx *Context) AddDisk(disk Disk) error {
// 	var err error
// 	id := disk.Id
// 	// TODO: check id exists on the market!
// 	//  --> http://mp.stratuslab.eu/metadata/<id>
// 	_, ok := ctx.diskdb[id]
// 	if ok {
// 		return fmt.Errorf("altolib.context: disk with id=%s already in db", id)
// 	}
// 	ctx.diskdb[id] = disk
// 	return err
// }

// func (ctx *Context) RemoveDisk(id string) error {
// 	_, ok := ctx.diskdb[id]
// 	if !ok {
// 		return fmt.Errorf("altolib.context: no such disk [id=%s] in db", id)
// 	}
// 	delete(ctx.diskdb, id)
// 	return nil
// }

func (ctx *Context) GetDisk(id string) (Disk, error) {
	var err error
	disk, ok := ctx.diskdb[id]
	if !ok {
		err = fmt.Errorf("altolib.context: no such disk [id=%s] in db", id)
		return Disk{}, err
	}
	return disk, nil
}

func (ctx *Context) AddBox(box Box) error {
	var err error
	id := box.Id
	_, ok := ctx.boxdb[id]
	if ok {
		return fmt.Errorf("altolib.context: box with id=%s already in db", id)
	}
	// TODO: check box.Vm and box.Disk exist in dbs as well
	ctx.boxdb[id] = box
	return err
}

func (ctx *Context) RemoveBox(id string) error {
	_, ok := ctx.boxdb[id]
	if !ok {
		return fmt.Errorf("altolib.context: no such box [id=%s] in db", id)
	}
	delete(ctx.boxdb, id)
	return nil
}

func (ctx *Context) GetBox(id string) (Box, error) {
	var err error
	box, ok := ctx.boxdb[id]
	if !ok {
		err = fmt.Errorf("altolib.context: no such box [id=%s] in db", id)
		return Box{}, err
	}
	return box, nil
}

// EOF
