package mdui

import (
	"net/http"

	"github.com/StevenZack/mux"

	"github.com/StevenZack/mdui/views"
)

// HandleAll e.g prefix="/res/"
func HandleAll(server *mux.Server, prefix string) {
	server.HandleWoff(prefix+"fonts/roboto/Roboto-Medium.woff", views.Bytes_RobotoMediumWoff)
	server.HandleWoff(prefix+"fonts/roboto/Roboto-Medium.woff2", views.Bytes_RobotoMediumWoff2)
	server.HandleWoff(prefix+"icons/material-icons/MaterialIcons-Regular.woff", views.Bytes_MaterialiconsRegularWoff)
	server.HandleWoff(prefix+"icons/material-icons/MaterialIcons-Regular.woff2", views.Bytes_MaterialiconsRegularWoff2)
	server.HandleCss(prefix+"css/mdui.min.css", views.Bytes_MduiMinCss)
	server.HandleJs(prefix+"js/mdui.min.js", views.Bytes_MduiMinJs)
}

func HandleCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(views.Bytes_MduiMinCss)
}

func HandleJs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	w.Write(views.Bytes_MduiMinJs)
}
