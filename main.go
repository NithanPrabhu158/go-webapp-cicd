package main

import (
	"log"
	"net/http"
)


func main() {
    fs := http.FileServer(http.Dir("landing_page"))
    http.Handle("/", http.StripPrefix("/", fs))

    log.Println("Server started at :8080")
    err := http.ListenAndServe("0.0.0.0:8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}