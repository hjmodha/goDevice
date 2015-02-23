package goDevice

import (
  "net/http"
  "strings"
)

func IsMobile(r *http.Request) string {
    userAgent := r.Header.Get("User-Agent")
    deviceType := "Web";

    if strings.Contains(userAgent,"Android") ||
        strings.Contains(userAgent,"webOS") ||
        strings.Contains(userAgent,"iPhone") ||
        strings.Contains(userAgent,"BlackBerry") ||
        strings.Contains(userAgent,"Windows Phone") {
          deviceType = "Mobile"
    }else if strings.Contains(userAgent,"iPad") ||
        strings.Contains(userAgent,"iPod") ||
        (strings.Contains(userAgent,"tablet") ||
        strings.Contains(userAgent,"RX-34") ||
        strings.Contains(userAgent,"FOLIO")) ||
        (strings.Contains(userAgent,"Kindle") ||
        strings.Contains(userAgent,"Mac OS") &&
        strings.Contains(userAgent,"Silk")) ||
        (strings.Contains(userAgent,"AppleWebKit") &&
        strings.Contains(userAgent,"Silk")){
        deviceType = "Tab"
    }else if strings.Contains(userAgent,"TV") ||
      strings.Contains(userAgent,"GoogleTV") ||
      strings.Contains(userAgent,"SmartTV") ||
      strings.Contains(userAgent,"Internet TV") ||
      strings.Contains(userAgent,"NetCast") ||
      strings.Contains(userAgent,"NETTV") ||
      strings.Contains(userAgent,"AppleTV") ||
      strings.Contains(userAgent,"boxee") ||
      strings.Contains(userAgent,"Kylo") ||
      strings.Contains(userAgent,"Roku") ||
      strings.Contains(userAgent,"DLNADOC"){
        deviceType="TV"
      }
  return deviceType
}
