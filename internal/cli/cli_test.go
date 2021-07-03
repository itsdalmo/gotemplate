package cli_test

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"reflect"
	"strings"
	"testing"

	"github.com/itsdalmo/gotemplate/internal/cli"
)

func TestCLI(t *testing.T) {
	tests := []struct {
		description string
		command     []string
		expected    string
	}{
		{
			description: "works",
			command:     []string{"--times", "2", "hello", "there"},
			expected: strings.TrimSpace(`
hello
there
hello
there
			 `),
		},
		{
			description: "help message",
			command:     []string{"--help"},
			expected: strings.TrimSpace(`
USAGE
  gotemplate [-n times] <arg>

FLAGS
  -times 1  Number of times to print the messages
			 `),
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			var b bytes.Buffer
			app := cli.New(cli.Options{Writer: &b})

			if err := app.ParseAndRun(context.Background(), tc.command); err != nil && !errors.Is(err, flag.ErrHelp) {
				t.Fatalf("unexpected error: %s", err)
			}
			eq(t, tc.expected, strings.TrimSpace(b.String()))
		})
	}
}

func eq(t *testing.T, expected, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\nexpected:\n%v\n\ngot:\n%v", expected, got)
	}
}
