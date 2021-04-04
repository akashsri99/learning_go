package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris", "English")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("Mismatch")
	}

}

func TestHello2(t *testing.T) {

	assertCorrectMsg := func(t testing.TB, got, want string) {

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying Hello to people if name is supplied", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertCorrectMsg(t, got, want)

	})

	t.Run("if name is not passed, return Hello, World", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMsg(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMsg(t, got, want)
	})
}
