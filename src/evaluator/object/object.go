package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.ibm.com/Kai-Mumford-CIC-UK/brisk/src/parser/ast"
)

// Type represents the type of an object
type Type string

// Object Type representations
const (
	INTEGER_OBJ      = "INTEGER"
	STRING_OBJ       = "STRING"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	DICTIONARY_OBJ   = "DICTIONARY"
)

// Object is an evaluated object that contains the type of the object
// its evaluated value
type Object interface {
	Type() Type
	Inspect() string
}

// DictionaryKey represents the key of a dictionary
type DictionaryKey struct {
	Type  Type
	Value uint64
}

// Hashable represents a hashable object
type Hashable interface {
	DictionaryKey() DictionaryKey
}

// Integer represents an integer in BRISK
type Integer struct {
	Value int64
}

// Inspect returns the string representation of the object
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type returns the type of the object
func (i *Integer) Type() Type { return INTEGER_OBJ }

// DictionaryKey the dictionary key of the object
func (i *Integer) DictionaryKey() DictionaryKey {
	return DictionaryKey{Type: i.Type(), Value: uint64(i.Value)}
}

// String represents a string in BRISK
type String struct {
	Value string
}

// Inspect returns the string representation of the object
func (s *String) Inspect() string { return s.Value }

// Type returns the type of the object
func (s *String) Type() Type { return STRING_OBJ }

// DictionaryKey the dictionary key of the object
func (s *String) DictionaryKey() DictionaryKey {
	h := fnv.New64a()
	_, err := h.Write([]byte(s.Value))
	if err != nil {
		return DictionaryKey{}
	}
	return DictionaryKey{Type: s.Type(), Value: h.Sum64()}
}

// Boolean represents a boolean in BRISK
type Boolean struct {
	Value bool
}

// Inspect returns the string representation of the object
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type returns the type of the object
func (b *Boolean) Type() Type { return BOOLEAN_OBJ }

// DictionaryKey the dictionary key of the object
func (b *Boolean) DictionaryKey() DictionaryKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return DictionaryKey{Type: b.Type(), Value: value}
}

// Null represents a null value in BRISK
type Null struct{}

// Inspect returns the string representation of the object
func (n *Null) Inspect() string { return "null" }

// Type returns the type of the object
func (n *Null) Type() Type { return NULL_OBJ }

// ReturnValue represents a return statement in BRISK
type ReturnValue struct {
	Value Object
}

// Inspect returns the string representation of the object
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Type returns the type of the object
func (rv *ReturnValue) Type() Type { return RETURN_VALUE_OBJ }

// Error represents an error in BRISK
type Error struct {
	Message string
}

// Inspect returns the string representation of the object
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Type returns the type of the object
func (e *Error) Type() Type { return ERROR_OBJ }

// Function represents a function in BRISK
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Inspect returns the string representation of the object
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	_, err := out.WriteString("func")
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
	_, err = out.WriteString(") {\n")
	if err != nil {
		return ""
	}
	_, err = out.WriteString(f.Body.String())
	if err != nil {
		return ""
	}
	_, err = out.WriteString("\n}")
	if err != nil {
		return ""
	}

	return out.String()
}

// Type returns the type of the object
func (f *Function) Type() Type { return FUNCTION_OBJ }

// BuiltinFunction represents an instance of a builtin function
type BuiltinFunction func(args ...Object) Object

// Builtin represents the structure of a builtin function
type Builtin struct {
	Fn BuiltinFunction
}

// Inspect returns the string representation of the object
func (b *Builtin) Inspect() string { return "builtin function" }

// Type returns the type of the object
func (b *Builtin) Type() Type { return BUILTIN_OBJ }

// Array represents an array in BRISK
type Array struct {
	Elements []Object
}

// Inspect returns the string representation of the object
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
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

// Type returns the type of the object
func (a *Array) Type() Type { return ARRAY_OBJ }

// DictionaryPair represents a key value pair in a dictionary
type DictionaryPair struct {
	Key   Object
	Value Object
}

// Dictionary represents a dictionary in BRISK
type Dictionary struct {
	Pairs map[DictionaryKey]DictionaryPair
}

// Type returns the type of the object
func (d *Dictionary) Type() Type { return DICTIONARY_OBJ }

// Inspect returns the string representation of the object
func (d *Dictionary) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range d.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
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
