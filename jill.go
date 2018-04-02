package jill

import (
	"github.com/Jeffail/gabs"
	"github.com/TobiEiss/jill/lexer"
)

// Container holds the data.
// Currently only JSON (gabs.Container). Later more datatypes.
type Container struct {
	JSONContainer *gabs.Container
}

// ParseJSON is the function to get a Container.
func ParseJSON(content []byte) (*Container, error) {
	jsonContainer, err := gabs.ParseJSON(content)
	return &Container{JSONContainer: jsonContainer}, err
}

// Query the container
func (container *Container) Query(query string) (string, error) {
	_, err := lexer.NewParser(query).Parse()
	return "sucessfully", err
}
