package main

import (
	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	//gocfg "github.com/sbinet/go-config/config"
)

func alto_make_cmd_market() *commander.Commander {
	cmd := &commander.Commander{
		Name:  "market",
		Short: "interact with the marketplace",
		Commands: []*commander.Command{
			alto_make_cmd_market_ls(),
		},
		Flag: flag.NewFlagSet("alto-market", flag.ExitOnError),
	}
	return cmd
}

// EOF
