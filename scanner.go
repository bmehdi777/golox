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
	default:
		globalLox.Error(s.line, "Unexpected character")
		break;
	}
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
