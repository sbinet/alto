package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	// "path/filepath"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	"github.com/sbinet/alto/altolib"
)

func alto_make_cmd_up() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_up,
		UsageLine: "up [options]",
		Short:     "launch a box on StratusLab",
		Long: `
up launches a box (VM+pdisk) on StratusLab using the configuration from the current directory.

ex:
 $ alto up
`,
		Flag: *flag.NewFlagSet("alto-up", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_up(cmd *commander.Command, args []string) {
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
		fmt.Printf("%s: starting...\n", n)
	}

	const fname = "AltoFile"
	if !path_exists(fname) {
		err = fmt.Errorf("%s: no such file [%s]. did you run 'alto init some-box-name' ?", n, fname)
		handle_err(err)
	}

	f, err := os.Open(fname)
	handle_err(err)
	defer f.Close()

	var box altolib.Box
	err = json.NewDecoder(f).Decode(&box)
	handle_err(err)

	const id_fname = ".alto.id"
	if path_exists(id_fname) {
		data, err := ioutil.ReadFile(id_fname)
		handle_err(err)

		id := string(bytes.Trim(data, " \r\n"))

		err = fmt.Errorf("%s: the box [%s] has already been instantiated (run-id=%s)", n, box.Id, id)
		handle_err(err)
	}

	slab_args := []string{
		fmt.Sprintf("--ram=%d", box.Ram),
		fmt.Sprintf("--cpu=%d", box.Cpus),
	}
	if box.Disk.Guid != "" {
		slab_args = append(
			slab_args,
			fmt.Sprintf("--persistent-disk=%s", box.Disk.Guid),
		)
	}
	slab_args = append(slab_args,
		"--output="+id_fname,
		box.Vm.Id,
	)
	slab := exec.Command("stratus-run-instance", slab_args...)
	slab.Stdin = os.Stdin
	slab.Stdout = os.Stdout
	slab.Stderr = os.Stderr
	err = slab.Run()
	handle_err(err)

	if !quiet {
		fmt.Printf("%s: starting... [done]\n", n)
	}
	return
}

// EOF
