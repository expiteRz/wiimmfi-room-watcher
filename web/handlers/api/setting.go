package api

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func SaveSettingHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "we only receive POST method", http.StatusNotAcceptable)
		return
	}

	var postData utils.Config
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&postData); err != nil {
		http.Error(w, "except at JSON parse: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := utils.ValidateStoredConfig(postData); err != nil {
		http.Error(w, "except at JSON parse: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := utils.WriteConfig(postData); err != nil {
		http.Error(w, "except at storing posted settings: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "setting save done")
}
