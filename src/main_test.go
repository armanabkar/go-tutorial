package main

import "testing"

// To make a test, create a new file ending in _test.go
// To run all tests in a package, use "go test"
// What do you really care about in the project? Test that thing!

func TestExample(t *testing.T) {
	name := "Arman"

	if name != "Arman" {
		t.Errorf("Expected name to be Arman, But got %v", name)
	}
}
