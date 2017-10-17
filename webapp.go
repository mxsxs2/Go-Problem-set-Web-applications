//Go guessing game webapp by Krisztian Nagy
package main

import (
	"net/http"
)

func main() {
	//Add the handler function for the landing page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	//Start a webserver which listens at port 8080
	http.ListenAndServe(":8080", nil)
}
