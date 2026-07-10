package main

import (
	"testing"
)

func TestDemo(t *testing.T) {
	result := "hello"
	if result != "hello" {
		t.Error("test failed")
	}
}
