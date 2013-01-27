package main

import (
	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	//gocfg "github.com/sbinet/go-config/config"
)

func alto_make_cmd_pdisk() *commander.Commander {
	cmd := &commander.Commander{
		Name:  "pdisk",
		Short: "add/remove/list persistent disks",
		Commands: []*commander.Command{
			alto_make_cmd_pdisk_add(),
		},
		Flag: flag.NewFlagSet("alto-pdisk", flag.ExitOnError),
	}
	return cmd
}

// EOF
