package main

import "fmt"

type TokenType string

const (
	// single character
	LEFT_PAREN TokenType = "LEFT_PAREN"
	RIGHT_PAREN = "RIGHT_PAREN"
	LEFT_BRACE = "LEFT_BRACE"
	RIGHT_BRACE = "RIGHT_BRACE"
	COMMA = "COMMA"
	DOT = "DOT"
	MINUS = "MINUS"
	PLUS = "PLUS"
	SEMICOLON = "SEMICOLON"
	SLASH = "SLASH"
	STAR = "STAR"

	// one or two character token
	BANG = "BANG"
	BANG_EQUAL = "BANG_EQUAL"
	EQUAL = "EQUAL"
	EQUAL_EQUAL = "EQUAL_EQUAL"
	GREATER = "GREATER"
	GREATER_EQUAL = "GREATER_EQUAL"
	LESS = "LESS"
	LESS_EQUAL = "LESS_EQUAL"

	// literals
	IDENTIFIER = "IDENTIFIER"
	STRING = "STRING"
	NUMBER = "NUMBER"

	// keywords
	AND = "AND"
	CLASS = "CLASS"
	ELSE = "ELSE"
	FALSE = "FALSE"
	FUN = "FUN"
	FOR = "FOR"
	IF = "IF"
	NIL = "NIL"
	OR = "OR"
	PRINT = "PRINT"
	RETURN = "RETURN"
	SUPER = "SUPER"
	THIS = "THIS"
	TRUE = "TRUE"
	VAR = "VAR"
	WHILE = "WHILE"

	EOF = "EOF"
)

type Token struct {
	tokenType TokenType
	lexeme    string
	line      int
	value     string // every value will be stored as string (no dynamic type in golang)
}

func (t *Token) toString() string {
	return fmt.Sprintf("%v %v %d", t.tokenType, string(t.lexeme), t.line)
}

func NewToken(tokenType TokenType, lexeme string, line int, value string) *Token {
	return &Token{
		tokenType,
		lexeme,
		line,
		value,
	}
}
