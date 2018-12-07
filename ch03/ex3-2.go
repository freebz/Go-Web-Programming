// Listing 3.2  Web server with additional configuration

package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
