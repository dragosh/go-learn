package sample

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello %v!\n", name)
}

func depart(name string) string {
	return fmt.Sprintf("Good bye %v!\n", name)
}
