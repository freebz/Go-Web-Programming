// Listing 3.1  The simplest web server

package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("127.0.0.1:8080", nil)
}
