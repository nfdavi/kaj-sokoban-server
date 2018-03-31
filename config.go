package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"log"
)

type Config struct {
	Database struct {
		Server   string
		User     string
		Password string
		DBname   string
	}
}

func CreateDsnFromConfig(configFile string) string {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, configFile)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Server, cfg.Database.DBname)
}
