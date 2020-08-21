package lexer

import (
	"testing"

	"github.ibm.com/Kai-Mumford-CIC-UK/brisk/src/lexer/token"
)

func TestNextToken(t *testing.T) {
	input := `var five = 5;
	var ten = 10;
	var add = func(x, y) {
		x + y;
	};
	var result = add(five, ten);
	!-/*5;
	5 < 10 > 5;
	if (5 < 10) {
		return true;
	} else {
		return false;
	};
	
	10 == 10;
	5 <= 10;
	10 >= 5;
	5 != 10;
	"foo"
	"foo bar"
	[1, 2];
	{"foo": "bar"}`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.VAR_DECLARATION, "var"},
		{token.IDENT, "five"},
		{token.VAR_EQUALS, "="},
		{token.INT, "5"},
		{token.END_OF_LINE, ";"},
		{token.VAR_DECLARATION, "var"},
		{token.IDENT, "ten"},
		{token.VAR_EQUALS, "="},
		{token.INT, "10"},
		{token.END_OF_LINE, ";"},
		{token.VAR_DECLARATION, "var"},
		{token.IDENT, "add"},
		{token.VAR_EQUALS, "="},
		{token.FUNCTION, "func"},
		{token.LEFT_BRACKET, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RIGHT_BRACKET, ")"},
		{token.LEFT_CURLY_BRACKET, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.END_OF_LINE, ";"},
		{token.RIGHT_CURLY_BRACKET, "}"},
		{token.END_OF_LINE, ";"},
		{token.VAR_DECLARATION, "var"},
		{token.IDENT, "result"},
		{token.VAR_EQUALS, "="},
		{token.IDENT, "add"},
		{token.LEFT_BRACKET, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RIGHT_BRACKET, ")"},
		{token.END_OF_LINE, ";"},
		{token.CONDITION_NOT, "!"},
		{token.MINUS, "-"},
		{token.DIVIDE, "/"},
		{token.MULTIPLY, "*"},
		{token.INT, "5"},
		{token.END_OF_LINE, ";"},
		{token.INT, "5"},
		{token.CONDITION_LESS_THAN, "<"},
		{token.INT, "10"},
		{token.CONDITION_MORE_THAN, ">"},
		{token.INT, "5"},
		{token.END_OF_LINE, ";"},
		{token.COMMAND_IF, "if"},
		{token.LEFT_BRACKET, "("},
		{token.INT, "5"},
		{token.CONDITION_LESS_THAN, "<"},
		{token.INT, "10"},
		{token.RIGHT_BRACKET, ")"},
		{token.LEFT_CURLY_BRACKET, "{"},
		{token.FUNCTION_RETURN, "return"},
		{token.BOOL_TRUE, "true"},
		{token.END_OF_LINE, ";"},
		{token.RIGHT_CURLY_BRACKET, "}"},
		{token.COMMAND_ELSE, "else"},
		{token.LEFT_CURLY_BRACKET, "{"},
		{token.FUNCTION_RETURN, "return"},
		{token.BOOL_FALSE, "false"},
		{token.END_OF_LINE, ";"},
		{token.RIGHT_CURLY_BRACKET, "}"},
		{token.END_OF_LINE, ";"},
		{token.INT, "10"},
		{token.CONDITION_EQUALS, "=="},
		{token.INT, "10"},
		{token.END_OF_LINE, ";"},
		{token.INT, "5"},
		{token.CONDITION_LESS_THAN_EQUAL, "<="},
		{token.INT, "10"},
		{token.END_OF_LINE, ";"},
		{token.INT, "10"},
		{token.CONDITION_MORE_THAN_EQUAL, ">="},
		{token.INT, "5"},
		{token.END_OF_LINE, ";"},
		{token.INT, "5"},
		{token.CONDITION_NOT_EQUAL, "!="},
		{token.INT, "10"},
		{token.END_OF_LINE, ";"},
		{token.STRING, "foo"},
		{token.STRING, "foo bar"},
		{token.LEFT_SQUARE_BRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RIGHT_SQUARE_BRACKET, "]"},
		{token.END_OF_LINE, ";"},
		{token.LEFT_CURLY_BRACKET, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RIGHT_CURLY_BRACKET, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] (line[%d]) - token type wrong. expected %q got %q", i, 33+i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected %q got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
