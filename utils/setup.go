package utils

import (
	"net/http"
)

var PORT string = ":3000"; 

func Startup() {
	apps := ReadAll()
	for _, app := range apps {
		RestartApp(app)
		CreateApp(app)
	}
}

func Setup() {
	Startup()
	http.HandleFunc("GET /apps", List)
	http.HandleFunc("GET /apps/stats", Performance)
	http.HandleFunc("POST /apps", Deploy)
	http.HandleFunc("PUT /apps", Restart)
	http.HandleFunc("DELETE /apps", Delete)
	http.ListenAndServe(PORT, nil)
}
