package main

import (
	"log"
	"testing"
)

func TestDependency_IsInstalled_true(t *testing.T) {
	installCommand := []string{"blah"}
	updateCommand := []string{"blah"}
	dependency := NewDependency("python", installCommand, updateCommand)
	if installed, err := dependency.IsInstalled(); !installed || err != nil {
		t.Error("python should be installed")
	}
}

func TestDependency_IsInstalled_false(t *testing.T) {
	installCommand := []string{"blah"}
	updateCommand := []string{"blah"}
	dependency := NewDependency("notinstalled", installCommand, updateCommand)
	if installed, err := dependency.IsInstalled(); installed || err == nil {
		log.Println(err)
		t.Error("notinstalled should not be installed")
	}
}
