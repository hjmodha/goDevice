package deviceChecker

import (
  "net/http"
  "strings"
)

func isMobile(r *http.Request) bool {

  userAgent := r.Header.Get("User-Agent")
  isMobile := strings.Contains(userAgent,"Android") ||
      strings.Contains(userAgent,"webOS") ||
      strings.Contains(userAgent,"iPad") ||
      strings.Contains(userAgent,"iPhone") ||
      strings.Contains(userAgent,"iPod") ||
      strings.Contains(userAgent,"BlackBerry") ||
      strings.Contains(userAgent,"Windows Phone")

  return isMobile;
}
