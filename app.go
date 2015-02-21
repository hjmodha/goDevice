package main

import (
  "fmt"
  "net/http"
  "time"
  "strings"
)

func RouteHandler(w http.ResponseWriter, r *http.Request) {
  start := time.Now()
  userAgent := r.Header.Get("User-Agent")

  isMobile := strings.Contains(userAgent,"Android") ||
      strings.Contains(userAgent,"webOS") ||
      strings.Contains(userAgent,"iPad") ||
      strings.Contains(userAgent,"iPhone") ||
      strings.Contains(userAgent,"iPod") ||
      strings.Contains(userAgent,"BlackBerry") ||
      strings.Contains(userAgent,"Windows Phone")

  if isMobile {
    fmt.Fprintf(w,"<h1>Mobile</h1>")
  }else{
    fmt.Fprintf(w,"<h1>Web</h1>")
  }

  elapsed := time.Since(start)
  fmt.Printf("Time take to serve static file %s\n\n", elapsed)
}

func main() {
  http.HandleFunc("/", RouteHandler)
  http.ListenAndServe(":3000", nil)
}
