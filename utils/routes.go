package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type App struct {
	Name string `json:"name"`
	Repo string `json:"repo"`
	Pid string `json:"pid"`
}

func Deploy(w http.ResponseWriter, r *http.Request) {
	app := App{}
	json.NewDecoder(r.Body).Decode(&app)
	cmd := exec.Command("git", "clone", app.Repo, fmt.Sprintf("./apps/%s", app.Name))
	cmd.Run()
	Convert(app.Name)
	pid := CreateApp(app)
	w.Write([]byte(pid))
}


func Restart(w http.ResponseWriter, r *http.Request) {
	app := App{}
	json.NewDecoder(r.Body).Decode(&app)
	existingApp := Read(app.Name)
	if existingApp.Name != "" {
		Kill(existingApp.Pid)
	}
	RestartApp(app)
	pid := CreateApp(app)
	w.Write([]byte(pid))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	app := App{}
	json.NewDecoder(r.Body).Decode(&app)
	DeleteApp(app.Name)
}

func List(w http.ResponseWriter, r *http.Request) {
	apps := ReadAll()
	list, err := json.Marshal(apps)
	check(err)
	w.Write(list)
}
