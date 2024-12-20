package tbe_test

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/peterramaldes/tbe"
)

func TestExhaustive(t *testing.T) {
	data := make([]string, 999)
	for i := range data {
		data[i] = generateRandomString(10)
	}

	for i := 0; i < len(data); i++ {
		v := data[i]
		_, err := tbe.Foo(v)
		if err == nil {
			t.Fatalf("expected an error but none error received")
		}

		errMsg := fmt.Sprintf("%s is not allowed", v)
		if err.Error() != errMsg {
			t.Fatalf("expected %s but %s", errMsg, err)
		}
	}

}

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

var stst = []struct {
	input  string
	output string
}{
	{"foo", "foo"},
	{"bar", "bar"},
	{"baz", "baz"},
	{"Brasil", "Brasil"},
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
