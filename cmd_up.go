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

func alto_make_cmd_up() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_up,
		UsageLine: "up [options]",
		Short:     "launch a box on StratusLab",
		Long: `
up launches a box (VM+pdisk) on StratusLab using the configuration from the current directory.

ex:
 $ alto up
`,
		Flag: *flag.NewFlagSet("alto-up", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_up(cmd *commander.Command, args []string) {
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