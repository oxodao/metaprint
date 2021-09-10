package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/oxodao/metaprint/modules"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Battery    map[string]modules.Battery    `yaml:"battery"`
	Custom     map[string]modules.Custom     `yaml:"custom"`
	DateTime   map[string]modules.Date       `yaml:"datetime"`
	Ip         map[string]modules.IP         `yaml:"ip"`
	Music      map[string]modules.Music      `yaml:"music"`
	PulseAudio map[string]modules.PulseAudio `yaml:"pulseaudio"`
	Ram        map[string]modules.Ram        `yaml:"ram"`
	Storage    map[string]modules.Storage    `yaml:"storage"`
	Uptime     map[string]modules.Uptime     `yaml:"uptime"`
}

func GetModulesAvailable() []string {
	modulesAvailable := []string{}

	t := reflect.TypeOf(Config{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		modulesAvailable = append(modulesAvailable, field.Tag.Get("yaml"))
	}

	return modulesAvailable
}

func getFieldNameFromModuleName(moduleType string) string {
	t := reflect.TypeOf(Config{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("yaml") == moduleType {
			return field.Name
		}
	}

	return ""
}

func (c Config) FindModule(moduleType, name string) (modules.Module, error) {
	fieldName := getFieldNameFromModuleName(moduleType)
	if len(fieldName) == 0 {
		return nil, nil
	}

	t := reflect.ValueOf(c)
	value := t.FieldByName(fieldName).MapIndex(reflect.ValueOf(name))

	if value.Kind() == reflect.Invalid {
		return nil, errors.New("could not find the " + fieldName + " module named " + name)
	}

	module := value.Interface().(modules.Module)

	return module, nil
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
