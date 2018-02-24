package main

import (
	"os/user"
	"testing"
)

func TestUtils_Index(t *testing.T) {
	a := []string{"val1", "val2"}

	if Index(a, "val3") != -1 {
		t.Error("Expecting to get -1 as index for values not in array")
	}
	if Index(a, "val1") != 0 || Index(a, "val2") != 1 {
		t.Error("Expecting array to have val1 at index 0 and val2 at index 1")
	}
}

func TestUtils_Includes(t *testing.T) {
	a := []string{"val1", "val2"}

	if Includes(a, "val3") {
		t.Error("Expecting Includes to return false for values not in array")
	}
	if !Includes(a, "val1") || !Includes(a, "val2") {
		t.Error("Expecting array to include val1 and val2")
	}
}

func TestUtils_ReplaceWithHome(t *testing.T) {
	path := "~/path"
	usr, _ := user.Current()
	expectedPath := usr.HomeDir + "/path"

	if ReplaceWithHome(path) != expectedPath {
		t.Error("User home should have been replaced")
	}
}
