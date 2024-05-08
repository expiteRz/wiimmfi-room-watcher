package utils

import (
	"errors"
	"flag"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Pid      string `json:"pid"`             // Set PID of targeted player
	Interval int    `json:"interval,string"` // Set seconds for interval getting api response (Default: 10, Min: 5)
	ServerIp string `json:"server_ip"`
}

var configFilename = flag.String("config", "config.yml", "Set specific config file")
var configPath string
var LoadedConfig Config

//func ReadConfig() {
//	flag.Parse()
//	ex, err := os.Executable()
//	if err != nil {
//		panic(err)
//	}
//	exPath := filepath.Join(filepath.Dir(ex), *configFilename)
//
//	var (
//		config cnf.Config
//		parsed map[string]string
//	)
//	config, err = cnf.SetFile(exPath)
//	if err != nil && errors.Is(err, cnf.ErrDoesNotExist) {
//		d := []byte(`[General] ; interval = seconds, pid = your pid
//interval = 5
//pid = 600000000
//
//[Web]
//serverip = 127.0.0.1:24050
//`)
//		if err = os.WriteFile(exPath, d, 0644); err != nil {
//			panic(err)
//		}
//		log.SetPrefix("[Config] ")
//		log.Printf("%s not found. It's automatically generated, please edit it, and restart the program.", *configFilename)
//		time.Sleep(5 * time.Second)
//		os.Exit(1)
//		return
//	}
//	parsed, err = config.Parse()
//	if err != nil {
//		panic(err)
//	}
//
//	LoadedConfig = Config{
//		Pid:      parsed["pid"],
//		Interval: cast.ToInt(parsed["interval"]),
//		ServerIp: parsed["serverip"],
//	}
//
//	if LoadedConfig.Interval < 5 {
//		LoadedConfig.Interval = 5
//		log.Println("The program reset your configured interval to 5 to avoid the possibility of dos.")
//	}
//}

func ReadConfig(callback chan bool) {
	var err error
	configPath, err = getAbsPath()
	if err != nil {
		log.Fatalln(err)
		return
	}
	file, err := os.Open(configPath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatalln(err)
			return
		}
		if err = WriteConfig(); err != nil {
			log.Fatalln(err)
			return
		}
		// TODO: At the plan, the user can edit config in built-in setting screen and the app will reload config automatically
		log.Println("No config file was found. wiimmfi-room-watcher created a new config file (config.toml), and please edit it as your wish.")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if err = yaml.Unmarshal(bytes, &LoadedConfig); err != nil {
		log.Fatalln(err)
		return
	}
	if LoadedConfig.Interval < 5 {
		LoadedConfig.Interval = 5
		log.Println("The program reset your configured interval to 5 to get rid of the possibility of DoS.")
	}
	callback <- true
}

//func CreateConfig(path string) {
//	config := Config{
//		Pid:      600000000,
//		Interval: 10,
//		ServerIp: "127.0.0.1:24050", // Inspired of gosumemory
//	}
//	bytes, err := json.MarshalIndent(config, "", "    ")
//	if err != nil {
//		panic(err)
//	}
//
//	file, err := os.Create(path)
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//	_, err = file.Write(bytes)
//	if err != nil {
//		panic(err)
//	}
//}

func WriteConfig(config ...Config) error {
	if len(config) <= 0 {
		config = append(config, Config{
			Pid:      "600000000",
			Interval: 5,
			ServerIp: "127.0.0.1:24050",
		})
	}
	bytes, err := yaml.Marshal(config[0])
	if err != nil {
		return err
	}
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func getAbsPath() (string, error) {
	flag.Parse()
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir, _ := filepath.Split(ex)
	var exPath string
	if !filepath.IsAbs(*configFilename) {
		exPath = filepath.Join(dir, *configFilename)
	}
	return exPath, nil
}

func ValidateStoredConfig(data Config) error {
	if data.ServerIp == "" {
		return errors.New("server ip is not defined")
	} else if data.Pid == "" {
		return errors.New("pid is not defined")
	}
	return nil
}
