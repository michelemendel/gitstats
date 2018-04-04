package main

import (
	"bufio"
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

// Parse returns a list of configurations.
func Parse(filename string) Configs {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	var cfgs Configs
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())

		// There may be empty arrays.
		if len(s) == 2 {
			cfgs = append(cfgs, Config{s[0], s[1]})
		} else if len(s) > 2 || len(s) == 1 {
			log.Fatal("Error in config file. Check that each line has the key-value format.")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading standard input:", err)
	}

	return cfgs

}
