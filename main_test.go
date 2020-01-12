package main

import (
	"testing"
)

func TestIntro(t *testing.T) {
	want := "simpledb (0.0.1)"
	if got := Intro(); got != want {
		t.Errorf("Intro() = %q, want %q", got, want)
	}
}
