package tbe_test

import (
	"testing"

	"github.com/peterramaldes/tbe"
)

func TestFoo(t *testing.T) {
	if err := tbe.Foo(); err != nil {
		t.Fatal(err)
	}
}
