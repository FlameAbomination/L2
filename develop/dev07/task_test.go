package main

import (
	"testing"
	"time"
)

func TestDict(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)
	local := time.Since(start)
	if local > 6*time.Second {
		t.Errorf("difference between %s %s is bigger than a second", start, local)
	}
}
