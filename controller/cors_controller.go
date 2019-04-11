package controller

import (
	"net/http"
	"time"

    "github.com/igor822/gomicro/utils"
)

// setPreflightHeadersAndStatus add the required preflight headers on OPTIONS requests
func setPreflightHeadersAndStatus(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "authorization,content-type,refresh_token")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE,FETCH")
	w.Header().Set("Vary", "Access-Control-Request-Headers")
	w.WriteHeader(http.StatusNoContent)
}

func addNotIgnoredHeader(w http.ResponseWriter, r *http.Request, ignoreHeaders []string, headerKey string, headerValue string) {
	if !utils.StringInSlice(headerKey, ignoreHeaders) {
		w.Header().Set(headerKey, headerValue)
	}
}

// CorsHandler manages cors activities
func corsHandler(w http.ResponseWriter, r *http.Request) {
	addNotIgnoredHeader(w, r, []string{}, "Access-Control-Allow-Origin", "*")
	addNotIgnoredHeader(w, r, []string{}, "Access-Control-Expose-Headers", "Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,refresh_token,access_token,storeid")
	addNotIgnoredHeader(w, r, []string{}, "Date", time.Now().Format(time.RFC1123))
	addNotIgnoredHeader(w, r, []string{}, "Connection", "keep-alive")
	addNotIgnoredHeader(w, r, []string{}, "X-Powered-By", "Golang")

	setPreflightHeadersAndStatus(w)
}

