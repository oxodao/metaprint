package main

import (
	"fmt"
	"os"

	"github.com/oxodao/metaprint/config"
	"github.com/oxodao/metaprint/pulse"
)

const (
	AUTHOR        = "Oxodao"
	VERSION       = "0.3"
	SOFTWARE_NAME = "metaprint"
)

var otherCommands = []string{"pulseaudio-infos"}

func main() {
	cfg := config.Load()

	hasOtherCommands := hasOtherCommand()
	if len(os.Args) < 3 && !hasOtherCommands {
		printUsage()
		os.Exit(1)
	}

	if hasOtherCommands {
		switch os.Args[1] {
		case "pulseaudio-infos":
			pulse.PrintInfos()
		default:
			printUsage()
			os.Exit(1)
		}
		return
	}

	module, err := cfg.FindModule(os.Args[1], os.Args[2])

	if module == nil && err == nil {
		printUsage()
		return
	} else if module == nil {
		fmt.Println(err)
		os.Exit(1)
	}

	text := module.Print(os.Args[3:])

	if len(module.GetPrefix()) > 0 {
		text = module.GetPrefix() + " " + text
	}

	if len(module.GetSuffix()) > 0 {
		text += " " + module.GetSuffix()
	}

	fmt.Println(text)
}

func printUsage() {
	fmt.Println("Usage: metaprint <module> <name> [params]")
	fmt.Println("       metaprint other-command")
	fmt.Println()
	fmt.Println("Other commands: ")
	for _, cmd := range otherCommands {
		fmt.Println("\t- " + cmd)
	}
	fmt.Println()
	fmt.Println("Available modules: ")
	for _, mod := range config.GetModulesAvailable() {
		fmt.Println("\t- " + mod)
	}
}

func hasOtherCommand() bool {
	if len(os.Args) != 2 {
		return false
	}

	for _, val := range otherCommands {
		if os.Args[1] == val {
			return true
		}
	}

	return false
}
