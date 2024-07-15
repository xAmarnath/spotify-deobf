package handler

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

func Deobfuscate(w http.ResponseWriter, r *http.Request) {
	key_id := r.URL.Query().Get("key_id")
	file_id := r.URL.Query().Get("file_id")

	if key_id == "" || file_id == "" {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("bash", "-c", "./playplay", key_id, file_id)
	
	out, err := cmd.CombinedOutput()

		if err != nil {
		ls := exec.Command("ls")
		ls_out, _ := ls.CombinedOutput()
		http.Error(w, "Failed to deobfuscate: "+err.Error()+" "+string(ls_out), http.StatusInternalServerError)
		return
	}

	resp := map[string]string{
		"resp_key": string(out),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
