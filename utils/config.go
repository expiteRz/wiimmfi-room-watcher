package utils

import (
	"errors"
	"flag"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Pid        string `json:"pid"`             // Set PID of targeted player
	Interval   int    `json:"interval,string"` // Set seconds for interval getting api response (Default: 10, Min: 5)
	ServerIp   string `json:"server_ip"`
	ServerPort int    `json:"server_port,string"`
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
		if _, err = UpdateConfig(); err != nil {
			log.Fatalln(err)
			return
		}
		log.SetPrefix("[Utils] ")
		log.Println("it seems you're a newbie for wiimmfi-room-watcher. access to http://localhost:24050/?tab=1 and edit your config first")
		//time.Sleep(5 * time.Second)
		callback <- true
		return
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
	if strings.Contains(LoadedConfig.ServerIp, ":") {
		log.Println("unnecessary port number included in server ip. wiimmfi-room-watcher will cut the port number")
		LoadedConfig.ServerIp, _, _ = strings.Cut(LoadedConfig.ServerIp, ":")
	}
	if LoadedConfig.Interval < 5 {
		LoadedConfig.Interval = 5
		log.Println("The program reset your configured interval to 5 to get rid of the possibility of DoS.")
	}
	if LoadedConfig.ServerPort <= 0 {
		LoadedConfig.ServerPort = 24050
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

func UpdateConfig(config ...Config) (needRestart bool, err error) { // bool = true if required values are changed for server hosting
	if len(config) <= 0 {
		config = append(config, Config{
			Pid:        "600000000",
			Interval:   10,
			ServerIp:   "127.0.0.1",
			ServerPort: 24050,
		})
	}
	needRestart = LoadedConfig.ServerPort != config[0].ServerPort
	LoadedConfig = config[0]
	bytes, err := yaml.Marshal(config[0])
	if err != nil {
		return
	}
	file, err := os.Create(configPath)
	if err != nil {
		return
	}
	_, err = file.Write(bytes)
	if err != nil {
		return
	}
	return
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
	}
	if data.Pid == "" {
		return errors.New("pid is not defined")
	}
	if data.ServerPort <= 0 {
		return errors.New("server port is not defined or is invalid")
	}
	return nil
}
