package main

import (
	"os"
)

var globalLox Lox

func main() {
	globalLox = Lox{}
	if len(os.Args) > 1 {
		path := os.Args[1]
		globalLox.RunFile(path)
	} else {
		globalLox.RunInteractive()
	}
}

