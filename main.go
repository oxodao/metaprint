package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/oxodao/metaprint/config"
	"github.com/oxodao/metaprint/modules"
)

const (
	AUTHOR        = "Oxodao"
	VERSION       = "0.1"
	SOFTWARE_NAME = "metaprint"
)

func main() {
	cfg := config.Load()

	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
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
	fmt.Println()
	fmt.Println("Available modules: ")
	fmt.Println("\t- battery")
	fmt.Println("\t- date")
	fmt.Println("\t- ip")
	fmt.Println("\t- music")
	fmt.Println("\t- ram")
	fmt.Println("\t- storage")
	fmt.Println("\t- uptime")
	fmt.Println("\t- custom")
}
