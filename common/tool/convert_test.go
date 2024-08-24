package tool

import (
	"testing"
)

func TestBytesToString(t *testing.T) {
	original := []byte("Hello, World!")
	converted := BytesToString(original)
	expected := "Hello, World!"

	if converted != expected {
		t.Errorf("BytesToString failed: expected %s, got %s", expected, converted)
	}
}

func TestStringToBytes(t *testing.T) {
	original := "Hello, World!"
	converted := StringToBytes(original)
	expected := []byte("Hello, World!")

	if string(converted) != string(expected) {
		t.Errorf("StringToBytes failed: expected %s, got %s", expected, string(converted))
	}
}
