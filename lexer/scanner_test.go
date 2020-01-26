package lexer_test

import (
	"strings"
	"testing"

	"../lexer"
)

func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok lexer.TokenType
		lit string
	}{
		{s: `func`, tok: lexer.TOKEN_COMMAND_FUNC, lit: "func"},
		{s: `=`, tok: lexer.TOKEN_VAR_EQUALS, lit: "="},
		{s: `==`, tok: lexer.TOKEN_CONDITION_EQUALS, lit: "=="},
		{s: ``, tok: lexer.TOKEN_EOF},
		{s: `#`, tok: lexer.TOKEN_ERROR, lit: `#`},
		{s: ` `, tok: lexer.TOKEN_WHITESPACE, lit: " "},
	}

	for i, tt := range tests {
		s := lexer.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		} else {
			t.Logf("%d. %q PASS: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}

}
