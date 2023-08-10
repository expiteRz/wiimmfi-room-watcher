package utils

import (
	"encoding/json"
	"errors"
	cnf "github.com/l3lackShark/config"
	"github.com/spf13/cast"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Pid      int `json:"pid"`      // Set PID of targeted player
	Interval int `json:"interval"` // Set seconds for interval getting api response (Default: 10, Min: 5)

	ServerIp string `json:"server_ip"`
}

const configFilename = "config.ini"

var LoadedConfig Config

func ReadConfig() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Join(filepath.Dir(ex), configFilename)

	var (
		config cnf.Config
		parsed map[string]string
	)
	config, err = cnf.SetFile(exPath)
	if err != nil && errors.Is(err, cnf.ErrDoesNotExist) {
		d := []byte(`[General] ; interval = seconds, pid = your pid
interval = 5
pid = 600000000

[Web]
serverip = 127.0.0.1:24050
`)
		if err = os.WriteFile(exPath, d, 0644); err != nil {
			panic(err)
		}
		log.Printf("%s not found. It's automatically generated, please edit it, and restart the program.", configFilename)
		time.Sleep(5 * time.Second)
		os.Exit(1)
		return
	}
	parsed, err = config.Parse()
	if err != nil {
		panic(err)
	}

	LoadedConfig = Config{
		Pid:      cast.ToInt(parsed["pid"]),
		Interval: cast.ToInt(parsed["interval"]),
		ServerIp: parsed["serverip"],
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
