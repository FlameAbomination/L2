package main

import (
	"errors"
	"testing"
)

type Test struct {
	packed   string
	expected string
	err      error
}

var Tests = []Test{
	{"a4bc2d5e", "aaaabccddddde", nil},
	{"abcd", "abcd", nil},
	{"45", "", errors.New("incorrect string")},
	{"", "", nil},
	{"qwe\\4\\5", "qwe45", nil},
	{"qwe\\45", "qwe44444", nil},
	{"qwe\\\\5", "qwe\\\\\\\\\\", nil},
}

func TestUnpack(t *testing.T) {
	for _, test := range Tests {
		if output, _ := UnpackString(test.packed); output != test.expected {
			t.Errorf("Output %q not equal to expected %q or", output, test.expected)
		}
	}
}
