package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
			got := Hello("Rider", "")
			want := "hello, Rider"

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
	})

	t.Run("say 'hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "hello, World"
		assertCorrectMessage(t, got, want) // because we called t.Helper(), the test suite tells us the issue is on this line, instead of in the helper function
	})

	t.Run("in Spanish", func(t *testing.T)  {
		got := Hello("Rider", "Spanish")
		want := "Hola, Rider"
		assertCorrectMessage(t, got, want)
	})
}

// helper functions take in testing.TB => which means it can be used in a test or benchmark
func assertCorrectMessage(t testing.TB, got, want string) { 
	t.Helper() // we need to call this to let the test
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}