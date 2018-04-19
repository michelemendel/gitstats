package main

import "testing"

// Benchmark the default config file on
func Benchmark(b *testing.B) {
	cfgFilename := getCfgFilename()
	cfgs := GetConfigs(cfgFilename)

	for i := 0; i < b.N; i++ {
		showGitStatuses(cfgs)
	}
}
