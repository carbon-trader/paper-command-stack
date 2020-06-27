package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

//Config database
type Config struct {
	Server   string
	Database string
}

//Read func
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
