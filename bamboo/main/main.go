package main

import (
	"fmt"
	"monkey/command"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		command.Start(os.Stdin, os.Stdout, true)
	} else if len(os.Args) == 2 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("no such file")
		}
		command.Start(f, os.Stdout, false)
	} else {

	}
}
