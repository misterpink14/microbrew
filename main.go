package main

import (
	"flag"
)

const TemplatesPath = "./templates"

type FlagValues struct {
	IsCopyTemplates bool
}

func parseFlags() FlagValues {
	isCopyTemplates := flag.Bool("t", false, "Copy template files")
	flag.Parse()
	return FlagValues{
		IsCopyTemplates: *isCopyTemplates,
	}
}

func main() {
	flags := parseFlags()
	config := GetConfig()

	if flags.IsCopyTemplates {
		NewTemplates(TemplatesPath, config.Templates).CopyTemplates()
	}
}
