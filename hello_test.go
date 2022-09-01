package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
			got := Hello("Rider")
			want := "hello, Rider"

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
	})
	t.Run("say 'hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}