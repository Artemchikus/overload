package main

import (
	"flag"
	"log"
	"overload/internal/apiserver"

	"github.com/BurntSushi/toml"
)

var (
	configPath string // путь к кнофигу
)

func init() {
	flag.StringVar(&configPath, "config_path", "configs/overload.toml", "path to config file") // назначение флага и дефолтного значения переменной пути
}

func main() {
	flag.Parse() // парсинг флага

	config := apiserver.NewConfig() // получение конфига сервера по умолчанию
	_, err := toml.DecodeFile(configPath, config) // заполнения полей конфига значениями из томл файла
	if err != nil {
		log.Fatal(err)
	} 

	// запуск сервера
	log.Printf("Starting API server on port %s", config.BindAddr)
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}