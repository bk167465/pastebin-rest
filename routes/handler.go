package routes

import (
	"encoding/json"
	"net/http"
)

type paste struct {
	content string `json:"content"`
	expiration_interval int `json:"expire_after"`
	short_link  string `json:"short_link"`
	created_at  int `json:"created_at"`
}

func CreateNewPaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	var req map[string]string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
    
	res := map[string]string {
	   "short_link": "temp",
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetPaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
        http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	params := r.URL.Query()
	short_link := params.Get("short_link")

	res := map[string]string {
		"content": short_link,
	} 

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)

}
