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
)

func alto_make_cmd_vm_rm() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_vm_rm,
		UsageLine: "rm [options] <vm-name|vm-id>",
		Short:     "rm a VM from the repository of VMs",
		Long: `
rm removes a VM from the alto repository.

ex:
 $ alto vm rm archlinux-64b
 $ alto vm rm AVijhrWdomWxMT5V34bqEPvArCB
`,
		Flag: *flag.NewFlagSet("alto-vm-rm", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_vm_rm(cmd *commander.Command, args []string) {
	var err error
	n := "alto-" + cmd.Name()

	vm_id := ""
	switch len(args) {
	case 1:
		vm_id = args[0]
	default:
		err = fmt.Errorf("%s: takes one argument (vm-name or vm-id)\n", n)
		handle_err(err)
	}

	quiet := cmd.Flag.Lookup("q").Value.Get().(bool)
	if !quiet {
		fmt.Printf("%s: remove VM [%s] from repository...\n", n, vm_id)
	}

	found := false
	vms := g_ctx.Vms()
	for _, vm := range vms {
		switch vm_id {
		case vm.Tag:
			found = true
			err = g_ctx.RemoveVm(vm.Id)
			handle_err(err)
			break
		case vm.Id:
			found = true
			err = g_ctx.RemoveVm(vm.Id)
			handle_err(err)
			break
		}
	}
	if !found {
		err = fmt.Errorf("%s: could not find the VM [%s] in repository", n, vm_id)
		handle_err(err)
	}

	if !quiet {
		fmt.Printf("%s: remove VM [%s] from repository... [done]\n", n, vm_id)
	}
	return
}

// EOF
