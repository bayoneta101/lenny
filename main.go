package main

import (
	"os"
	"strings"
)

func main() {

	arguments := os.Args[1:]

	// TODO impement first list
	// should read config json to get what packet manager is configured and create a json with the packages on the system
	println(strings.Join(arguments, " ") + "")
}
