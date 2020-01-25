package lexer

import (
	"bufio"
	"bytes"
	"io"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) Scan() (tok TokenType, lit string) {
	ch := s.Read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isChar(ch) || isSpecial(ch) {
		s.unread()
		return s.scanIndent()
	}

	switch ch {
	case eof:
		return TOKEN_EOF, ""
	case '(':
		return TOKEN_LEFT_BRACKET, string(ch)
	case ')':
		return TOKEN_RIGHT_BRACKET, string(ch)
	case '{':
		return TOKEN_LEFT_CURLY_BRACKET, string(ch)
	case '}':
		return TOKEN_RIGHT_CURLY_BRACKET, string(ch)
	case '[':
		return TOKEN_LEFT_SQUARE_BRACKET, string(ch)
	case ']':
		return TOKEN_RIGHT_SQUARE_BRACKET, string(ch)
	case '"':
		return TOKEN_STR_QUOTE, string(ch)
	case '\'':
		return TOKEN_CHAR_QUOAT, string(ch)
	case '-':
		return TOKEN_MINUS, string(ch)
	case '+':
		return TOKEN_PLUS, string(ch)
	case '/':
		return TOKEN_DIVIDE, string(ch)
	case '*':
		return TOKEN_MULTIPLY, string(ch)
	case '%':
		return TOKEN_MOD, string(ch)
	}

	return TOKEN_ERROR, string(ch)
}

//TODO scanWhiteSpace

func (s *Scanner) scanWhitespace() (tok TokenType, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.Read())

	for {
		if ch := s.Read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return TOKEN_WHITESPACE, buf.String()
}

//TODO scanIndent

func (s *Scanner) scanIndent() (tok TokenType, lit string) {
	return TOKEN_ERROR, ""
}
