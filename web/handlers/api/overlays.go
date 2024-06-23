package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func OpenFolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if runtime.GOOS != "windows" {
		log.Println("your os is unsupported opening overlay folder")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "your os is unsupported opening overlay folder"})
		return
	}
	folderName := r.PathValue("folderName")
	exPath, err := os.Executable()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "failed to open folder"})
		return
	}
	exactPath := filepath.Dir(exPath)
	overlayPath := filepath.Join(exactPath, "static", folderName)
	_, err = exec.Command("cmd", "/c", "start", overlayPath).Output()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "failed to open folder"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "operation succeeded"})
}
