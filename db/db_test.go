package db

import (
	"regexp"
	"testing"
	"time"
)

func TestGenerateSomeKey(t *testing.T) {
	t.Log("TestGenerateSomeKey")
	key := GenerateSomeKey("1234", "title", time.Now())
	t.Logf(`GeneratedSomeKey = %q`, key)
}

func TestParseSomeKey(t *testing.T) {
	t.Log("TestParseSomeKey")
	now := time.Now()
	want := regexp.MustCompile("/title/1234")
	key := GenerateSomeKey("1234", "title", now)
	t.Logf(`GeneratedSomeKey = %q`, key)

	decodedKey, err := ParseSomeKey(key)
	if err != nil {
		t.Fatalf(`ParseSomeKey error %q`, err)
	}
	if !want.MatchString(decodedKey) {
		t.Fatalf(`ProductName = %q, want match for %#q`, decodedKey, want)
	} else {
		t.Logf(`ProductName = %q, want match for %#q`, decodedKey, want)
	}

}
