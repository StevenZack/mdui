package mdui

import "net/http"

func HandleCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(Bytes_MduiMinCss)
}

func HandleJs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	w.Write(Bytes_MduiMinJs)
}
