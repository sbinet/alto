package main

import (
	"fmt"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func alto_make_cmd_vm_ls() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_vm_ls,
		UsageLine: "ls [options]",
		Short:     "list VMs from the repository of VMs",
		Long: `
ls lists the owner's repository of VMs.

ex:
 $ alto vm ls
`,
		Flag: *flag.NewFlagSet("alto-vm-ls", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_vm_ls(cmd *commander.Command, args []string) {
	var err error
	n := "alto-" + cmd.Name()

	switch len(args) {
	case 0:
		// ok
	default:
		err = fmt.Errorf("%s: does not take any argument\n", n)
		handle_err(err)
	}

	quiet := cmd.Flag.Lookup("q").Value.Get().(bool)

	if !quiet {
		fmt.Printf("%s: listing VMs...\n", n)
	}

	vms := g_ctx.VMs()
	if len(vms) > 0 {
		for _, vm := range vms {
			fmt.Printf("%v\n", vm)
		}
	} else {
		fmt.Printf("no vm\n")
	}

	if !quiet {
		fmt.Printf("%s: listing VMs... [done]\n", n)
	}

	return
}

// EOF
