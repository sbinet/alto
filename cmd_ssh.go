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

func alto_make_cmd_ssh() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_ssh,
		UsageLine: "ssh [options]",
		Short:     "connect to a (running) box on StratusLab",
		Long: `
ssh connects to (running) box (VM+pdisk) on StratusLab using the configuration from the current directory.

ex:
 $ alto ssh
`,
		Flag: *flag.NewFlagSet("alto-ssh", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_ssh(cmd *commander.Command, args []string) {
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
		fmt.Printf("%s: connecting...\n", n)
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

	ssh := exec.Command("stratus-connect-instance", string(id))
	ssh.Stdin = os.Stdin
	ssh.Stdout = os.Stdout
	ssh.Stderr = os.Stderr
	err = ssh.Run()
	handle_err(err)

	if !quiet {
		fmt.Printf("%s: connecting... [done]\n", n)
	}
	return
}

// EOF
