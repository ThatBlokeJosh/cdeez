package utils

import (
	"fmt"
	"os/exec"
)

func CreateApp(app App) string {
	nut, err := exec.Command("bash", fmt.Sprintf("./apps/%s/nut.sh", app.Name)).Output()
	check(err)
	app.Pid = string(nut)[:len(nut)-1]
	check(err)
	Write(app)
	return app.Pid
}

func RestartApp(app App) {
	cmd := exec.Command("cd", fmt.Sprintf("./apps/%s;", app.Name), "git", "pull")
	cmd.Run()
}

func Kill(pid string) {
	exec.Command("pkill", "-P", pid).Run()
}
