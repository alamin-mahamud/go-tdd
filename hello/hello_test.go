package hello

import "testing"

func testEqual(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("got '%s' want '%s'\n", got, want)
	}
}

func TestHello(t *testing.T) {
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
