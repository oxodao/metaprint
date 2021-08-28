package main

import (
	"fmt"
	"os"

	"github.com/oxodao/metaprint/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	os.Exit(startReal(cfg))
}

func startReal(cfg *config.Config) int {
	if len(os.Args) < 3 {
		printUsage()
		return 1
	}

	// @TODO: Make golang do something ok there
	// @TODO: Handle error message when module name is not found
	switch os.Args[1] {
	case "date":
		module := cfg.DateTime[os.Args[2]]
		printResponse(module.Prefix, module.Print(os.Args[3:]), module.Suffix)
	case "ram":
		module := cfg.Ram[os.Args[2]]
		printResponse(module.Prefix, module.Print(os.Args[3:]), module.Suffix)
	case "ip":
		module := cfg.Ip[os.Args[2]]
		printResponse(module.Prefix, module.Print(os.Args[3:]), module.Suffix)
	}

	return 0
}

func printUsage() {
	fmt.Println("Usage: metaprint <module> <name> [params]")
	fmt.Println()
	fmt.Println("Available modules: ")
	fmt.Println("\t- date")
	fmt.Println("\t- ram")
	fmt.Println("\t- ip")
}

func printResponse(prefix, text, suffix string) {
	if len(prefix) > 0 {
		prefix += " "
	}

	if len(suffix) > 0 {
		suffix = " " + suffix
	}

	fmt.Println(prefix + text + suffix)
}

func startDebug(cfg *config.Config) {
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
}
