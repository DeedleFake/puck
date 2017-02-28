package util

import (
	"reflect"
	"testing"
)

func TestExpandStruct(t *testing.T) {
	type test struct {
		One string `expand:"one"`
		Two string `expand:"two"`
	}

	tests := []struct {
		name  string
		in    interface{}
		out   interface{}
		panic bool
	}{
		{
			name: "Simple",
			in:   &test{One: "two: ${two}", Two: "two"},
			out:  &test{One: "two: two", Two: "two"},
		},
		{
			name:  "Simple Recursion",
			in:    &test{One: "${two}", Two: "${one}"},
			out:   nil,
			panic: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if test.panic && (recover() == nil) {
					t.Fatalf("Expected test to panic.")
				}
			}()

			out := ExpandStruct(test.in, "expand")
			if !reflect.DeepEqual(out, test.out) {
				t.Fatalf("Expected %#v\nGot %#v", test.out, out)
			}
		})
	}
}
