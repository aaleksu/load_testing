package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	BasicAuth string `json:"basic_auth"`
}

var config Config

func GetConfig(configFileName string) Config {
	configFile, err := os.Open(configFileName)

	if err != nil {
		fmt.Printf("\nError: %s\n", err)
	}

	defer configFile.Close()

	configData, _ := ioutil.ReadAll(configFile)

	fmt.Printf("\nconfigData: %s\n", configData)

	json.Unmarshal(configData, &config)

	fmt.Printf("\nconfig: %s\n", config)

	return config
}


