package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}


// Constructor example
//  func new.....() pointerTo
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

// Method example
// func(type) name (args)
func (uc userController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// res.Write([]byte("Hello from the User Controller"))
	if req.URL.Path == '/users' {
		switch req.Method {
		case http.MethodGet:
				uc.getAll(res, req)
				// ..... other use cases
		default:
				res.Write(http.StatusNotImplemented)
		}
	} else {
		// handle the user id regexp patter
	}
}
