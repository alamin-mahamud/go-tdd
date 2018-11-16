package hello

import "testing"

func TestHello(t *testing.T) {

	var testEqual = func(got, want string, t *testing.T) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'\n", got, want)
		}
	}

	t.Run("saying hello to one man", func(t *testing.T) {
		got := Hello("Alamin")
		want := "Hello, Alamin"
		testEqual(got, want, t)
	})

	t.Run("saying without passing anything", func(t *testing.T) {
		got := Hello()
		want := "Hello, World"
		testEqual(got, want, t)
	})

	t.Run("saying hello to lots of people by name", func(t *testing.T) {
		got := Hello("Alamin", "Billal", "Chowdhury", "Drake", "Eminem", "Falcon")
		want := "Hello, Alamin, Billal, Chowdhury, Drake, Eminem, Falcon"
		testEqual(got, want, t)
	})

	t.Run("saying hello with an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		testEqual(got, want, t)
	})
}
