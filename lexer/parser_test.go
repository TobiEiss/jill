package lexer_test

import (
	"strings"
	"testing"

	"github.com/TobiEiss/jill/lexer"
)

func TestLegalStatements(t *testing.T) {
	var tests = []struct {
		Statement string
		legal     bool
	}{
		// legal cases
		{"SUM ( json1.field1, json2.field2 )", true},
		{"SUM ( json1.field1, json2.field2, json3.field1 )", true},
		{"SUM ( json1.field1, json2.field2, SUM ( json3.field1, json3.field2 ) )", true},
		{"SUM ( json1.field1, json2.field2, SUM ( json3.field2 ) )", true},
		{"SUM(json1.field1,json2.field2,SUM(json3.field2))", true},

		// illegal cases
		{"SUMM ( json1.field1, json2.field2 )", false},
		{"SUM ( )", false},
		{"SUM", false},
	}

	// iterate all tests
	for index, test := range tests {
		_, err := lexer.NewParser(strings.NewReader(test.Statement)).Parse()
		if (err != nil) == test.legal {
			t.Errorf("%d failed: %s", index, err)
		}
	}
}
