package main

import (
	"testing"
)

func TestAddArray(t *testing.T) {
	array := []string{"1", "2"}
    got := AddArray(array)
    if got != 3 {
        t.Errorf("AddArray[1, 2] = %d; want 3", got)
    }
}