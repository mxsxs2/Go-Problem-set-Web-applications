//Go guessing game webapp by Krisztian Nagy
package main

import (
	"net/http"
)

func main() {
	//Add the handler function
	http.Handle("/", http.FileServer(http.Dir("guess")))
	//Start a webserver which listens at port 8080
	http.ListenAndServe(":8080", nil)
}
