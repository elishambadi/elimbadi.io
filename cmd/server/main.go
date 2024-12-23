package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, GOTH!")
    })

    http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome home. Still work in progress!")
    })

    fmt.Println("Server running on http://localhost:8000")
    http.ListenAndServe(":8000", nil)
}