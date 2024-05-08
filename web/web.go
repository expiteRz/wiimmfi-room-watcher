//go:build !debug

package web

// All codes based on gosumemory (https://github.com/l3lackShark/gosumemory)

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
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
		log.SetPrefix("[Web] ")
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
	overlayFs := http.FileServer(http.Dir(filepath.Join(parentPath, "static")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path[1:]) > 0 {
			overlayFs.ServeHTTP(w, r)
			return
		}
		htmlTitle := "wiimmfi-room-watcher 1.0.0"
		var body string
		switch r.URL.Query().Get("tab") {
		case "1":
			htmlTitle = "Settings | " + htmlTitle
			htmlMain := makeSettingPage()
			body = strings.Replace(GetWebAsset("assets/templates/index.html"), "{{BODY}}", htmlMain, -1)
			break
		case "2":
			htmlTitle = "How to add overlay on OBS | " + htmlTitle
			body = strings.Replace(GetWebAsset("assets/templates/index.html"), "{{BODY}}", "tut written here", -1)
		default:
			body = strings.Replace(GetWebAsset("assets/templates/index.html"), "{{BODY}}", "hello", -1)
		}
		body = strings.Replace(body, "{{TITLE}}", htmlTitle, -1)
		if _, err := fmt.Fprint(w, body); err != nil {
			log.Println(err)
		}
	})
	assetFs, err := fs.Sub(webUiAssets, "assets/deps")
	if err != nil {
		log.Println(err)
	} else {
		http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assetFs))))
	}
	http.HandleFunc("/json", handle)

	// There is no way to print the notification that the user can access after serving, so print it here instead
	fmt.Printf("You can now access to http://%s or add it as a browser source in OBS!\n", utils.LoadedConfig.ServerIp)

	err = http.ListenAndServe(utils.LoadedConfig.ServerIp, nil)
	if err != nil {
		log.SetPrefix("[Web] ")
		log.Println(err)
		time.Sleep(5 * time.Second)
		log.Fatalln(err)
	}
}

func handle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(JSONByte))
}

func makeSettingPage() (html string) {
	elements := reflect.ValueOf(&utils.LoadedConfig).Elem()
	size := elements.NumField()

	for i := range size {
		element := elements.Field(i)
		html += strings.Replace(settingItemTemplate, "{NAME}", settingItemDict[i]["name"], -1)
		html = strings.Replace(html, "{DESC}", settingItemDict[i]["description"], -1)

		child := strings.Replace(settingItemInputTemplate, "{NAME}", settingItemDict[i]["name"], -1)
		switch element.Type().String() {
		case "string":
			child = strings.Replace(child, "{TYPE}", "text", -1)
			if element.String() != "" {
				child = strings.Replace(child, "{ADDON}", `min="0"`, -1)
			} else {
				child = strings.Replace(child, "{ADDON}", "", -1)
			}
			child = strings.Replace(child, "{VALUE}", element.String(), -1)
			break
		case "int":
			child = strings.Replace(child, "{TYPE}", "number", -1)
			if element.Int() != 0 {
				if elements.Type().Field(i).Name == "Interval" {
					child = strings.Replace(child, "{ADDON}", `min="5"`, -1)
				} else {
					child = strings.Replace(child, "{ADDON}", `min="0"`, -1)
				}
			} else {
				child = strings.Replace(child, "{ADDON}", "", -1)
			}
			child = strings.Replace(child, "{VALUE}", strconv.FormatInt(element.Int(), 10), -1)
			break
		case "bool":
			child = strings.Replace(child, "{TYPE}", "checkbox", -1)
			if element.Bool() {
				child = strings.Replace(child, "{ADDON}", `checkbox="true"`, -1)
			} else {
				child = strings.Replace(child, "{ADDON}", "", -1)
			}
			child = strings.Replace(child, "{VALUE}", strconv.FormatBool(element.Bool()), -1)
			break
		}

		html = strings.Replace(html, "{INPUT}", child, -1)
	}

	return html
}
