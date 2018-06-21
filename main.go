///
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jun 21 09:59:27 2018
//
package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/MiyamonY/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
