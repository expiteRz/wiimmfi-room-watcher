//go:build debug

package web

// All codes based on gosumemory (https://github.com/l3lackShark/gosumemory)

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"app.rz-public.xyz/wiimmfi-room-watcher/utils/log"
	"app.rz-public.xyz/wiimmfi-room-watcher/web/handlers/api"
	"fmt"
	"github.com/gorilla/websocket"
	"io/fs"
	"net"
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
		log.Logger.Error().Err(err).Send()
	}
	for {
		ws.WriteMessage(1, JSONByte)
		time.Sleep(time.Duration(5) * time.Second)
	}
}

type ApiFunc func(http.ResponseWriter, *http.Request)

var apiHandleList = map[string]ApiFunc{
	"/api/setting/save":               api.SaveSettingHandle,
	"/api/overlays/open/{folderName}": api.OpenFolder,
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsEndpoint)
	ex, err := os.Executable()
	if err != nil {
		log.Logger.Fatal().Err(err).Send()
	}
	parentPath := filepath.Dir(ex)
	overlayFs := http.FileServer(http.Dir(filepath.Join(parentPath, "static")))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path[1:]) > 0 {
			overlayFs.ServeHTTP(w, r)
			return
		}
		// For debug
		bytes, err := os.ReadFile("web/assets/templates/index.html")
		if err != nil {
			return
		}
		s := string(bytes)
		htmlTitle := "wiimmfi-room-watcher 1.0.0"
		var body string
		switch r.URL.Query().Get("tab") {
		case "1":
			htmlTitle = "Settings | " + htmlTitle
			htmlMain := makeSettingPage()
			//body = strings.Replace(GetWebAsset("assets/templates/index.html"), "{{BODY}}", htmlMain, -1)
			body = strings.Replace(s, "{{BODY}}", htmlMain, -1)
			break
		case "2":
			htmlTitle = "How to add overlay on OBS | " + htmlTitle
			//body = strings.Replace(GetWebAsset("assets/templates/index.html"), "{{BODY}}", "tut written here", -1)
			// TODO: View instruction to add overlay on OBS
			body = strings.Replace(s, "{{BODY}}", "tut written here", -1)
		default:
			htmlTitle = "Local overlays | " + htmlTitle
			//body = strings.Replace(GetWebAsset("assets/templates/index.html"), "{{BODY}}", "hello", -1)
			// TODO: View overlay list
			body = strings.Replace(s, "{{BODY}}", makeLibrary(), -1)
		}
		body = strings.Replace(body, "{{TITLE}}", htmlTitle, -1)
		body = strings.Replace(body, "{{DEBUGSCRIPT}}", "", -1)
		if _, err := fmt.Fprint(w, body); err != nil {
			log.Logger.Error().Err(err).Send()
		}
	})
	//assetFs, err := fs.Sub(webUiAssets, "assets/deps")
	assetFs, err := fs.Sub(os.DirFS("./web"), "assets/deps")
	if err != nil {
		log.Logger.Error().Err(err).Send()
	} else {
		mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assetFs))))
	}
	mux.HandleFunc("/json", jsonHandle)
	// Serve API handlers
	for s := range apiHandleList {
		mux.HandleFunc(s, apiHandleList[s])
	}

	return mux
}

func StartServer() {
	addr, err := portCheck(utils.LoadedConfig.ServerIp, utils.LoadedConfig.ServerPort)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Logger.Fatal().Err(err).Send()
		return
	}
	mux := setupRoutes()
	// There is no way to print the notification that the user can access after serving, so print it here instead
	log.Logger.Info().Msgf("Start hosting on http://%s", addr)

	if err = http.Serve(l, mux); err != nil {
		log.Logger.Error().Err(err).Send()
		time.Sleep(5 * time.Second)
		os.Exit(-1)
	}
}

func jsonHandle(w http.ResponseWriter, _ *http.Request) {
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
		child = strings.Replace(child, "{ID}", settingItemDict[i]["id"], -1)
		switch element.Type().String() {
		case "string":
			child = strings.Replace(child, "{TYPE}", "text", -1)
			if element.String() != "" && elements.Type().Field(i).Name == "ServerIp" {
				child = strings.Replace(child, "{ADDON}", `maxlength="15"`, -1)
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
				} else if elements.Type().Field(i).Name == "ServerPort" {
					child = strings.Replace(child, "{ADDON}", `min="1" max="65535"`, -1)
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

	html += "<div>" + settingSubmitSaveTemplate + "</div>"

	return html
}

func makeLibrary() (html string) {
	exPath, err := os.Executable()
	if err != nil {
		return internalErrorTemplate
	}
	exPath = filepath.Join(filepath.Dir(exPath), "static")
	dir, err := os.ReadDir(exPath)
	if err != nil {
		return internalErrorTemplate
	}
	for _, entry := range dir {
		html += strings.Replace(overlayItemTemplate, "{NAME}", entry.Name(), -1)
		html = strings.Replace(html, "{ID}", entry.Name(), -1)
		html = strings.Replace(html, "{URL}",
			fmt.Sprint("http://", utils.LoadedConfig.ServerIp, ":", utils.LoadedConfig.ServerPort, "/", entry.Name(), "/index.html"),
			-1)
		html = strings.Replace(html, "{ORIGINPATH}", fmt.Sprint("http://", utils.LoadedConfig.ServerIp, ":", utils.LoadedConfig.ServerPort, "/", entry.Name()), -1)
	}

	return
}
