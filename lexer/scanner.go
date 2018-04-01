package lexer

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	reader *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{reader: bufio.NewReader(reader)}
}

// Scan returns the next token and literal value.
func (scanner *Scanner) Scan() (token Token, literal string) {
	// Read the next rune.
	character := scanner.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	// If we see a digit then consume as a number.
	if isWhitespace(character) {
		scanner.unread()
		return scanner.scanWhitespace()
	} else if isLetter(character) {
		scanner.unread()
		return scanner.scanIdent()
	}

	// Otherwise read the individual character.
	if character == eof {
		return EOF, ""
	}
	if token, ok := MiscCharMap[character]; ok {
		return token, string(character)
	}

	return ILLEGAL, string(character)
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (scanner *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(scanner.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if character := scanner.read(); character == eof {
			break
		} else if !isLetter(character) && !isDigit(character) && character != '_' {
			scanner.unread()
			break
		} else {
			_, _ = buf.WriteRune(character)
		}
	}

	// If the string matches a keyword then return that keyword.
	if token, ok := KeyWordMap[strings.ToUpper(buf.String())]; ok {
		return token, buf.String()
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (scanner *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(scanner.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if character := scanner.read(); character == eof {
			break
		} else if !isWhitespace(character) {
			scanner.unread()
			break
		} else {
			buf.WriteRune(character)
		}
	}

	return WHITESPACE, buf.String()
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (scanner *Scanner) read() rune {
	character, _, err := scanner.reader.ReadRune()
	if err != nil {
		return eof
	}
	return character
}

// unread places the previously read rune back on the reader.
func (scanner *Scanner) unread() {
	_ = scanner.reader.UnreadRune()
}

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}
