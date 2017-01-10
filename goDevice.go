package goDevice

import (
	"net/http"
	"strings"
)

type Device string

const (
	MOBILE Device = "Mobile"
	TABLET Device = "Tab"
	WEB    Device = "Web"
	TV     Device = "TV"
)

func GetType(r *http.Request) Device {
	userAgent := r.Header.Get("User-Agent")
	deviceType := WEB

	if strings.Contains(userAgent, "Android") ||
		strings.Contains(userAgent, "webOS") ||
		strings.Contains(userAgent, "iPhone") ||
		strings.Contains(userAgent, "BlackBerry") ||
		strings.Contains(userAgent, "Windows Phone") {
		deviceType = MOBILE
	} else if strings.Contains(userAgent, "iPad") ||
		strings.Contains(userAgent, "iPod") ||
		(strings.Contains(userAgent, "tablet") ||
			strings.Contains(userAgent, "RX-34") ||
			strings.Contains(userAgent, "FOLIO")) ||
		(strings.Contains(userAgent, "Kindle") ||
			strings.Contains(userAgent, "Mac OS") &&
				strings.Contains(userAgent, "Silk")) ||
		(strings.Contains(userAgent, "AppleWebKit") &&
			strings.Contains(userAgent, "Silk")) {
		deviceType = TABLET
	} else if strings.Contains(userAgent, "TV") ||
		strings.Contains(userAgent, "NetCast") ||
		strings.Contains(userAgent, "boxee") ||
		strings.Contains(userAgent, "Kylo") ||
		strings.Contains(userAgent, "Roku") ||
		strings.Contains(userAgent, "DLNADOC") {
		deviceType = TV
	}
	return deviceType
}
