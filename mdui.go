package mdui

import "net/http"

func HandleCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(Str_MduiMinCss))
}

func HandleJs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	w.Write([]byte(Str_MduiMinJs))
}
