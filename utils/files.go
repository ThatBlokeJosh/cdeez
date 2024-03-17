package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func Write(app App) {
	file, _ := os.Open("./apps/nuts.json")
	defer file.Close()
	var apps map[string]App = make(map[string]App)
	decoder := json.NewDecoder(file)
	decoder.Decode(&apps)
	apps[app.Name] = app
	toWrite, _ := json.Marshal(apps)
	file, _ = os.Create("./apps/nuts.json")
	file.Write(toWrite)
}

func Read(name string) App {
	file, _ := os.Open("./apps/nuts.json")
	defer file.Close()
	var apps map[string]App = make(map[string]App)
	decoder := json.NewDecoder(file)
	decoder.Decode(&apps)
	return apps[name] 
}

func ReadAll() map[string]App {
	file, _ := os.Open("./apps/nuts.json")
	defer file.Close()
	var apps map[string]App = make(map[string]App)
	decoder := json.NewDecoder(file)
	decoder.Decode(&apps)
	return apps 
}


func DeleteApp(name string) {
	file, _ := os.Open("./apps/nuts.json")
	defer file.Close()
	var apps map[string]App = make(map[string]App)
	decoder := json.NewDecoder(file)
	decoder.Decode(&apps)
	Kill(apps[name].Pid)
	delete(apps, name)
	toWrite, _ := json.Marshal(apps)
	file, _ = os.Create("./apps/nuts.json")
	file.Write(toWrite)
	os.RemoveAll(fmt.Sprintf("./apps/%s", name))
}
