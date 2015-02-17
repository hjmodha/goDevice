package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world!")
}

func home(w http.ResponseWriter,r *http.Request){
  fmt.Fprintf(w,"Home Page")
}

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/home", home)
  http.ListenAndServe(":3000", nil)
}
