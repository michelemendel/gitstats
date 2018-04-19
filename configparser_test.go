package main

import "testing"

func TestIsDir(t *testing.T) {
	expected := true
	actual := homeDir()
	if !isDir(actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
