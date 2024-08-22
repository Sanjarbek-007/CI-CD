package main

import "testing"

func TestSayHello(t *testing.T) {
	name := "John Doe"
    expected := "Hello, John Doe!"

    actual := SayHello(name)

    if actual != expected {
        t.Errorf("Expected '%s', got '%s'", expected, actual)
    }
}