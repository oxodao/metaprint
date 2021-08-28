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

	DEBUG = false
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	if DEBUG {
		os.Exit(startDebug(cfg))
	}

	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	// @TODO: Make this not ugly, like find something to put modules in map but make things work correctly and iterate on it
	// => This could also be used to autogenerate the printUsage thing
	var module modules.Module = nil
	var ok bool = false

	switch os.Args[1] {
	case "date":
		module, ok = cfg.DateTime[os.Args[2]]
	case "ram":
		module, ok = cfg.Ram[os.Args[2]]
	case "ip":
		module, ok = cfg.Ip[os.Args[2]]
	case "music":
		module, ok = cfg.Music[os.Args[2]]
	default:
		printUsage()
		os.Exit(1)
	}

	if !ok {
		fmt.Printf("Could not find the %v module named \"%v\" !\n", strings.ToLower(os.Args[1]), os.Args[2])
		os.Exit(1)
	}

	printResponse(module)
}

func printUsage() {
	fmt.Println("Usage: metaprint <module> <name> [params]")
	fmt.Println()
	fmt.Println("Available modules: ")
	fmt.Println("\t- date")
	fmt.Println("\t- ip")
	fmt.Println("\t- music")
	fmt.Println("\t- ram")
}

func printResponse(module modules.Module) {
	text := module.Print(os.Args[3:])

	if len(module.GetPrefix()) > 0 {
		text = module.GetPrefix() + " " + text
	}

	if len(module.GetSuffix()) > 0 {
		text += " " + module.GetSuffix()
	}

	fmt.Println(text)
}

func startDebug(cfg *config.Config) int {
	fmt.Println("Date")
	for k, v := range cfg.DateTime {
		fmt.Println("\t" + k)
		fmt.Println("\t\t" + v.Print([]string{}))
	}

	fmt.Println()

	fmt.Println("Ram")
	for k, v := range cfg.Ram {
		fmt.Println("\t" + k)
		fmt.Println("\t\t" + v.Print([]string{}))
	}

	fmt.Println()

	fmt.Println("IP")
	for k, v := range cfg.Ip {
		fmt.Println("\t" + k)
		fmt.Println("\t\t" + v.Print([]string{}))
	}

	fmt.Println()

	fmt.Println("Music")
	for k, v := range cfg.Music {
		fmt.Println("\t" + k)
		fmt.Println("\t\t" + v.Print([]string{}))
	}

	return 0
}
