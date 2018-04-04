package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

const cfgFilenameDefault = "dirs.cfg"

func main() {
	// Get filename from CLI args
	cfgFilename := flag.String("f", cfgFilenameDefault, "a config file")
	flag.Parse()

	// Get cfgs as an array of structs
	cfgs := Parse(*cfgFilename)

	printTableOfCfgs(cfgs)
}

func printTableOfCfgs(cfgs Configs) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for i, cfg := range cfgs {
		fmt.Fprintf(w, "%v\t%v\t%v\n", i+1, cfg.Name, cfg.Dir)
	}
	w.Flush()
}
