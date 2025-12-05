// This is how u say hello world in golang in web interface

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello u have send the request to the pasth", r.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
