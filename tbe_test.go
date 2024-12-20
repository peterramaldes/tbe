package tbe_test

import (
	"fmt"
	"testing"

	"github.com/peterramaldes/tbe"
)

var stst = []struct {
	input  string
	output string
}{
	{"foo", "foo"},
	{"bar", "bar"},
	{"baz", "baz"},
}

func TestSuccessFoo(t *testing.T) {
	for _, e := range stst {
		result, err := tbe.Foo(e.input)
		if err != nil {
			t.Fatal(err)
		}

		if result != e.output {
			t.Errorf("Foo(%v) = %v, want %v", e.input, result, e.output)
		}
	}
}

var ftst = []struct {
	input string
}{
	{"ti"},
	{"boo"},
	{"baa"},
	{"iii"},
}

func TestFailedFoo(t *testing.T) {
	for _, e := range ftst {
		_, err := tbe.Foo(e.input)
		if err == nil {
			t.Fatal("expecting an error for not allowed word")
		}

		errMsg := fmt.Sprintf("%s is not allowed", e.input)
		if err.Error() != errMsg {
			t.Fatalf("expected %s but %s", errMsg, err)
		}
	}
}
