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

	box := ""
	vm := ""
	disk := ""
	switch len(args) {
	case 2:
		box = args[0]
		vm = args[1]
	case 3:
		box = args[0]
		vm = args[1]
		disk = args[2]
	default:
		err = fmt.Errorf("%s: takes at least 2 arguments\n", n)
		handle_err(err)
	}

	quiet := cmd.Flag.Lookup("q").Value.Get().(bool)

	if !quiet {
		fmt.Printf("%s: adding new box [%s]...\n", n, box)
	}

	fmt.Printf(">>> box=%q vm=%q disk=%q\n", box, vm, disk)
	if !quiet {
		fmt.Printf("%s: adding new box [%s]... [done]\n", n, box)
	}
	return
}

// EOF
