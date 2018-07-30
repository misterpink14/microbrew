package util

import (
	"log"
	"os/user"
	"path/filepath"
)

// Index returns the index of a given string in a string array
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Includes returns a bool value determining the existence of a
// given string in a string array
func Includes(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// GetHome returns the user's home path.
// Any failure results in a Fatal log.
func GetHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

// ReplaceWithHome replaces any instance of "~/" with the path to
// the user's home directory
func ReplaceWithHome(path string) string {
	if path[:2] == "~/" {
		return filepath.Join(GetHome(), path[2:])
	}
	return path
}
