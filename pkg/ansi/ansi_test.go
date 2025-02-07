package ansi

import (
	"testing"
)

func TestNew(t *testing.T) {
	find := New(1,2,3)
	expected := AnsiCode{1,2,3}
	if find != expected {
		t.Errorf("Error: found %v, expected %x", find, expected)
	}
}
