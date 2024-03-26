package cmd

import (
	"fmt"
	"os"
)

func Init(name string) {
	nutFile, err := os.Create("Nutfile")
	Check(err)
	_, err = nutFile.Write([]byte(name))
	Check(err)
	defer nutFile.Close()
	gitignore, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	Check(err)
	_, err = gitignore.Write([]byte("nut.sh"))
	Check(err)
	defer gitignore.Close()
	fmt.Println("Please write your commands into the Nutfile")
}
