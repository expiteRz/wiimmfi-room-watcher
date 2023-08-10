package web

// All codes based on gosumemory (https://github.com/l3lackShark/gosumemory)

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var JSONByte []byte
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	for {
		ws.WriteMessage(1, JSONByte)
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func SetupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}

func StartServer() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	parentPath := filepath.Dir(ex)
	fs := http.FileServer(http.Dir(filepath.Join(parentPath, "static")))
	http.Handle("/", fs)

	http.HandleFunc("/json", handle)

	// There is no way to print the notification that the user can access after serving, so print it here instead
	fmt.Printf("You can now access to http://%s or add it as a browser source in OBS!\n", utils.LoadedConfig.ServerIp)

	err = http.ListenAndServe(utils.LoadedConfig.ServerIp, nil)
	if err != nil {
		log.Println(err)
		time.Sleep(5 * time.Second)
		log.Fatalln(err)
	}
}

func handle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(JSONByte))
}
