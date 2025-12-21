package main

import "fmt"

type TokenType int

const (
	// single character
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// one or two character token
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOr
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

type Token struct {
	tokenType TokenType
	lexeme    string
	line      int
}

func (t *Token) toString() string {
	return fmt.Sprintf("%v %v %d", t.tokenType, t.lexeme, t.line)
}

func NewToken(tokenType TokenType, lexeme string, line int) *Token {
	return &Token{
		tokenType,
		lexeme,
		line,
	}
}
