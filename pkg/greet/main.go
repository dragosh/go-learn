package greet

import (
	"strings"
)

// Private visibility
var name strings.Builder

// Public visibility
var Name = "World"

// Private visibility
func hello() string {
	name.WriteString("Hello")
	name.WriteString(" ")
	return name.String()
}

// Public visibility
func Hello(n string) string {
	return hello() + n
}
