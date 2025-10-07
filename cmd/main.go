package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Resp struct {
	Message string `json:"message"`
	Time    string `json:"time"`
	Version string `json:"version,omitempty"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rsp := Resp{
			Message: "Hello from go-ci-app",
			Time:    time.Now().Format(time.RFC3339),
			Version: os.Getenv("APP_VERSION"),
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(rsp)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
