// Listing 3.4  Serving through HTTPS

package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
