package main

import (
	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	//gocfg "github.com/sbinet/go-config/config"
)

func alto_make_cmd_disk() *commander.Commander {
	cmd := &commander.Commander{
		Name:  "disk",
		Short: "add/remove/list persistent disks",
		Commands: []*commander.Command{
			alto_make_cmd_disk_add(),
			alto_make_cmd_disk_ls(),
		},
		Flag: flag.NewFlagSet("alto-disk", flag.ExitOnError),
	}
	return cmd
}

// EOF
