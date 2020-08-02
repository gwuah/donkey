package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/gwuah/donkey/repl"
)

func print(value interface{}) {
	fmt.Println(value)
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nHello %s!, Welcome to DonkeyLang REPL!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)

}
