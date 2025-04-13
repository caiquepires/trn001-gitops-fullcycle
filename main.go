package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server")

	// Register the helloHandler function to handle requests to the root URL ("/").
	// The http.HandleFunc function takes two arguments:
	// 1. The URL pattern ("/") to match incoming requests.
	// 2. The handler function (helloHandler) that will be called to handle those requests.
	// The helloHandler function is defined below and is responsible for sending a response back to the client.
	// It uses the http.ResponseWriter to write the response and the http.Request to access request details.
	// The http.HandleFunc function is part of the net/http package and is used to register HTTP handlers.
	// The http package provides a simple way to create web servers and handle HTTP requests in Go.
	// The http.HandleFunc function is a convenient way to associate a URL pattern with a handler function.
	// When a request is received that matches the specified URL pattern,
	// the corresponding handler function is called to process the request.
	http.HandleFunc("/", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		// The server will listen on port 8080 for incoming HTTP requests.
		// If the server starts successfully, it will continue running until terminated.
		// The http.ListenAndServe function blocks the main goroutine,
		// allowing the server to handle incoming requests concurrently.
		// The nil argument indicates that we are using the default ServeMux for routing requests.
		// The default ServeMux is a multiplexer that matches incoming requests to registered handlers.

	}
	// The log.Fatalf function is used to log an error message and exit the program if the server fails to start.

}

// helloHandler responds with a simple "Hello, World!" message.
// It is registered to handle requests to the root URL ("/").
// The function takes an http.ResponseWriter and an http.Request as parameters.
// The http.ResponseWriter is used to send a response back to the client,
// and the http.Request contains information about the incoming request.
// In this case, we simply write "Hello, World!" to the response writer.
// This will be displayed in the client's web browser when they access the root URL.
// The server listens on port 8080, and when a request is received,
// the helloHandler function is called to handle it.
// helloHandler is a simple HTTP handler function that responds with "Hello, World!".
// It is registered to handle requests to the root URL ("/").
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
