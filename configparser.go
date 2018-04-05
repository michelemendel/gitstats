package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Config is a single configuration.
type Config struct {
	Name string
	Dir  string
}

// Configs is a list of configurations.
type Configs []Config

// GetConfigs returns a list of configurations.
func GetConfigs(filename string) Configs {
	var cfgs Configs
	file := getFile(filename)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fs := getFields(scanner)

		if len(fs) == 2 && isDir(fs[1]) {
			cfgs = append(cfgs, Config{fs[0], fs[1]})
		} else if len(fs) > 2 || len(fs) == 1 { // There may be empty arrays, but they are ignored.
			printWarning("Entry is not in correct key-value format", strings.Join(fs, " "))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading standard input:", err)
	}

	return cfgs
}

func getFile(filename string) *os.File {
	file, err := os.Open(filename)
	errFatal(err)
	return file
}

func getFields(scanner *bufio.Scanner) []string {
	line := strings.Trim(scanner.Text(), " ")

	// Do not add comments
	if strings.HasPrefix(line, "#") {
		return []string{}
	}

	return strings.Fields(line)
}

// If dir doesn't exists, don't add it
func isDir(dirname string) bool {
	if _, err := os.Stat(dirname); err != nil {
		printWarning("Directory doesn't exists", dirname)
		return false
	}
	return true
}

func printWarning(msg string, obj string) {
	fmt.Printf("\n%vWARNING: %v (%v).%v", Red, msg, obj, AttrOff)
}
