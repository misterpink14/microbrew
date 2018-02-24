package main

import (
	//"os"
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

const ConfigPath = "./config.json"

var (
	configSingleton Config
	configOnce      sync.Once
)

type Config struct {
	Templates map[string]string `json:"templates"`
}

func GetConfig() Config {
	configOnce.Do(func() {
		byteValue, err := ioutil.ReadFile(ConfigPath)
		if err != nil {
			log.Fatalln("Cannot open config file:", ConfigPath)
		}
		json.Unmarshal(byteValue, &configSingleton)
	})
	return configSingleton
}
