package object

// NewEnvironment creates a new environment with a new variable store
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Environment is the environment of the program. It is responsible for storing
// and binding identifers to their object values
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get returns the object that is binded to the string sent in
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set creates a new binding in the environment between the string and the
// object that have been sent in
func (e *Environment) Set(name string, value Object) Object {
	e.store[name] = value
	return value
}

// NewEnclosedEnvironment creates a nested environment for function executions
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
