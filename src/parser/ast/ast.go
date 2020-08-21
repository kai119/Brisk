package ast

import (
	"bytes"
	"strings"

	"github.ibm.com/Kai-Mumford-CIC-UK/brisk/src/lexer/token"
)

// Node represents a node on the AST. It contains the token literal for that node
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents a statement and is attached to a node on the AST
type Statement interface {
	Node
	statementNode()
}

// Expression represents a expression and is attached to a node on the AST
type Expression interface {
	Node
	expressionNode()
}

// Program represents the parsed AST and contains a list of statements that are each
// nodes in the tree
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the token literal of the root node in the AST
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		_, err := out.WriteString(s.String())
		if err != nil {
			return ""
		}
	}

	return out.String()
}

// VarStatement is an example of a variable declaration eg. var num = 5
type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode() {}

//TokenLiteral returns the token literal in the variable statement
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }

func (vs *VarStatement) String() string {
	var out bytes.Buffer

	_, err := out.WriteString(vs.TokenLiteral() + " ")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(vs.Name.String())
	if err != nil {
		return ""
	}
	_, err = out.WriteString(" = ")
	if err != nil {
		return ""
	}

	if vs.Value != nil {
		_, err = out.WriteString(vs.Value.String())
		if err != nil {
			return ""
		}
	}

	_, err = out.WriteString(";")
	if err != nil {
		return ""
	}

	return out.String()
}

// Identifier represents the identifier of a variable. For example, in the
// statement 'var num = 5', 'num' is the identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

//TokenLiteral returns the token literal of the identifier
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

// ReturnStatement represents a return statement. For example, a
// return statement could be 'return 5 + 5'
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

//TokenLiteral returns the token literal of the return statement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	_, err := out.WriteString(rs.TokenLiteral() + " ")
	if err != nil {
		return ""
	}

	if rs.ReturnValue != nil {
		_, err = out.WriteString(rs.ReturnValue.String())
		if err != nil {
			return ""
		}
	}

	_, err = out.WriteString(";")
	if err != nil {
		return ""
	}

	return out.String()
}

// ExpressionStatement represencts an expression. An example of an
// expression would be '5 + 3 == 8 * 1'
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

//TokenLiteral returns the token literal of the expression statement
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// IntegerLiteral represents an integer
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

//TokenLiteral returns the token literal of the integer
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

func (il *IntegerLiteral) String() string { return il.Token.Literal }

// StringLiteral represents an string
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

//TokenLiteral returns the token literal of the integer
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

func (sl *StringLiteral) String() string { return sl.Token.Literal }

// PrefixExpression represents a prefix expression. Prefix expressions can either
// be !<expression> or -<expression>
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

//TokenLiteral returns the token literal of the prefix expression
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	_, err := out.WriteString("(")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(pe.Operator)
	if err != nil {
		return ""
	}
	_, err = out.WriteString(pe.Right.String())
	if err != nil {
		return ""
	}
	_, err = out.WriteString(")")
	if err != nil {
		return ""
	}

	return out.String()
}

// InfixExpression represents an infix expression. Examples of these
// include '5 + 5' or '5 * 5 + 5 / 5'
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

//TokenLiteral returns the token literal of the infix expression
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	_, err := out.WriteString("(")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(ie.Left.String())
	if err != nil {
		return ""
	}
	_, err = out.WriteString(" " + ie.Operator + " ")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(ie.Right.String())
	if err != nil {
		return ""
	}
	_, err = out.WriteString(")")
	if err != nil {
		return ""
	}

	return out.String()
}

// Boolean represents a boolean (true or false)
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

//TokenLiteral returns the token literal of the boolean
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

func (b *Boolean) String() string { return b.Token.Literal }

// BlockStatement represents a block of statements. This is usually
// used for lines in between curly braces of if statements or functions
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

//TokenLiteral returns the token literal of the block statement
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		_, err := out.WriteString(s.String())
		if err != nil {
			return ""
		}
	}

	return out.String()
}

// IfExpression represents if statements. If statements are presented in the form
// "if (<condition>) {<consequence>} else {<alternative>}"
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

//TokenLiteral returns the token literal of the if statement
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	_, err := out.WriteString("if")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(ie.Condition.String())
	if err != nil {
		return ""
	}
	_, err = out.WriteString(" ")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(ie.Consequence.String())
	if err != nil {
		return ""
	}

	if ie.Alternative != nil {
		_, err = out.WriteString("else")
		if err != nil {
			return ""
		}
		_, err = out.WriteString(ie.Alternative.String())
		if err != nil {
			return ""
		}
	}

	return out.String()
}

// FunctionLiteral represents a function literal. these are in the form of
// var add = func(a, b) { return a + b; }
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

//TokenLiteral returns the token literal of the function literal
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	_, err := out.WriteString(fl.TokenLiteral())
	if err != nil {
		return ""
	}
	_, err = out.WriteString("(")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(strings.Join(params, ", "))
	if err != nil {
		return ""
	}
	_, err = out.WriteString(")")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(fl.Body.String())
	if err != nil {
		return ""
	}

	return out.String()
}

// CallExpression represents the call of a function literal. an example of this is
// add(1, 5 + 2)
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

//TokenLiteral returns the token literal of the function literal
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, p := range ce.Arguments {
		args = append(args, p.String())
	}

	_, err := out.WriteString(ce.Function.String())
	if err != nil {
		return ""
	}
	_, err = out.WriteString("(")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(strings.Join(args, ", "))
	if err != nil {
		return ""
	}
	_, err = out.WriteString(")")
	if err != nil {
		return ""
	}

	return out.String()
}

// ArrayLiteral represents an array and its elements
type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

//TokenLiteral returns the token literal of the array literal
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range al.Elements {
		elements = append(elements, e.String())
	}

	_, err := out.WriteString("[")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(strings.Join(elements, ", "))
	if err != nil {
		return ""
	}
	_, err = out.WriteString("]")
	if err != nil {
		return ""
	}

	return out.String()
}

// IndexExpression represents an call to an index of an array
type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}

//TokenLiteral returns the token literal of the array literal
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	_, err := out.WriteString("(")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(ie.Left.String())
	if err != nil {
		return ""
	}
	_, err = out.WriteString("[")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(ie.Index.String())
	if err != nil {
		return ""
	}
	_, err = out.WriteString("])")
	if err != nil {
		return ""
	}

	return out.String()
}

// DictionaryLiteral represents a dictionary
type DictionaryLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (dl *DictionaryLiteral) expressionNode() {}

//TokenLiteral returns the token literal of the dictionary literal
func (dl *DictionaryLiteral) TokenLiteral() string { return dl.Token.Literal }

func (dl *DictionaryLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range dl.Pairs {
		pairs = append(pairs, key.String()+": "+value.String())
	}

	_, err := out.WriteString("{")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(strings.Join(pairs, ", "))
	if err != nil {
		return ""
	}
	_, err = out.WriteString("}")
	if err != nil {
		return ""
	}

	return out.String()
}
