package main

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	// "os"
	// "os/exec"
	// "path/filepath"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	"github.com/sbinet/alto/altolib"
)

func alto_make_cmd_box_add() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_box_add,
		UsageLine: "add [options] <box-name> <vm-name> [<pdisk-name>]",
		Short:     "add a box (VM+pdisk) to the repository of boxes",
		Long: `
add adds a box (VM+pdisk) on StratusLab.

ex:
 $ alto box add archlinux-64b my-archlinux-vm my-archlinux-disk
`,
		Flag: *flag.NewFlagSet("alto-box-add", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_box_add(cmd *commander.Command, args []string) {
	var err error
	n := "alto-" + cmd.Name()

	box_id := ""
	vm_id := ""
	disk_id := ""
	switch len(args) {
	case 2:
		box_id = args[0]
		vm_id = args[1]
	case 3:
		box_id = args[0]
		vm_id = args[1]
		disk_id = args[2]
	default:
		err = fmt.Errorf("%s: takes at least 2 arguments\n", n)
		handle_err(err)
	}

	quiet := cmd.Flag.Lookup("q").Value.Get().(bool)

	if !quiet {
		fmt.Printf("%s: adding new box [%s]...\n", n, box_id)
	}

	vms := g_ctx.Vms()
	found_vm := false
	for _, vm := range vms {
		switch vm_id {
		case vm.Id:
			vm_id = vm.Id
			found_vm = true
			break
		case vm.Tag:
			vm_id = vm.Id
			found_vm = true
			break
		}
	}
	if !found_vm {
		err = fmt.Errorf("%s: no such VM %q in db", n, vm_id)
		handle_err(err)
	}

	vm, err := g_ctx.GetVm(vm_id)
	handle_err(err)

	var disk altolib.Disk
	if disk_id != "" {
		found_disk := false
		disks := g_ctx.Disks()
		for _, disk := range disks {
			switch disk_id {
			case disk.Guid:
				disk_id = disk.Guid
				found_disk = true
				break
			case disk.Tag:
				disk_id = disk.Guid
				found_disk = true
				break
			}
		}
		if !found_disk {
			err = fmt.Errorf("%s: no such disk %q in db", n, disk_id)
			handle_err(err)
		}
		disk, err = g_ctx.GetDisk(disk_id)
		handle_err(err)
	}

	err = g_ctx.AddBox(altolib.Box{
		Id:   box_id,
		Vm:   vm,
		Disk: disk,
	})
	handle_err(err)
	if !quiet {
		fmt.Printf("%s: adding new box [%s]... [done]\n", n, box_id)
	}
	return
}

// EOF
