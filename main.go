package main 

import (
	"fmt"
	"os"
	"os/user"
	"go_interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type is some commmands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
