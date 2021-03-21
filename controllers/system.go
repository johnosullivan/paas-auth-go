package controllers

import (
	"encoding/json"
	"net/http"
	"time"
)

type SystemStatus struct {
	Date   time.Time
	Status bool
}

func PingLink(w http.ResponseWriter, r *http.Request) {
	sysStatus := SystemStatus{time.Now(), true}
	js, err := json.Marshal(sysStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
