package utils

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Stat struct {
	CPU float64 `json:"cpu"`
	MEM float64 `json:"mem"`
}

func CreateApp(app App) string {
	nut, err := exec.Command("bash", fmt.Sprintf("./apps/%s/nut.sh", app.Name)).Output()
	check(err)
	app.Pid = string(nut)[:len(nut)-1]
	check(err)
	Write(app)
	return app.Pid
}

func RestartApp(app App) {
	cmd, _ := exec.Command("cd", fmt.Sprintf("./apps/%s;", app.Name), "git", "pull").Output()
	fmt.Println(cmd)
}

func Kill(pid string) {
	exec.Command("pkill", "-P", pid).Run()
}

func Stats(pid string) ([]byte) {
	stats, err := exec.Command("bash", "scripts/performance.sh", pid).Output()
	check(err)
	data := strings.Split(string(stats)[:len(stats)-1], " ")
	var stat Stat

	for i, v := range data {
		conv, err := strconv.ParseFloat(v, 64)
		check(err)
		if (i % 2 == 0) { stat.CPU += conv } else { stat.MEM += conv }
	}

	out, err := json.Marshal(&stat)

	return out 
}
