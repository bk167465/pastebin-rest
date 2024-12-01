package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type paste struct {
	content string `json:"content"`
	expiration_interval int `json:"expire_after"`
	short_link  string `json:"short_link"`
	created_at  int `json:"created_at"`
}

func CreateNewPaste(w http.ResponseWriter, req *http.Request) {
    var data map[string]interface{}
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
       http.Error(w, "Inva;id JSON Body", http.StatusBadRequest)
	   return
	}

	fmt.Println(data)

	res := map[string]string {
		"short_link": "temp",
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}
