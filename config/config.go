package config

import (
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
func Load() (*Config, error) {
	data, err := ioutil.ReadFile(getPath())
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)

	return &config, err
}

func getPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "./config.yml"
	}

	configPath := dirname + "/.config/metaprint"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return "./config.yml"
	}

	return configPath + "/config.yml"
}
