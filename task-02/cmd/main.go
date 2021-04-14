package main

import (
	"log"
	"net/http"
	"os"
)

func greet(w http.ResponseWriter, r *http.Request) {
	s := os.Getenv("MONGODB_URI")

	w.Write([]byte("MongoDB URI: " + s))
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		u, err := os.ReadFile("/config/username")
		if err != nil {
			w.Write([]byte("Error: " + err.Error()))
		} else {
			w.Write(u)
		}
	})
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
