package main

import "testing"

func TestGreet(t *testing.T) {
	if got := Greet("Вася"); got != "Привет, Вася!" {
		t.Errorf("got %q", got)
	}
	if got := Greet(""); got != "Привет, мир!" {
		t.Errorf("got %q", got)
	}
}
