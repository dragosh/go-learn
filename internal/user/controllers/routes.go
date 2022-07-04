package controllers

import "net/http"

func RegisterControllers() {

	// Interface example
	// newUserController implements http.Handle
	// so it must have	ServeHTTP as return method
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
