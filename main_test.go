package gotest

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add([]int{1, 2, 3})
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}
