package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 ||
		(len(arguments) == 1 && arguments[0] == "help") {
		fmt.Println("Help text //TODO")
	}

	executeComand(arguments[0])

	// TODO impement first list
	// should read config json to get what packet manager is configured and create a json with the packages on the system
	println(strings.Join(arguments, " ") + "")
}

func executeComand(command string) {
}
