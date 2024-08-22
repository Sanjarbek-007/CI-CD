package main

import "testing"

func TestSayHello(t *testing.T) {
	name := "Sanjarbek"
    expected := "Hello, Sanjar!"

    actual := SayHello(name)

    if actual != expected {
        t.Errorf("Expected '%s', got '%s'", expected, actual)
    }
}