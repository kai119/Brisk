package repl

import (
	"reflect"
	"testing"
)

func TestReplFindArgs(t *testing.T) {
	repl := CmdRepl

	tests := []struct {
		args     []string
		expected map[string][]string
	}{
		{
			[]string{
				"-c",
				"Hello World",
				"exit",
			}, map[string][]string{
				"-c": {
					"Hello World",
					"exit",
				},
			},
		},
		{
			[]string{
				"-c",
				"Hello World",
				"exit",
				"-i",
			}, map[string][]string{
				"-c": {
					"Hello World",
					"exit",
				},
				"-i": {"true"},
			},
		},
	}

	for i, tt := range tests {
		repl.Args = map[string][]string{}
		err := repl.ProcessArgs(tt.args)
		if err != nil {
			t.Fatalf("tests[%d] - received the following error when processing args: %s", i, err)
		}

		if !reflect.DeepEqual(tt.expected, repl.Args) {
			t.Fatalf("tests[%d] - result %s does not equal expected %s", i, repl.Args, tt.expected)
		}
	}
}

func TestReplFindArgsFail(t *testing.T) {
	repl := CmdRepl

	tests := []struct {
		args     []string
		expected map[string][]string
	}{
		{
			[]string{"-v"},
			nil,
		},
		{
			[]string{
				"-c",
				"Hello World",
				"exit",
				"-v",
			}, nil,
		},
		{
			[]string{
				"-c",
				"Hello World",
				"exit",
				"-c",
				"Hello World 2",
				"exit",
			}, nil,
		},
		{
			[]string{
				"-c",
				"Hello World",
				"exit",
				"-c",
				"Hello World 2",
				"exit",
				"-i",
			}, nil,
		},
	}

	for i, tt := range tests {
		repl.Args = map[string][]string{}
		err := repl.ProcessArgs(tt.args)
		if err == nil {
			t.Fatalf("test[%d] - did not receive expected error", i)
		}

		if !reflect.DeepEqual(tt.expected, repl.Args) {
			t.Fatalf("test[%d] - result %s does not equal expected %s", i, repl.Args, tt.expected)
		}
	}
}
