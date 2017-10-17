//Go guessing game webapp by Krisztian Nagy
package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//The boundaries of the random number generator
const RANDMAX = 20
const RANDMIN = 10

//Struct for the template
type Message struct {
	Message string
}

func main() {
	//Seet the random generator once the application is started
	rand.Seed(time.Now().UnixNano())

	//Add the handler function for the landing page
	http.HandleFunc("/", landingPageHandler)
	//Add the handler function for the guess folder
	http.HandleFunc("/guess/", guessPageHandler)
	//Start a webserver which listens at port 8080
	http.ListenAndServe(":8080", nil)
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	//Set the html hconent type
	w.Header().Set("Content-Type", "text/html")
	//Serve the index file
	http.ServeFile(w, r, "index.html")
}

func guessPageHandler(w http.ResponseWriter, r *http.Request) {
	//Check if the cookie exist, if not then create it
	checkAndSetTargetCookie(w, r)

	//Set the html hconent type
	w.Header().Set("Content-Type", "text/html")
	//Parse the template file
	gTemplate := template.Must(template.ParseFiles("guess/guess.tmpl"))
	//Merge the template with the message
	gTemplate.Execute(w, Message{Message: "Guess a number between 1 and 20"})

	//http.ServeFile(w, r, "guess/guess.html")
}

func checkAndSetTargetCookie(w http.ResponseWriter, r *http.Request) {
	//Get the cookie
	cookie, err := r.Cookie("target")

	//Check if cookie doest not exists
	if err != nil || cookie == nil {

		//Create the random number
		randomNumber := rand.Intn(RANDMAX-RANDMIN) + RANDMIN
		//Create the cookie
		cookie := http.Cookie{Name: "target", Value: strconv.Itoa(randomNumber), Expires: time.Now().Add(365 * 24 * time.Hour)}
		//Set the cookie
		http.SetCookie(w, &cookie)
	}
}
