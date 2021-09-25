package server

import (
	"fmt"
	"html"
	"net/http"
)

// CREATE A HTTP SERVER

// http.ResponseWriter assembles the servers response and writes to
// the client
// http.Request is the clients request
func root(w http.ResponseWriter, r *http.Request) {
	// Writes to the client
	fmt.Fprintf(w, "Home, %q", html.EscapeString(r.URL.Path))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

}

func Start() {

	// Calls for function handlers output to match the directory /
	http.HandleFunc("/", root)

	// Calls for function handler2 output to match directory /earth
	http.HandleFunc("/hello", hello)

	// Listen to port 8080 and handle requests
	http.ListenAndServe(":8080", nil)

}
