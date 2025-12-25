package main

type Scanner struct {
	source string
	tokens []Token

	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		tokens: make([]Token, 0),

		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) scanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current

		s.scanToken()
	}

	s.tokens = append(s.tokens, Token{EOF, "", s.line})
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
		break
	case ')':
		s.addToken(RIGHT_PAREN)
		break
	case '{':
		s.addToken(LEFT_BRACE)
		break
	case '}':
		s.addToken(RIGHT_BRACE)
		break
	case ',':
		s.addToken(COMMA)
		break
	case '.':
		s.addToken(DOT)
		break
	case '-':
		s.addToken(MINUS)
		break
	case '+':
		s.addToken(PLUS)
		break
	case ';':
		s.addToken(SEMICOLON)
		break
	case '*':
		s.addToken(STAR)
		break

	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
		break
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
		break
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
		break
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
		break
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
		break

	case ' ':
	case '\r':
	case '\t':
		// ignore whitespaces
		break
	case '\n':
		s.line += 1
		break

	default:
		globalLox.Error(s.line, "Unexpected character")
		break
	}
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != byte(expected) {
		return false
	}

	s.current += 1
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}

	return rune(s.source[s.current])
}

func (s *Scanner) addToken(tokenType TokenType) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{tokenType, text, s.line})
}

func (s *Scanner) advance() rune {
	s.current += 1
	return rune(s.source[s.current])
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
