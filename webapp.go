//Go guessing game webapp by Krisztian Nagy
package main

import (
	"html/template"
	"net/http"
)

//Struct for the template
type Message struct {
	Message string
}

func main() {

	//Add the handler function for the landing page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Set the html hconent type
		w.Header().Set("Content-Type", "text/html")
		//Serve the index file
		http.ServeFile(w, r, "index.html")
	})
	//Add the handler function for the guess folder
	http.HandleFunc("/guess/", func(w http.ResponseWriter, r *http.Request) {
		//Set the html hconent type
		w.Header().Set("Content-Type", "text/html")
		//Parse the template file
		gTemplate := template.Must(template.ParseFiles("guess/guess.tmpl"))
		//Merge the template with the message
		gTemplate.Execute(w, Message{Message: "Guess a number between 1 and 20"})

		//http.ServeFile(w, r, "guess/guess.html")
	})
	//Start a webserver which listens at port 8080
	http.ListenAndServe(":8080", nil)
}
