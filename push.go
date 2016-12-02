package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))))
	http.HandleFunc("/main.html", serveIndex)
	log.Fatal(http.ListenAndServeTLS(":9000", "server.crt", "server.key", nil))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher) // HL
	if ok {
		// if push is supported, push the resource rather than waiting the client to request it.
		log.Println("Pushing css files")
		if err := pusher.Push("/statics/css/bootstrap.min.css", nil); err != nil { // HL
			log.Printf("failed to push: %v", err)
		}

		log.Println("Pushing js files")
		if err := pusher.Push("/statics/js/angular.min.js", nil); err != nil { // HL
			log.Printf("failed to push: %v", err)
		}
	}

	log.Printf("Serving file: %s \n", r.URL.Path[1:])
	http.ServeFile(w, r, "statics/"+r.URL.Path[1:])
}
