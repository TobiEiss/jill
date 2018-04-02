package lexer

// Token represents a lexical token.
type Token int

const (
	ILLEGAL Token = iota
	EOF
	WHITESPACE

	IDENT // identifier

	BRACKETLEFT  // (
	BRACKETRIGHT // )
	COMMA        // ,

	SUM
	ADD
)

// eof represents a marker rune for the end of the reader.
var eof = rune(0)

// MiscCharMap is a map from the rune to the Token
var MiscCharMap = map[rune]Token{
	'(': BRACKETLEFT,
	')': BRACKETRIGHT,
	',': COMMA,
}

// KeyWordMap is a map from the string to the Token
var KeyWordMap = map[string]Token{
	"SUM": SUM,
	"ADD": ADD,
}

func isTokenAKeyWord(token Token) bool {
	for _, value := range KeyWordMap {
		if value == token {
			return true
		}
	}
	return false
}
