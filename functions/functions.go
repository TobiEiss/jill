package functions

import (
	"github.com/TobiEiss/jill/lexer"
)

type Function interface {
	Float64(args ...float64) float64
}

// FunctionsMap holds all functions for all lexer.Token
var FunctionsMap = map[lexer.Token]Function{
	lexer.ADD: newAdd(),
}
