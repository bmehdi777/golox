package main

import (
	"os"
)
var globalLox

func main() {
	globalLox = Lox{}
	if len(os.Args) > 1 {
		path := os.Args[1]
		globalLox.RunFile(path)
	} else {
		globalLox.RunInteractive()
	}
}

