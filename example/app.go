package main

import (
  "fmt"
  "net/http"
  "time"
  "./deviceChecker"
)

func RouteHandler(w http.ResponseWriter, r *http.Request) {
  start := time.Now()

  deviceType := deviceChecker.IsMobile(r)

  if deviceType=="Mobile" {
    fmt.Fprintf(w,"<h1>Mobile</h1>")
  }else if deviceType=="Web"{
    fmt.Fprintf(w,"<h1>Web</h1>")
  }else if deviceType=="Tab"{
    fmt.Fprintf(w,"<h1>Tablet</h1>")
  }

  elapsed := time.Since(start)
  fmt.Printf("Time take to serve static file %s\n\n", elapsed)
}

func main() {
  http.HandleFunc("/", RouteHandler)
  error := http.ListenAndServe(":8000", nil)
  fmt.Println(error)
}
