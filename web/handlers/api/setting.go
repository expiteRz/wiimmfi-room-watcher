package api

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"app.rz-public.xyz/wiimmfi-room-watcher/utils/log"
	"encoding/json"
	"net/http"
)

func SaveSettingHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		log.Logger.Info().Msg("we only receive POST method")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "we only receive POST method"})
		return
	}

	var postData utils.Config
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&postData); err != nil {
		log.Logger.Info().Msg(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "except at JSON parse: " + err.Error()})
		return
	}
	if err := utils.ValidateStoredConfig(postData); err != nil {
		log.Logger.Info().Msg(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "except at JSON parse: " + err.Error()})
		return
	}
	needRestart, err := utils.UpdateConfig(postData)
	if err != nil {
		log.Logger.Info().Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "except at storing posted settings: " + err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	if needRestart {
		json.NewEncoder(w).Encode(map[string]any{"status": "success", "need_restart": true, "message": "setting saved. restart the server to apply update"})
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"status": "success", "need_restart": false, "message": "setting saved"})
}
