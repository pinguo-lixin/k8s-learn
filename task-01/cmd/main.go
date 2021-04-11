package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func gethostname(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request")
	hName, _ := os.Hostname()
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h2>Hostname is < %s ></h2>", hName)
}

func main() {
	http.HandleFunc("/", gethostname)
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
