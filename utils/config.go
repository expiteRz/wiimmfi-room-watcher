package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Pid      int `json:"pid"`      // Set PID of targeted player
	Interval int `json:"interval"` // Set seconds for interval getting api response (Default: 10, Min: 5)

	ServerIp string `json:"server_ip"`
}

const configFilename = "config.json"

var LoadedConfig Config

func ReadConfig() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Join(filepath.Dir(ex), configFilename)
	file, err := os.Open(exPath)
	if err != nil {
		CreateConfig(exPath)
		log.Fatalf("%s not found. It's automatically generated, please edit it, and restart the program.", configFilename)
	}
	bytes, err := io.ReadAll(file)
	if err = json.Unmarshal(bytes, &LoadedConfig); err != nil {
		panic(err)
	}

	if LoadedConfig.Interval < 5 {
		LoadedConfig.Interval = 5
		log.Println("The program reset your configured interval to 5 to avoid the possibility of dos.")
	}
}

func CreateConfig(path string) {
	config := Config{
		Pid:      600000000,
		Interval: 10,
		ServerIp: "127.0.0.1:24050", // Inspired of gosumemory
	}
	bytes, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(bytes)
	if err != nil {
		panic(err)
	}
}
