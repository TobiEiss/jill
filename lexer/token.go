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
}
