package cli_test

import (
	"bytes"
	"flag"
	"testing"

	cli "github.com/taranpreetnatt/todo/cmd/todo/cli"
)

func TestGetCreateArgs(t *testing.T) {

	tests := []struct {
		name      string
		input     []string
		expected  string
		expectErr bool
	}{
		{"Create a task with one word", []string{"./todo", "--create", "test"}, "test", false},
		{"Create a task with two words", []string{"./todo", "--create", "one two"}, "one two", false},
		{"Create with empty string", []string{"./todo", "--create", ""}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet(tt.input[0], flag.ExitOnError)
			var buf bytes.Buffer
			got, err := cli.GetArgs(tt.input, &buf)

			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected an error when argument to create is an empty string, but got none.")
				}
			}

			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
