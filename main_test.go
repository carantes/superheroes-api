package main

import "testing"

func testHello(t *testing.T) {
	expected := "API is running"
	result := startServer()

	if result != expected {
		t.Errorf("Hello function return error, expected %q, got %q", expected, result)
	}
}
