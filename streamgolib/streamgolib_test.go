package streamgolib

import (
	"regexp"
	"testing"
)

func TestProductName(t *testing.T) {
	name := "streamgolib"
	want := regexp.MustCompile("streamgolib")
	SetProductName(name)
	msg := ProductName()
	if !want.MatchString(msg) {
		t.Fatalf(`ProductName = %q, want match for %#q`, msg, want)
	} else {
		t.Logf(`ProductName = %q, want match for %#q`, msg, want)
	}
}
