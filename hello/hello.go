package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	fmt.Println("Hello world")
	message := greetings.Hello("Baltinha")
	fmt.Println(message)
}
