package greetings

import (
	"fmt"

	"golang.org/x/text/message"
)

func hello(name string) string {
	// var message string
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}