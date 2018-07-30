package dependency

import (
	"os/exec"
)

// type IDependency interface {
// IsInstalled() (bool, error)
// Install() error
// Update() error
// }

type Dependency struct {
	Name           string
	InstallCommand []string
	UpdateCommand  []string
}

func NewDependency(name string, installCommand []string, updateCommand []string) Dependency {
	return Dependency{
		Name:           name,
		InstallCommand: installCommand,
		UpdateCommand:  updateCommand,
	}
}

func (d *Dependency) execute(cmd []string) error {
	first, rest := cmd[0], cmd[1:]
	return exec.Command(first, rest...).Run()
}

func (d *Dependency) Install() error {
	return d.execute(d.InstallCommand)
}

func (d *Dependency) Update() error {
	return d.execute(d.UpdateCommand)
}

func (d *Dependency) IsInstalled() (bool, error) {
	cmdOut, err := exec.Command("which", d.Name).Output()
	if err != nil || string(cmdOut) == "" {
		return false, err
	}
	return true, nil
}
