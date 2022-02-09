package main

import (
	"flag"
	"log"
	"overload/internal/apiserver"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config_path", "configs/overload.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting API server on port %s", config.BindAddr)
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}