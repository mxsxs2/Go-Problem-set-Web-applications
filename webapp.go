//Go guessing game webapp by Krisztian Nagy
package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Add the handler function
	http.HandleFunc("/", requestHandler)
	//Start a webserver which listens at port 8080
	http.ListenAndServe(":8080", nil)
}

//Handle the request
func requestHandler(w http.ResponseWriter, r *http.Request) {
	//Set the html hconent type
	w.Header().Set("Content-Type", "text/html")

	//Write the response
	fmt.Fprintln(w, "<h1>Guessing game</h1>")

}
