package main

import (
	"fmt"
	"github.com/artarts36/specw"
	"gopkg.in/yaml.v3"
)

type Config struct {
	URL      specw.URL       `yaml:"url" json:"url"`
	IP       specw.IP        `yaml:"ip" json:"ip"`
	LogLevel specw.SlogLevel `yaml:"log_level" json:"log_level"`

	EnvString specw.Env[string]   `yaml:"env_string"`
	EnvIP     specw.Env[specw.IP] `yaml:"env_ip"`
}

const content = `
url: http://google.com
ip: 192.168.0.1
log_level: info
`

func main() {
	var cfg Config

	_ = yaml.Unmarshal([]byte(content), &cfg)

	fmt.Println(cfg)
}
