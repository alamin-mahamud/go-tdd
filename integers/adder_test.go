package integers

import "testing"

func TestAdder(t *testing.T) {

	var assertEqual = func(t *testing.T, want, got int) {
		t.Helper()

		if got != want {
			t.Errorf("EXPECTED: '%d', GOT: '%d'", want, got)
		}
	}

	t.Run("2+2 should be 4", func(t *testing.T) {
		want := 4
		got := Adder(2, 2)
		assertEqual(t, want, got)
	})
}
