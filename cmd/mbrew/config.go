package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

var (
	configSingleton Config
	configOnce      sync.Once
)

type Config struct {
	Templates map[string]string `json:"templates"`
}

// GetConfig gets config information from a given path and filename
func GetConfig(path, filename string) Config {
	configOnce.Do(func() {
		configJSON := fmt.Sprintf("%s/%s.json", path, filename)
		log.Println(configJSON)
		byteValue, err := ioutil.ReadFile(configJSON)
		if err != nil {
			log.Fatalln("Cannot open config file:", configJSON)
		}
		json.Unmarshal(byteValue, &configSingleton)
	})
	return configSingleton
}
