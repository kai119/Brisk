package lexer

import "github.ibm.com/Kai-Mumford-CIC-UK/brisk/src/lexer/token"

// Lexer is a struct that contains information on the string that was inputted,
// the position of the pointer in the string, the read position of the next character in the string
// and the character that the pointer is pointing to
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New creates a new Lexer with the input that was parsed, the position at 0, the read position
// at 1 and the character equal to the first character in the string
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()

	return l
}

//TODO support Unicode with runes at some point

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken will check the character at the current pointer position to see if it matches any
// single-character defined tokens, if it does match it will return a new token with the matching
// type and string literal. If the single character doesn't match a token it will continue searching
// through the input string until whitespace is found. Then it check to see if the string found either
// matches a BRISK command or if it is a user-defined identifier
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.CONDITION_EQUALS, Literal: literal}
		} else {
			tok = newToken(token.VAR_EQUALS, l.ch)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.END_OF_LINE, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '(':
		tok = newToken(token.LEFT_BRACKET, l.ch)
	case ')':
		tok = newToken(token.RIGHT_BRACKET, l.ch)
	case '{':
		tok = newToken(token.LEFT_CURLY_BRACKET, l.ch)
	case '}':
		tok = newToken(token.RIGHT_CURLY_BRACKET, l.ch)
	case '[':
		tok = newToken(token.LEFT_SQUARE_BRACKET, l.ch)
	case ']':
		tok = newToken(token.RIGHT_SQUARE_BRACKET, l.ch)
	case '\'':
		tok = newToken(token.CHAR_QUOAT, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '/':
		tok = newToken(token.DIVIDE, l.ch)
	case '*':
		tok = newToken(token.MULTIPLY, l.ch)
	case '%':
		tok = newToken(token.MOD, l.ch)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.CONDITION_LESS_THAN_EQUAL, Literal: literal}
		} else {
			tok = newToken(token.CONDITION_LESS_THAN, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.CONDITION_MORE_THAN_EQUAL, Literal: literal}
		} else {
			tok = newToken(token.CONDITION_MORE_THAN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.CONDITION_NOT_EQUAL, Literal: literal}
		} else {
			tok = newToken(token.CONDITION_NOT, l.ch)
		}
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isInteger(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isInteger(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

//TODO add float support here
func isInteger(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
