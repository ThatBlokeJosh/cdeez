package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Convert(name string) (appName string) {
	file, err := os.Open(fmt.Sprintf("./apps/%s/Nutfile", name))
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var commands []string = []string{fmt.Sprintf("cd ./apps/%s", name)}
	i := 0

	for scanner.Scan() {
		if i == 0 {
			appName = scanner.Text() 
		} else {
			commands = append(commands, scanner.Text() + " > /dev/null 2>&1")

		}
		i += 1
	}

	commands[len(commands)-1] += " &"
	commands = append(commands, "echo $!")

	shell, err  := os.Create(fmt.Sprintf("./apps/%s/nut.sh", name))
	check(err)
	s := strings.Join(commands, "\n")
	shell.WriteString(s)
	return
} 
