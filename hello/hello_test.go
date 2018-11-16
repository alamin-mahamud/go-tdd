package hello

import "testing"

func testEqual(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("got '%s' want '%s'\n", got, want)
	}
}

func TestHelloWithNoParam(t *testing.T) {
	got := Hello()
	want := "Hello, World"
	testEqual(got, want, t)
}

func TestHelloWithSingleParam(t *testing.T) {
	got := Hello("Alamin")
	want := "Hello, Alamin"
	testEqual(got, want, t)
}

func TestHelloWithLotsOfParams(t *testing.T) {
	got := Hello("Alamin", "Billal", "Chowdhury", "Drake", "Eminem", "Falcon")
	want := "Hello, Alamin, Billal, Chowdhury, Drake, Eminem, Falcon"
	testEqual(got, want, t)
}
