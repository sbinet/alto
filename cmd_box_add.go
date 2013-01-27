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
		UsageLine: "add [options] <box-name> <box-configuration>",
	Short:     "add a box (VM+pdisk) to the repository of boxes",
		Long: `
add adds a box (VM+pdisk) on StratusLab.

ex:
 $ alto box add archlinux-64b
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

	switch len(args) {
	case 0:
		// ok
	default:
		err = fmt.Errorf("%s: does not take any argument\n", n)
		handle_err(err)
	}

	//quiet := cmd.Flag.Lookup("q").Value.Get().(bool)

	return
}

// EOF
