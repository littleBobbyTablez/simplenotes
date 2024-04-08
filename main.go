package main

import (
    "log"
    "net/http"
    handler "simplenotes/handler"
)


func main() {
    http.HandleFunc("/", handler.Home)

    http.HandleFunc("POST /", handler.GetNote)

    http.Handle("GET /output.css", http.FileServer(http.Dir("./templates/static/")))

    log.Fatal(http.ListenAndServe(":8088", nil))
}

