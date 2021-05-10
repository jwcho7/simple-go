package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "Helloworld"

	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func TestGoodbye(t *testing.T) {
	got := goodbye()
	want := "goodbye"

	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}
