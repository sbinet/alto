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

func alto_make_cmd_vm_add() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_vm_add,
		UsageLine: "add [options] <vm-name> <vm-id>",
		Short:     "add a VM to the repository of VMs",
		Long: `
add adds a VM on StratusLab.

ex:
 $ alto vm add archlinux-64b
`,
		Flag: *flag.NewFlagSet("alto-vm-add", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_vm_add(cmd *commander.Command, args []string) {
	var err error
	n := "alto-" + cmd.Name()

	vm_name := ""
	vm_id := ""
	switch len(args) {
	case 2:
		vm_name = args[0]
		vm_id = args[1]
	default:
		err = fmt.Errorf("%s: does not take any argument\n", n)
		handle_err(err)
	}

	quiet := cmd.Flag.Lookup("q").Value.Get().(bool)
	if !quiet {
		fmt.Printf("%s: adding VM [%s] to repository...\n", n, vm_name)
	}

	err = g_ctx.AddVm(vm_name, vm_id)
	handle_err(err)

	if !quiet {
		fmt.Printf("%s: adding VM [%s] to repository... [done]\n", n, vm_name)
	}
	return
}

// EOF
