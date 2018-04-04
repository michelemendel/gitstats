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
	gitstatus()
}

func flatten(statusMsg string) string {
	return strings.Replace(statusMsg, "\n", "", -1)
}

func isDirClean(statusMsg string) bool {
	return strings.Contains(flatten(statusMsg), "nothing to commit, working directory clean")
}

func gitstatus() {
	cmdOut := git("status")
	if isDirClean(cmdOut) {
		fmt.Println("OK")
	} else {
		printBody(cmdOut)
	}
}

func gitfetch() {
	git("fetch")
}

func git(cmd string) string {
	cmdOutBytes, err := exec.Command("git", cmd).Output()
	checkErr(err)
	cmdOut := string(cmdOutBytes)
	printSubTitle(cmd)
	return cmdOut
}

func cwd() {
	cwd, _ := os.Getwd()
	printBody(cwd)
}

func printMainTitle(str string) {
	// fmt.Printf("\n/////////////%v\n", str)
	fmt.Printf("\n----------------------------------------------------------------------\n\t\t%v", str)
	fmt.Printf("\n----------------------------------------------------------------------\n")
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
