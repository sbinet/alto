package main

import (
	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	//gocfg "github.com/sbinet/go-config/config"
)

func alto_make_cmd_box() *commander.Commander {
	cmd := &commander.Commander{
		Name:  "box",
		Short: "add/remove/edit boxes",
		Commands: []*commander.Command{
			alto_make_cmd_box_add(),
			alto_make_cmd_box_ls(),
		},
		Flag: flag.NewFlagSet("alto-box", flag.ExitOnError),
	}
	return cmd
}

// EOF
