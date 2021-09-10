package main

import (
	"fmt"
	"github.com/oxodao/metaprint/config"
	"github.com/oxodao/metaprint/modules"
	"github.com/oxodao/metaprint/pulse"
	"os"
	"strings"
)

const (
	AUTHOR        = "Oxodao"
	VERSION       = "0.2"
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

	// @TODO: Make this not ugly, like find something to put modules in map but make things work correctly and iterate on it
	// => This could also be used to autogenerate the printUsage thing
	var module modules.Module = nil
	var ok bool = false

	switch os.Args[1] {
	case "battery":
		module, ok = cfg.Battery[os.Args[2]]
	case "date":
		module, ok = cfg.DateTime[os.Args[2]]
	case "ram":
		module, ok = cfg.Ram[os.Args[2]]
	case "ip":
		module, ok = cfg.Ip[os.Args[2]]
	case "music":
		module, ok = cfg.Music[os.Args[2]]
	case "pulseaudio":
		module, ok = cfg.PulseAudio[os.Args[2]]
	case "storage":
		module, ok = cfg.Storage[os.Args[2]]
	case "uptime":
		module, ok = cfg.Uptime[os.Args[2]]
	case "custom":
		module, ok = cfg.Custom[os.Args[2]]
	default:
		printUsage()
		os.Exit(1)
	}

	if !ok {
		fmt.Printf("Could not find the %v module named \"%v\" !\n", strings.ToLower(os.Args[1]), os.Args[2])
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
	fmt.Println("\t- battery")
	fmt.Println("\t- date")
	fmt.Println("\t- ip")
	fmt.Println("\t- music")
	fmt.Println("\t- pulseaudio")
	fmt.Println("\t- ram")
	fmt.Println("\t- storage")
	fmt.Println("\t- uptime")
	fmt.Println("\t- custom")
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