package main

import (
	"encoding/json"
	"net/http"
)

type ResponseApi struct {
	StatusCode int    `json:"status_code,omitempty"`
	Status     string `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := ResponseApi{
			StatusCode: http.StatusOK,
			Status:     "ok",
			Message:    "success run this app",
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&response)
	})

	err := http.ListenAndServe(":9010", r)
	if err != nil {
		panic(err.Error())
	}
	
}
