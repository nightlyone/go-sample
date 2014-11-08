package main

import (
	"testing"
)

func TestTruth(t *testing.T) {
	want := true
	got := true

	if got == want {
		t.Log("got", got)
	} else {
		t.Error("want:", want, "but got", got)
	}
}
