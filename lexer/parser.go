package lexer

import (
	"fmt"
	"io"
)

// Statement represents jill-statement
type Statement struct {
	Function   Token
	Fields     []string
	Statements []*Statement
}

// Parser represents a parser.
type Parser struct {
	scanner *Scanner
	buf     struct {
		token      Token  // last read token
		literal    string // last read literal
		buffersize int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(reader io.Reader) *Parser {
	return &Parser{scanner: NewScanner(reader)}
}

// Parse parses a jill statement.
func (parser *Parser) Parse() (*Statement, error) {
	token, literal := parser.scanIgnoreWhitespace()
	if token != SUM {
		return nil, fmt.Errorf("found %q, expected field", literal)
	}

	return findStatement(parser, token)
}

func findStatement(parser *Parser, function Token) (*Statement, error) {
	stmt := &Statement{Function: function, Fields: []string{}, Statements: []*Statement{}}
	// read a field
	token, literal := parser.scanIgnoreWhitespace()

	// check if token is legal
	if token == ILLEGAL || token != BRACKETLEFT {
		return nil, fmt.Errorf("found %q, expected field", literal)
	}

	// collect fields
	for {
		token, literal := parser.scanIgnoreWhitespace()

		// if token is an ident add statement with function "IDENT"
		if token == IDENT {
			stmt.Fields = append(stmt.Fields, literal)
		} else if token == SUM { // if there is a new function:
			innerStmt, err := findStatement(parser, token)
			if err != nil {
				return nil, err
			}
			stmt.Statements = append(stmt.Statements, innerStmt)
		} else {
			return nil, fmt.Errorf("missing argument")
		}

		// next have to be COMMA or BRACKETRIGHT
		token, literal = parser.scanIgnoreWhitespace()

		if token == BRACKETRIGHT {
			break
		} else if token != COMMA {
			return nil, fmt.Errorf("found %q, expected field", literal)
		}
	}
	return stmt, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (parser *Parser) scan() (token Token, literal string) {
	// If we have a token on the buffer, then return it.
	if parser.buf.buffersize != 0 {
		parser.buf.buffersize = 0
		return parser.buf.token, parser.buf.literal
	}

	// Otherwise read the next token from the scanner.
	token, literal = parser.scanner.Scan()

	// Save it to the buffer in case we unscan later.
	parser.buf.token, parser.buf.literal = token, literal
	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (parser *Parser) scanIgnoreWhitespace() (token Token, literal string) {
	token, literal = parser.scan()
	if token == WHITESPACE {
		token, literal = parser.scan()
	}
	return
}

// unscan pushes the previously read token back onto the buffer.
func (parser *Parser) unscan() {
	parser.buf.buffersize = 1
}
