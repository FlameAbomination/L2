package main

import (
	"reflect"
	"testing"
)

type Test struct {
	set      []string
	expected map[string][]string
}

var Tests = []Test{
	{[]string{
		"пятак", "пятка", "тяпка",
		"листок", "слиток", "столик",
	},
		map[string][]string{
			"пятак":  {"пятка", "тяпка"},
			"листок": {"слиток", "столик"},
		},
	},
}

func TestDict(t *testing.T) {
	for _, test := range Tests {
		if output := getDictionary(test.set); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q or", output, test.expected)
		}
	}
}
