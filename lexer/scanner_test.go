package lexer_test

import (
	"strings"
	"testing"

	"github.com/TobiEiss/jill/lexer"
)

func TestScanner(t *testing.T) {
	// Testdata
	var tests = []struct {
		sequence string
		token    lexer.Token
		expected string
	}{
		// Special tokens (EOF, ILLEGAL, WHITESPACE)
		{sequence: ``, token: lexer.EOF},
		{`#`, lexer.ILLEGAL, `#`},
		{` `, lexer.WHITESPACE, " "},
		{"\t", lexer.WHITESPACE, "\t"},
		{"\n", lexer.WHITESPACE, "\n"},

		// misc chars
		{"(", lexer.BRACKETLEFT, "("},
		{")", lexer.BRACKETRIGHT, ")"},

		// keywords
		{"FROM", lexer.FROM, "FROM"},
		{"SUM", lexer.SUM, "SUM"},

		// Identifiers
		{`foo`, lexer.IDENT, `foo`},
		{`abc_def-123`, lexer.IDENT, `abc_def`},
	}

	// run test
	for i, test := range tests {
		scanner := lexer.NewScanner(strings.NewReader(test.sequence))
		token, literal := scanner.Scan()
		if test.token != token {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, test.sequence, test.token, token, literal)
		} else if test.expected != literal {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, test.sequence, test.expected, literal)
		}
	}
}
