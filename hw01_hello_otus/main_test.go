package main

import "testing"

func TestReversePhrase(t *testing.T) {
	got := reversePhrase("Hello, OTUS!")
	expected := "!SUTO ,olleH"

	if got != expected {
		t.Fatalf("got %q, want %q", got, expected)
	}
}
