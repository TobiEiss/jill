package lexer_test

import (
	"testing"

	"github.com/TobiEiss/jill/lexer"
)

func TestLegalStatements(t *testing.T) {
	var tests = []struct {
		Statement string
		ErrorType lexer.ErrorType
	}{
		// legal cases
		{Statement: `SUM ( json1.field1, json2.field2 )`},
		{Statement: `SUM ( json1.field1, json2.field2, json3.field1 )`},
		{Statement: `SUM ( json1.field1, json2.field2, SUM ( json3.field1, json3.field2 ) )`},
		{Statement: `SUM ( json1.field1, json2.field2, SUM ( json3.field2 ) )`},
		{Statement: `SUM(json1.field1,json2.field2,SUM(json3.field2))`},
		{Statement: `SUM ( json1.field1, json2.field2, ADD ( json3.field2 ) )`},
		{Statement: `ADD ( ADD ( json3.field2 ) )`},

		// illegal cases
		{"SUMM ( json1.field1, json2.field2 )", lexer.ILLEGALTOKEN},
		{"SUM ( )", lexer.MISSINGARGUMENT},
		{"SUM", lexer.ILLEGALTOKEN},
	}

	// iterate all tests
	for index, test := range tests {
		_, err := lexer.NewParser(test.Statement).ParseStatement()
		if lexerError, ok := err.(*lexer.Error); ok {
			if lexerError.ErrorType != test.ErrorType {
				t.Errorf("%d failed: Not expected error: %s", index, err)
			}
			// error was expected
		}
	}
}
