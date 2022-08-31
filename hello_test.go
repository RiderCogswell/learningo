package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Rider")
	want := "hello, Rider"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}