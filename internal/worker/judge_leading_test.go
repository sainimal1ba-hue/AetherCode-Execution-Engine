
package worker

import (
	"testing"

	"github.com/stjosephsplacements/AetherCode-Execution-Engine/internal/model"
)

func TestExactLeadingBlankLinesSignificant(t *testing.T) {
	cases := []struct {
		name, expected, actual string
		want                   model.Verdict
	}{
		{"leading LF", "42", "\n42", model.VerdictWrongAnswer},
		{"leading CRLF", "42", "\r\n42", model.VerdictWrongAnswer},
		{"two leading LF", "42", "\n\n42", model.VerdictWrongAnswer},
		{"matching leading LF", "\n42", "\n42", model.VerdictAccepted},
		{"trailing newlines unchanged", "42", "42\n\n", model.VerdictAccepted},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := JudgeOutput("exact", tc.expected, tc.actual); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestNormalizeOutputPreservesLeadingBlankLines(t *testing.T) {
	cases := map[string]string{
		"\n42":   "\n42",
		"\n\n42": "\n\n42",
		"\r\n42": "\n42",
		"42\n\n": "42",
		"a\n\nb": "a\n\nb",
	}
	for input, want := range cases {
		if got := normalizeOutput(input); got != want {
			t.Errorf("normalizeOutput(%q) = %q, want %q", input, got, want)
		}
	}
}
