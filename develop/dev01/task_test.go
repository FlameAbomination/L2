package main

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	current, err := GetTime()
	local := time.Now()
	if err != nil {
		t.Errorf("got %v", err)
	}
	if current.Sub(local) > time.Second {
		t.Errorf("difference between %s %s is bigger than a second", current, local)
	}
}
