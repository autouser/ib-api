package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
// addr     = "0.0.0.0:4000"
// authAddr = "http://localhost:3000/auth"
)

func main() {
	addr := os.Getenv("BACKEND_ADDR")
	authAddr := os.Getenv("AUTH_ADDR")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API"))
	})

	http.HandleFunc("/backend", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		password := r.URL.Query().Get("password")
		command := r.URL.Query().Get("command")

		requestURL := fmt.Sprintf("http://%s/auth?username=%s&password=%s", authAddr, username, password)
		authResp, err := http.Get(requestURL)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}

		fmt.Printf("Auth Status: %d\n", authResp.StatusCode)

		if authResp.StatusCode == http.StatusOK {
			fmt.Printf("Authorized user '%s' executed command '%s'\n", username, command)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("SUCCESSFUL"))
		} else {
			fmt.Printf("Unauthorized user '%s' can't execute command '%s'\n", username, command)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("FAILED"))
		}

	})

	log.Printf("Starting server on %s...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
