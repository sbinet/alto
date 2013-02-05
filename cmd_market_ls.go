package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func alto_make_cmd_market_ls() *commander.Command {
	cmd := &commander.Command{
		Run:       alto_run_cmd_market_ls,
		UsageLine: "ls [options]",
		Short:     "list VMs from the repository of VMs",
		Long: `
ls lists the owner's repository of VMs.

ex:
 $ alto market ls
`,
		Flag: *flag.NewFlagSet("alto-market-ls", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func alto_run_cmd_market_ls(cmd *commander.Command, args []string) {
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
		fmt.Printf("%s: listing...\n", n)
	}

	resp, err := http.Get("https://marketplace.stratuslab.eu:443/metadata")
	handle_err(err)
	defer resp.Body.Close()

	type md_descr_t struct {
		Identifier string `xml:"dcterms:identifier"`
	}
	type MarketMetadata struct {
		XMLName xml.Name      `xml:"metadata"`
		Descr   []*md_descr_t `xml:"rdf:Description"`
	}

	md := MarketMetadata{}
	err = xml.NewDecoder(resp.Body).Decode(&md)
	handle_err(err)
	fmt.Printf(">>> %v\n", md)

	if !quiet {
		fmt.Printf("%s: listing... [done]\n", n)
	}

	return
}

// EOF
