package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var config Config
var configLoaded = false

const (
	ConfigFile string = "example.yaml"
)

type Pihole struct {
	Url    string `yaml:"url,omitempty"`
	Pwhash string `yaml:"pwhash,omitempty"`
}
type Config struct {
	Pihole Pihole `yaml:"pihole,omitempty"`
}

func Configuration() *Config {
	if !configLoaded {
		config = Config{}

		configdata, err := ioutil.ReadFile(ConfigFile)
		if err == nil {
			yaml.Unmarshal([]byte(configdata), &config)
		}

		if config.Pihole.Url == "" {
			config.Pihole.Url = "http://localhost/admin/api.php"
		}
		configLoaded = true
	}
	return &config
}
