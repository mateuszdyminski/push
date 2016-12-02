package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))))

	http.HandleFunc("/main.html", serveIndexClassic)

	log.Fatal(http.ListenAndServeTLS(":9001", "server.crt", "server.key", nil))
}

func serveIndexClassic(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving file: %s \n", r.URL.Path[1:])
	http.ServeFile(w, r, "statics/"+r.URL.Path[1:])
}
