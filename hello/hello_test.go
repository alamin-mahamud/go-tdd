package hello

import "testing"

func TestHelloWithNoParam(t *testing.T) {
	got := Hello()
	want := "Hello, World"
	if got != want {
		t.Errorf("got '%s' want '%s'\n", got, want)
	}
}

func TestHelloWithSingleParam(t *testing.T) {
	got := Hello("Alamin")
	want := "Hello, Alamin"
	if got != want {
		t.Errorf("got '%s' want '%s'\n", got, want)
	}
}

func TestHelloWithLotsOfParam(t *testing.T) {

}
