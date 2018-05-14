package main

import (
	"testing"
)

const (
	BashProfile         = "bash_profile"
	ExpectedBashProfile = "~/.bash_profile"

	NVim         = "init.vim"
	ExpectedNVim = "~/.config/nvim/init.vim"

	Vim         = "vimrc"
	ExpectedVim = "~/.vimrc"
)

func ValidateConfig(t *testing.T, given, expected string) {
	if given != expected {
		t.Errorf("%s is expected to be %s", given, expected)
	}
}

func TestGetConfig(t *testing.T) {
	config := GetConfig()

	if len(config.Templates) != 3 {
		t.Error("Incorrrect number of templates", len(config.Templates))
	}

	for template, destination := range config.Templates {
		switch template {
		case BashProfile:
			ValidateConfig(t, destination, ExpectedBashProfile)
		case NVim:
			ValidateConfig(t, destination, ExpectedNVim)
		case Vim:
			ValidateConfig(t, destination, ExpectedVim)
		}
	}
}
