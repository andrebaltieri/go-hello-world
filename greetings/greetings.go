package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi %vi. Welcome", name)
	return message
}
