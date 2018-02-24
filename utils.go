package main

import (
	"os/user"
	"path/filepath"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Includes(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func ReplaceWithHome(path string) string {
	usr, _ := user.Current()
	if path[:2] == "~/" {
		return filepath.Join(usr.HomeDir, path[2:])
	}
	return path
}
