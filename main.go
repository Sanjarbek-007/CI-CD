package main

import "fmt"

func main() {
	msg := SayHello("John Doe")

	println(msg)
}

func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
