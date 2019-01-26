package jill

import (
	"fmt"
	"reflect"

	"github.com/Jeffail/gabs"
	"github.com/TobiEiss/jill/functions"
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
func (container *Container) Query(query string) (interface{}, error) {
	stmt, err := lexer.NewParser(query).ParseStatement()
	if err != nil {
		return nil, err
	}
	return container.apply(stmt)
}

// apply calculates a statement
func (container *Container) apply(stmt *lexer.Statement) (interface{}, error) {
	type result struct {
		Value interface{}
		Type  reflect.Type
	}
	results := []result{}

	// collect all results of all fields
	for _, field := range stmt.Fields {
		value := container.JSONContainer.Path(field).Data()
		results = append(results, result{Value: value, Type: reflect.TypeOf(value)})
	}

	// collect all results of all statments
	for _, innerStmt := range stmt.Statements {
		value, err := container.apply(innerStmt)
		if err != nil {
			return nil, err
		}
		results = append(results, result{Value: value, Type: reflect.TypeOf(value)})
	}

	// check if all datatypes are the same
	for i := 0; i < len(results); i++ {
		if results[i].Type != results[0].Type {
			// TODO improve error-message
			return nil, fmt.Errorf("Value %s and Value %s are not the same datatype",
				results[i].Value, results[0].Value)
		}
	}

	// choose correct function
	function := functions.FunctionsMap[stmt.Function]
	switch results[0].Type.Kind() {
	case reflect.Float64:
		number1 := results[0].Value.(float64)
		number2 := results[1].Value.(float64)
		return function.Float64(number1, number2), nil
	}
	return "successfully", nil
}
