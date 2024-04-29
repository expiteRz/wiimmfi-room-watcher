package main

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/updater"
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"app.rz-public.xyz/wiimmfi-room-watcher/web"
)

func main() {
	updater.DoUpdate()
	utils.ReadConfig()
	go web.StartParseRoom()
	go web.SetupRoutes()
	web.StartServer()
}
