package main

import (
	"bytes"
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	// "path/filepath"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func alto_make_cmd_down() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_down,
		UsageLine: "down [options]",
		Short:     "shutdown a (running) box on StratusLab",
		Long: `
down sends the shutdown signal to a (running) box (VM+pdisk) on StratusLab using the configuration from the current directory.

ex:
 $ alto down
`,
		Flag: *flag.NewFlagSet("alto-down", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	cmd.Flag.Bool("kill", true, "kill the box")
	return cmd
}

func alto_run_cmd_down(cmd *commander.Command, args []string) {
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
	do_kill := cmd.Flag.Lookup("kill").Value.Get().(bool)

	if !quiet {
		fmt.Printf("%s: shutting down...\n", n)
	}

	const cfg_fname = "AltoFile"
	if !path_exists(cfg_fname) {
		err = fmt.Errorf("%s: no such file [%s]. did you run 'alto init some-box-name' ?", n, cfg_fname)
		handle_err(err)
	}

	const id_fname = ".alto.id"
	if !path_exists(id_fname) {
		err = fmt.Errorf("%s: no such file [%s]. did you run 'alto up' ?", n, id_fname)
		handle_err(err)
	}

	data, err := ioutil.ReadFile(id_fname)
	handle_err(err)

	id := bytes.Trim(data, " \r\n")

	cmd_name := "stratus-shutdown-instance"
	if do_kill {
		cmd_name = "stratus-kill-instance"
	}

	ssh := exec.Command(cmd_name, string(id))
	ssh.Stdin = os.Stdin
	ssh.Stdout = os.Stdout
	ssh.Stderr = os.Stderr
	err = ssh.Run()
	handle_err(err)

	err = os.Remove(id_fname)
	handle_err(err)

	if !quiet {
		fmt.Printf("%s: shutting down... [done]\n", n)
	}
	return
}

// EOF
