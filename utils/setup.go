package utils

import (
	"fmt"
	"net/http"
)

var PORT string = ":3000"; 

func Format(rgb string, message string) string {
  return fmt.Sprintf("\x1b[38;2;%sm%s\x1b[0m", rgb, message) 
}

func Startup() {
	apps := ReadAll()
	if len(apps) <= 0 {
		return
	}
	fmt.Println(Format("0;255;255", "~nutting your apps up~"))
	for _, app := range apps {
		RestartApp(app)
		CreateApp(app)
	}
	fmt.Println(Format("255;0;0", "~DONE!~"))
}

func Setup() {
	Startup()
	http.HandleFunc("GET /apps", List)
	http.HandleFunc("GET /apps/stats", Performance)
	http.HandleFunc("POST /apps", Deploy)
	http.HandleFunc("PUT /apps", Restart)
	http.HandleFunc("DELETE /apps", Delete)
	fmt.Println(Format("255;0;255", fmt.Sprintf("~http://127.0.0.1:%s~", PORT)))
	http.ListenAndServe(PORT, nil)
}
