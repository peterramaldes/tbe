package tbe

import (
	"fmt"
	"slices"
)

func Foo(input string) (string, error) {
	allowed := []string{"foo", "bar", "baz"}
	countries := []string{"Brasil"}

	isCountry := slices.Contains(countries, input)
	isAllowed := slices.Contains(allowed, input)

	if isCountry {
		return input, nil
	} else if !isAllowed {
		err := fmt.Errorf("%v is not allowed", input)
		return "", err
	}

	return input, nil
}
