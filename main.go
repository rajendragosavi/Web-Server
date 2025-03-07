package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Running Webserver on port 8080")
    http.HandleFunc("/hello", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
