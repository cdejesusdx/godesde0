package webserver

import (
	"net/http"
)

func MiWebServer() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}

func MiWebServer2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./webserver/index.html")
	})
	http.ListenAndServe(":3000", nil)
}
