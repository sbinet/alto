/*
alto manages stratuslab VMs and disks.
*/
package main

import (
	"os"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	"github.com/sbinet/alto/altolib"
)

var g_cmd *commander.Commander
var g_ctx *altolib.Context

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
			alto_make_cmd_pdisk(),
			alto_make_cmd_vm(),
		},
	}
}

func main() {

	var err error
	err = g_cmd.Flag.Parse(os.Args[1:])
	handle_err(err)

	g_ctx, err = altolib.NewContext()
	handle_err(err)
	defer func() {
		err = g_ctx.Sync()
		handle_err(err)
	}()

	args := g_cmd.Flag.Args()
	err = g_cmd.Run(args)
	handle_err(err)

	return
}
