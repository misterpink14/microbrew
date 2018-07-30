package main

import (
	"fmt"
	"os"
	"testing"
)

const (
	exampleConfigPath     = "github.com/misterpink14/microbrew/example/config"
	exampleConfigFileName = "config"

	bashProfile         = "bash_profile"
	expectedBashProfile = "~/.bash_profile"

	nVim         = "init.vim"
	expectedNVim = "~/.config/nvim/init.vim"

	vim         = "vimrc"
	expectedVim = "~/.vimrc"
)

func ValidateConfig(t *testing.T, given, expected string) {
	if given != expected {
		t.Errorf("%s is expected to be %s", given, expected)
	}
}

func TestGetConfig(t *testing.T) {
	config := GetConfig(
		fmt.Sprintf("%s/src/%s", os.Getenv("GOPATH"), exampleConfigPath),
		exampleConfigFileName)

	if len(config.Templates) != 3 {
		t.Error("Incorrrect number of templates", len(config.Templates))
	}

	for template, destination := range config.Templates {
		switch template {
		case bashProfile:
			ValidateConfig(t, destination, expectedBashProfile)
		case nVim:
			ValidateConfig(t, destination, expectedNVim)
		case vim:
			ValidateConfig(t, destination, expectedVim)
		}
	}
}
