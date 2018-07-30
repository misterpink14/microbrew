package main

import (
	"flag"
	"log"
)

const (
	configPath     = ".config/mbrew"
	configFilename = "config"
	templatesPath  = "./templates"
)

func main() {
	copyTemplates := flag.Bool("t", false, "Copy template files")
	flag.Parse()

	config := GetConfig(configPath, configFilename)

	if *copyTemplates {
		log.Println(config.Templates)
		// NewTemplates(templatesPath, config.Templates).CopyTemplates()
		log.Println("copy")
	}
}
