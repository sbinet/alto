/*
alto manages stratuslab VMs and disks.
*/
package main

import (
	"os"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

var g_cmd *commander.Commander

func init() {
	g_cmd = &commander.Commander{
		Name: os.Args[0],
		Commands: []*commander.Command{
			alto_make_cmd_init(),
			alto_make_cmd_up(),
		},
		Flag: flag.NewFlagSet("alto", flag.ExitOnError),
		Commanders: []*commander.Commander{
			alto_make_cmd_box(),
		},
	}
}

func main() {

	var err error
	err = g_cmd.Flag.Parse(os.Args[1:])
	handle_err(err)

	args := g_cmd.Flag.Args()
	err = g_cmd.Run(args)
	handle_err(err)

	return
}
