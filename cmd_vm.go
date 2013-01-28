package main

import (
	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	//gocfg "github.com/sbinet/go-config/config"
)

func alto_make_cmd_vm() *commander.Commander {
	cmd := &commander.Commander{
		Name:  "vm",
		Short: "add/remove/list VMs",
		Commands: []*commander.Command{
			alto_make_cmd_vm_add(),
			alto_make_cmd_vm_ls(),
		},
		Flag: flag.NewFlagSet("alto-vm", flag.ExitOnError),
	}
	return cmd
}

// EOF
