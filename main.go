package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"
	"sync"
	"text/tabwriter"
)

type clr struct {
	c string
}

const cfgFilenameDefault = "gitststatsdirs.cfg"

func main() {
	cfgFilename := getCfgFilename()
	cfgs := GetConfigs(cfgFilename)
	// printTableOfCfgs(cfgs)
	showGitStatuses(cfgs)
}

// Get filename from CLI args, unless you want to use default in home directory.
func getCfgFilename() string {
	filename := flag.String("f", path.Join(homeDir(), cfgFilenameDefault),
		"You need to provide a config file with -f or empty which will look for default.")
	flag.Parse()
	return *filename
}

func showTableOfCfgs(cfgs Configs) {
	fmt.Println("")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for i, cfg := range cfgs {
		fmt.Fprintf(w, "%v\t%v\t%v\n", i+1, cfg.Name, cfg.Dir)
	}
	w.Flush()
}

var wg sync.WaitGroup

func showGitStatuses(cfgs Configs) {
	for _, cfg := range cfgs {
		wg.Add(1)
		go func(cfg Config) {
			gitstatus(cfg)
			wg.Done()
		}(cfg)
	}
	wg.Wait()
}

func gitstatus(cfg Config) {
	git("fetch", cfg.Dir)
	cmdOut := git("status", cfg.Dir)

	printMainTitle(cfg.Name)
	fmt.Printf("%v\n", cfg.Dir)

	if isDirClean(cmdOut) {
		fmt.Printf("OK\n\n")
	} else {
		fmt.Printf("%v\n", cmdOut)
	}
}

func git(cmd string, dir string) string {
	cmdOutBytes, err := exec.Command("git", "-C", dir, cmd).Output()
	errFatal(err)
	return string(cmdOutBytes)
}

func isDirClean(statusMsg string) bool {
	sm := flatten(statusMsg)

	return strings.Contains(sm, "nothing to commit") &&
		strings.Contains(sm, "working") &&
		strings.Contains(sm, "clean")
}

func flatten(statusMsg string) string {
	return strings.Replace(statusMsg, "\n", "", -1)
}

func errFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homeDir() string {
	user, err := user.Current()
	errFatal(err)
	return user.HomeDir
}

func printMainTitle(str string) {
	const w80 = "--------------------------------------------------------------------------------"
	fmt.Printf("%v%v\n%v\n%v\n%v", Green, w80, str, w80, AttrOff)
	// fmt.Printf("\n%v\n%v", w80, Nocolor)
}
