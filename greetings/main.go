package main

import (
	"fmt"

	"golang.org/x/text/message"
)

func Hello(name string) string {
	// var message string
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}