package testmod

import (
	testmod "github.com/pehks1980/testmod"
	testmodv2 "github.com/pehks1980/testmod/v2"
	"testing"
)

func TestHi(t *testing.T) {
	want := "Hi, Murlitas"
	if got := testmod.Hi("Murlitas"); got != want {
		t.Errorf("Hi() = %q, want %q", got, want)
	}
}

func TestHiv2(t *testing.T) {
	want := "Hi, Murlitas"
	if got := testmodv2.Hi("Murlitas"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
	want = "Hello, Anyone"
	if got := testmodv2.Hi("Anyone"); got != want {
		t.Errorf("Hi() = %q, want %q", got, want)
	}
}

