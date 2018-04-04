package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"
)

const cfgFilenameDefault = "dirs.cfg"

func main() {
	cfgs := parseConfig()
	// printTableOfCfgs(cfgs)

	for _, cfg := range cfgs {
		presentGitStatus(cfg)
	}
}

func parseConfig() Configs {
	// Get filename from CLI args
	cfgFilename := flag.String("f", cfgFilenameDefault, "a config file")
	flag.Parse()

	// Get cfgs as an array of structs
	cfgs := Parse(*cfgFilename)

	return cfgs
}

func printTableOfCfgs(cfgs Configs) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for i, cfg := range cfgs {
		fmt.Fprintf(w, "%v\t%v\t%v\n", i+1, cfg.Name, cfg.Dir)
	}
	w.Flush()
}

func presentGitStatus(cfg Config) {
	os.Chdir(cfg.Dir)
	printMainTitle(cfg.Name)
	printSubTitle(cfg.Dir)
	gitfetch()
	res := gitstatus()
	fmt.Printf("\n###\n%v\n###\n", res)
}

func gitstatus() string {
	return git("status")
}

func gitfetch() string {
	return git("fetch")
}

func git(cmd string) string {
	cmdOutBytes, err := exec.Command("git", cmd).Output()
	checkErr(err)
	cmdOut := string(cmdOutBytes)
	printSubTitle(cmd)
	printBody(cmdOut)
	return cmdOut
}

func cwd() {
	cwd, _ := os.Getwd()
	printBody(cwd)
}

func printMainTitle(str string) {
	// fmt.Printf("\n/////////////%v\n", str)
	fmt.Printf("\n--------------------------------------------------\n%v", str)
	fmt.Printf("\n--------------------------------------------------\n")
}

func printSubTitle(str string) {
	fmt.Printf("---> %v\n", strings.ToUpper(str))
}

func printBody(str string) {
	fmt.Printf("%v\n", str)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
