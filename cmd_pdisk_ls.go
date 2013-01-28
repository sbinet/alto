package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"os"
	// "os/exec"
	"path/filepath"
	// "regexp"
	// "strings"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	"github.com/sbinet/alto/altolib"
)

func alto_make_cmd_pdisk_ls() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_pdisk_ls,
		UsageLine: "ls [options]",
		Short:     "list pdisks from the repository of persistent disks",
		Long: `
ls lists the owner's repository of persistent disks from StratusLab.

ex:
 $ alto pdisk ls
`,
		Flag: *flag.NewFlagSet("alto-pdisk-ls", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_pdisk_ls(cmd *commander.Command, args []string) {
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
		fmt.Printf("%s: listing pdisks...\n", n)
	}

	pdisks := g_ctx.Disks()
	for _, pdisk := range pdisks {
		fmt.Printf("%v\n", pdisk)
	}

	if !quiet {
		fmt.Printf("%s: listing pdisks... [done]\n", n)
	}

	// refresh cache of pdisks...
	fname := altolib.DiskDbFileName
	if path_exists(fname) {
		err = os.RemoveAll(fname)
		handle_err(err)
	}
	err = os.MkdirAll(filepath.Dir(fname), 0755)
	handle_err(err)
	f, err := os.Create(fname)
	handle_err(err)
	defer func() {
		err = f.Sync()
		handle_err(err)
		err = f.Close()
		handle_err(err)
	}()
	err = json.NewEncoder(f).Encode(&pdisks)
	handle_err(err)
	return
}

// EOF
