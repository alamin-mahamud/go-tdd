package locale

import "testing"

func TestHelloLocale(t *testing.T) {

	var assertEqual = func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'\n", got, want)
		}
	}

	t.Run("saying hello to one man", func(t *testing.T) {
		got := Hello("Alamin", "English")
		want := "Hello, Alamin"
		assertEqual(t, got, want)
	})

	t.Run("saying hello with an empty string", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertEqual(t, got, want)
	})

	t.Run("saying hello with an empty string and also no lang specified", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertEqual(t, got, want)
	})

	t.Run("saying hello for an unknown string", func(t *testing.T) {
		got := Hello("", "12321312")
		want := "Unknown Language. Try English"
		assertEqual(t, got, want)
	})

	t.Run("saying hello in english", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertEqual(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"
		assertEqual(t, got, want)
	})
}
