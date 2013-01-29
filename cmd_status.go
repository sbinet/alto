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

func alto_make_cmd_status() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_status,
		UsageLine: "status [options]",
		Short:     "display the status of a box on StratusLab",
		Long: `
status displays the status of a box on StratusLab.

ex:
 $ alto status
`,
		Flag: *flag.NewFlagSet("alto-status", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	cmd.Flag.Bool("all", false, "display the status of all boxes (not just the current one)")
	return cmd
}

func alto_run_cmd_status(cmd *commander.Command, args []string) {
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
	all_boxes := cmd.Flag.Lookup("all").Value.Get().(bool)

	if !quiet {
		fmt.Printf("%s...\n", n)
	}

	cmd_name := "stratus-describe-instance"
	id := ""
	if !all_boxes {
		const id_fname = ".alto.id"
		if !path_exists(id_fname) {
			err = fmt.Errorf("%s: no such file [%s]. did you run 'alto up' ?", n, id_fname)
			handle_err(err)
		}

		data, err := ioutil.ReadFile(id_fname)
		handle_err(err)

		id = string(bytes.Trim(data, " \r\n"))
	}

	ssh_args := []string{}
	if id != "" {
		ssh_args = append(ssh_args, id)
	}
	ssh := exec.Command(cmd_name, ssh_args...)
	ssh.Stdin = os.Stdin
	ssh.Stdout = os.Stdout
	ssh.Stderr = os.Stderr
	err = ssh.Run()
	handle_err(err)

	if !quiet {
		fmt.Printf("%s:... [done]\n", n)
	}
	return
}

// EOF
