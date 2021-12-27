package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

// Loads the environment configuration to a usable struct.
func LoadConfig() (Config, env.EnvSet) {
	var conf Config

	es, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf, es
}
