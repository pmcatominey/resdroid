package main

import (
	"testing"
)

func TestInvalidResPath(t *testing.T) {
	r, err := NewResDirectory("invalidpath/that/doesnt/exist")

	if err == nil {
		t.Error("Invalid res path should result in an error")
	}

	if r != nil {
		t.Error("r (ResDirectory) should be nil with invalid path")
	}
}
