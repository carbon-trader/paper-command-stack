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
	switch env := os.Getenv("GO_PROFILE_ACTIVE"); env {

	case "dev", "prd", "qa", "test":
		if _, err := toml.DecodeFile("/var/config/application_"+os.Getenv("GO_PROFILE_ACTIVE")+".toml", &c); err != nil {
			log.Fatal(err)
		}
		break
	default:
		if _, err := toml.DecodeFile("application.toml", &c); err != nil {
			log.Fatal(err)
		}
		break
	}
}
