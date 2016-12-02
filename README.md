#### Http2 Server Push

HTTP/2 Push allows a web server to send resources to a web browser before the browser gets to request them. It is, for the most part, a performance technique that can help some websites load faster.

#### Http2 Server Push in Golang

In following repository you can see the example of HTTP/2 Server Push written in Golang
 
```go
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

```

##### Requirements

Golang in version 1.8 or higher

##### Run 

```
go run push.go
```

##### Verification

Open browser and go to: [https://localhost:9000/main.html](https://localhost:9000/main.html)

Inspect html and open 'Network' tab. You should see following output: 

![Image of Http/2 server push]
(https://github.com/mateuszdyminski/push/raw/master/presentation/data/with-push.png)

