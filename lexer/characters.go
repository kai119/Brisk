package lexer

var eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isChar(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' || ch <= 'Z')
}

func isSpecial(ch rune) bool {
	return ch == '=' || ch == '<' || ch == '>' || ch == '!'
}