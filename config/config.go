package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/oxodao/metaprint/modules"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Battery  map[string]modules.Battery
	Custom   map[string]modules.Custom
	DateTime map[string]modules.Date
	Ram      map[string]modules.Ram
	Ip       map[string]modules.IP
	Music    map[string]modules.Music
	Uptime   map[string]modules.Uptime
}

// Load the configuration struct from a json file
func Load() *Config {
	var config Config

	globalPath := getPath() + "config.yml"

	hostname, err := os.Hostname()
	hasHostname := err == nil
	hostPath := getPath() + hostname + ".yml"

	anyFound := false
	if _, err := os.Stat(globalPath); !os.IsNotExist(err) {
		err = loadConfig(&config, globalPath)
		if err == nil {
			anyFound = true
		}
	}

	if !hasHostname {
		return &config
	}

	// Loading it after will override the previous values if those exists
	if _, err := os.Stat(hostPath); !os.IsNotExist(err) {
		err = loadConfig(&config, hostPath)
		if err == nil {
			anyFound = true
		}
	}

	if !anyFound {
		fmt.Println("Could not load any config !")
		os.Exit(1)
	}

	return &config
}

func loadConfig(cfg *Config, path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, cfg)

	return err
}

func getPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "./"
	}

	configPath := dirname + "/.config/metaprint"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return "./"
	}

	return configPath + "/"
}
