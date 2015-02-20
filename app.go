package main

import (
  "fmt"
  "net/http"
  "time"
)

func RouteHandler(w http.ResponseWriter, r *http.Request) {
  start := time.Now()

  http.ServeFile(w, r, "html/index.html")

  elapsed := time.Since(start)
  fmt.Printf("Time take to serve static file %s\n\n", elapsed)
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
