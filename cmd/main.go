package main

import (
	"awesomeProject2/apiServer"
	"awesomeProject2/configs"
	"awesomeProject2/repository"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := &configs.Config{}
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	db, err := repository.NewPostgresDB(config.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	dataBase := repository.NewDb(db)
	server, err := apiServer.New(config, dataBase)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
