package functions

type add struct{}

func newAdd() *add {
	return &add{}
}

// Float64 implements the add function for numbers
func (add *add) Float64(args ...float64) float64 {
	var result float64
	for _, value := range args {
		result += value
	}
	return result
}
