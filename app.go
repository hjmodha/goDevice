package main

import (
  "fmt"
  "net/http"
)

func RouteHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,"Root Requst")
}

func home(w http.ResponseWriter,r *http.Request){
  http.ServeFile(w, r, "html/index.html")
}

func notfound(w http.ResponseWriter,r *http.Request){
  fmt.Fprintf(w,"Page Not found")
}

func main() {
  http.HandleFunc("/", RouteHandler)
  http.ListenAndServe(":3000", nil)
}
