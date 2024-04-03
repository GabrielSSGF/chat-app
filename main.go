package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	messages := make(map[string]string)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		message := r.FormValue("message")
		dt := time.Now().Format("15:04")
		messages[message] = dt
	})

	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		for message, time := range messages {
			fmt.Fprintf(w, "%s\n - At:%s", message, time)
		}

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func login(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("Method not allowed %s", r.Method)
	}

	WriteJSON(w, 200, nil)
	return nil
}

func WriteJSON(writer http.ResponseWriter, status int, value any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(value)
}
