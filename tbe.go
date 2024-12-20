package tbe

import (
	"fmt"
	"slices"
)

func Foo(input string) (string, error) {
	allowed := []string{"foo", "bar", "baz"}

	if !slices.Contains(allowed, input) {
		err := fmt.Errorf("%v is not allowed", input)
		return "", err
	}

	return input, nil
}
