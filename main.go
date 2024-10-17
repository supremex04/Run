package main

import (
	"fmt"
	"os"
	"os/user"

	"run/repl"
)

func main(){
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey %s! This is Run programming language!\n", user.Username)
	fmt.Printf("Type commands here:\n")
	repl.Start(os.Stdin/*, os.Stdout*/)
}