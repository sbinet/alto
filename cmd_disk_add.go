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

func alto_make_cmd_disk_add() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_disk_add,
		UsageLine: "add [options] <disk-name> <disk-configuration>",
		Short:     "add a disk to the repository of persistent disks",
		Long: `
add adds a persistent disk on StratusLab.

ex:
 $ alto disk add archlinux-64b-data
`,
		Flag: *flag.NewFlagSet("alto-disk-add", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_disk_add(cmd *commander.Command, args []string) {
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
