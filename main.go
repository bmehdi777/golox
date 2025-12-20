package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		path := os.Args[1]
		runFile(path)
	} else {
		runInteractive()
	}
}

func runFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	content := string(data)
	fmt.Println("File content : ", content)
}

func runInteractive() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">> ")

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Error : ", err)
			os.Exit(1)
		}

		run(string(line))

		fmt.Print(">> ")
	}
}

func run(content string) {
	for pos, char := range content {
		fmt.Println("Char ", pos," ", string(char))
	}
}
