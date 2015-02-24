package goDevice

import (
	"net/http"
	"strings"
)

var Database = map[string][]string{
	"Mobile": []string{
		"Android",
		"webOS",
		"iPhone",
		"BlackBerry",
		"Windows Phone",
	},
	"Tab": []string{
		"iPad",
		"iPod",
		"tablet",
		"RX-34",
		"FOLIO",
		"Kindle",
		"Mac OS",
		"Silk",
		"AppleWebKit",
	},
	"TV": []string{
		"TV",
		"NetCast",
		"boxee",
		"Kylo",
		"Roku",
		"DLNADOC",
	},
}

func GetType(r *http.Request) string {
	userAgent := r.Header.Get("User-Agent")

	for typ, list := range Database {
		for _, extract := range list {
			if strings.Contains(userAgent, extract) {
				return typ
			}
		}
	}

	return "web"
}
