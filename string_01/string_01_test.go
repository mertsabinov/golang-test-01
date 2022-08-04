package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("Name")
	want := "Hello Name"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
