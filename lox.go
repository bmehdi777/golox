package main

import (
	"bufio"
	"fmt"
	"os"
)

type Lox struct {
	hadError bool
}

func (l *Lox) Error(line int, message string) {
	l.report(line, "", message)
}

func (l *Lox) report(line int, where string, message string) {
	fmt.Printf("[line %v] Error %v : %v\n", line, where, message)
	l.hadError = true
}

func (l *Lox) RunFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	content := string(data)
	scanner := NewScanner(content)
	tokens := scanner.ScanTokens()

	fmt.Println("Tokens : ", tokens)
}

func (l *Lox) RunInteractive() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">> ")

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Error : ", err)
			os.Exit(1)
		}

		//l.run(string(line))
		l.hadError = false

		scanner := NewScanner(string(line))
		tokens := scanner.ScanTokens()

		// debug purpose
		if !l.hadError {
			for _, tok := range tokens {
				fmt.Println(tok.toString())
			}
		}

		fmt.Print(">> ")
	}
}

func (l *Lox) run(content string) {
	for pos, char := range content {
		fmt.Println("Char ", pos, " ", string(char))
	}

	if l.hadError {
		os.Exit(1)
	}
}
