package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("John Snow")
	want := "Hello, John Snow"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
