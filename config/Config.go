package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//Config database
type Config struct {
	Server   string
	Database string
}

//Read func
func (c *Config) Read() {
	switch env := os.Getenv("ENVIRONMENT"); env {
	case "dev", "prd", "qa":
		if _, err := toml.DecodeFile("/var/config/config_"+os.Getenv("ENVIRONMENT")+".toml", &c); err != nil {
			log.Fatal(err)
		}
		break
	default:
		if _, err := toml.DecodeFile("config.toml", &c); err != nil {
			log.Fatal(err)
		}
		break
	}
}
