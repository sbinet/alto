package main

import (
	"fmt"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func alto_make_cmd_box_ls() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_box_ls,
		UsageLine: "ls [options]",
		Short:     "list boxes from the repository of boxes",
		Long: `
ls lists the owner's repository of boxes.

ex:
 $ alto box ls
`,
		Flag: *flag.NewFlagSet("alto-box-ls", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_box_ls(cmd *commander.Command, args []string) {
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
		fmt.Printf("%s: listing boxes...\n", n)
	}

	boxes := g_ctx.Boxes()
	for _, box := range boxes {
		const indent = "    "
		fmt.Printf(
			"::: box [%s]\n%s%v\n%s%v\n",
			box.Id,
			indent, box.Vm,
			indent, box.Disk,
		)
	}

	if !quiet {
		fmt.Printf("%s: listing boxes... [done]\n", n)
	}

	return
}

// EOF
