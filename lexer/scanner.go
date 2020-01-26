package lexer

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"
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
		return s.scanIdent()
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

func (s *Scanner) scanIdent() (tok TokenType, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.Read())

	for {
		if ch := s.Read(); ch == eof {
			break
		} else if !isChar(ch) && !unicode.IsDigit(ch) && !isSpecial(ch) && ch != '_' && ch != '-' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	switch strings.ToUpper(buf.String()) {
	case "IF":
		return TOKEN_COMMAND_IF, buf.String()
	case "FOR":
		return TOKEN_COMMAND_FOR, buf.String()
	case "VAR":
		return TOKEN_VAR_DECLARATION, buf.String()
	case "CONST":
		return TOKEN_CONST, buf.String()
	case "WHILE":
		return TOKEN_COMMAND_WHILE, buf.String()
	case "INT":
		return TOKEN_TYPE_INT, buf.String()
	case "STR":
		return TOKEN_TYPE_STRING, buf.String()
	case "BOOL":
		return TOKEN_TYPE_BOOL, buf.String()
	case "FLOAT":
		return TOKEN_TYPE_FLOAT, buf.String()
	case "DOUBLE":
		return TOKEN_TYPE_DOUBLE, buf.String()
	case "CHAR":
		return TOKEN_TYPE_CHAR, buf.String()
	case "FUNC":
		return TOKEN_COMMAND_FUNC, buf.String()
	case "=":
		return TOKEN_VAR_EQUALS, buf.String()
	case "==":
		return TOKEN_CONDITION_EQUALS, buf.String()
	case "<":
		return TOKEN_CONDITION_LESS_THAN, buf.String()
	case "<=":
		return TOKEN_CONDITION_LESS_THAN_EQUAL, buf.String()
	case ">":
		return TOKEN_CONDITION_MORE_THAN, buf.String()
	case ">=":
		return TOKEN_CONDITION_MORE_THAN_EQUAL, buf.String()
	case "!=":
		return TOKEN_CONDITION_NOT_EQUAL, buf.String()
	}

	return TOKEN_IDENT, buf.String()
}
