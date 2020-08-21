package token

// Type is the expected string value of the type of the token found
type Type string

// Token is a struct that contains information on both the type and the string literal
// of the token found
type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	END_OF_LINE = ";" // nolint: golint
	NEWLINE     = "\n"
	COMMA       = ","
	COLON       = ":"

	LEFT_BRACKET         = "("
	RIGHT_BRACKET        = ")"
	LEFT_CURLY_BRACKET   = "{"
	RIGHT_CURLY_BRACKET  = "}"
	LEFT_SQUARE_BRACKET  = "["
	RIGHT_SQUARE_BRACKET = "]"

	VAR_DECLARATION = "var"
	VAR_EQUALS      = "="

	TYPE_INT    = "int"
	TYPE_STR    = "str"
	TYPE_BOOL   = "bool"
	TYPE_DOUBLE = "double"
	TYPE_FLOAT  = "float"
	TYPE_CHAR   = "char"

	STR_QUOTE  = "\""
	CHAR_QUOAT = "'"

	PLUS     = "+"
	MINUS    = "-"
	DIVIDE   = "/"
	MOD      = "%"
	MULTIPLY = "*"

	FUNCTION        = "func"
	FUNCTION_RETURN = "return"

	COMMAND_IF    = "if"
	COMMAND_ELSE  = "else"
	COMMAND_FOR   = "for"
	COMMAND_WHILE = "while"

	CONDITION_EQUALS          = "=="
	CONDITION_NOT_EQUAL       = "!="
	CONDITION_NOT             = "!"
	CONDITION_MORE_THAN       = ">"
	CONDITION_MORE_THAN_EQUAL = ">="
	CONDITION_LESS_THAN       = "<"
	CONDITION_LESS_THAN_EQUAL = "<="

	BOOL_TRUE  = "true"
	BOOL_FALSE = "false"
)

var keywords = map[string]Type{
	"func":   FUNCTION,
	"var":    VAR_DECLARATION,
	"if":     COMMAND_IF,
	"else":   COMMAND_ELSE,
	"return": FUNCTION_RETURN,
	"true":   BOOL_TRUE,
	"false":  BOOL_FALSE,
}

// LookupIdent checks the keywords map to see if the string parsed is a BRISK command
// or an identifier
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
