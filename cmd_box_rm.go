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

func alto_make_cmd_box_rm() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_box_rm,
		UsageLine: "rm [options] <box-name>",
		Short:     "rm a box from the repository of boxes",
		Long: `
rm removes a box from the alto repository.

ex:
 $ alto box rm archlinux-64b
`,
		Flag: *flag.NewFlagSet("alto-box-rm", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_box_rm(cmd *commander.Command, args []string) {
	var err error
	n := "alto-" + cmd.Name()

	box_id := ""
	switch len(args) {
	case 1:
		box_id = args[0]
	default:
		err = fmt.Errorf("%s: need the name of the box to remove\n", n)
		handle_err(err)
	}

	quiet := cmd.Flag.Lookup("q").Value.Get().(bool)
	if !quiet {
		fmt.Printf("%s: remove box [%s] from repository...\n", n, box_id)
	}

	found := false
	boxes := g_ctx.Boxes()
	for _, box := range boxes {
		switch box_id {
		case box.Id:
			found = true
			err = g_ctx.RemoveBox(box.Id)
			handle_err(err)
			break
		}
	}
	if !found {
		err = fmt.Errorf("%s: could not find the box [%s] in repository", n, box_id)
		handle_err(err)
	}

	if !quiet {
		fmt.Printf("%s: remove VM [%s] from repository... [done]\n", n, box_id)
	}
	return
}

// EOF
