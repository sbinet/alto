package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"os"
	// "os/exec"
	// "path/filepath"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func alto_make_cmd_init() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_init,
		UsageLine: "init [options] <box-name>",
		Short:     "create a box on StratusLab",
		Long: `
init prepares a directory for hosting a box (VM+pdisk) on StratusLab.

ex:
 $ alto init archlinux-64b
`,
		Flag: *flag.NewFlagSet("alto-init", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_init(cmd *commander.Command, args []string) {
	var err error
	n := "alto-" + cmd.Name()

	box_id := ""
	switch len(args) {
	case 1:
		box_id = args[0]
	default:
		err = fmt.Errorf("%s: need a box id\n", n)
		handle_err(err)
	}

	quiet := cmd.Flag.Lookup("q").Value.Get().(bool)
	if !quiet {
		fmt.Printf("%s with box=%s...\n", n, box_id)
	}

	box, err := g_ctx.GetBox(box_id)
	handle_err(err)

	const fname = "AltoFile"
	if path_exists(fname) {
		err = fmt.Errorf("%s: a file [%s] already exists", n, fname)
		handle_err(err)
	}

	f, err := os.Create(fname)
	handle_err(err)
	defer f.Close()

	err = json.NewEncoder(f).Encode(&box)
	handle_err(err)
	handle_err(f.Sync())
	handle_err(f.Close())

	if !quiet {
		fmt.Printf("%s with box=%s... [done]\n", n, box_id)
	}
	return
}

// EOF
