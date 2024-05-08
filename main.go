package main

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/updater"
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"app.rz-public.xyz/wiimmfi-room-watcher/web"
)

func main() {
	configDone := make(chan bool)

	updater.DoUpdate()
	go utils.ReadConfig(configDone)
	<-configDone

	go web.StartParseRoom()
	go web.SetupRoutes()
	web.StartServer()
}
