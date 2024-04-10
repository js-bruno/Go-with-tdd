package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("say hello in french", func(t *testing.T) {
		result := Hello("bruno!", "french")
		expected := "bonjour, bruno!"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("say hello in english", func(t *testing.T) {
		result := Hello("Chris!", "")
		expected := "hello, Chris!"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("say 'Hello World!!!' when string passed is empty", func(t *testing.T) {
		result := Hello("", "")
		expected := "hello, World!!!"

		assertCorrectMessage(t, result, expected)
	})
}

