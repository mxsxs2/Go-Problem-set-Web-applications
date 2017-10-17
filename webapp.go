//Go guessing game webapp by Krisztian Nagy
package main

import (
	"bytes"
	"errors"
	"fmt"
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
	Message string //The message
	Guess   int    //The guessed number
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

	//Create the message
	message := Message{
		Message: "Guess a number between 1 and 20",
	}

	//Check if there is any guess was made
	if guess, err := getGuessedNumberParameter(r); err == nil {
		//Set the guessed number if there was any
		message.Guess = guess
	}

	//Merge the template with the message
	gTemplate.Execute(w, message)

	//getRequestDetails(w, r)

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

//Function used to get the guessed number form the url parameters
func getGuessedNumberParameter(r *http.Request) (int, error) {
	//Get the guess parameters. It might be more than one
	guessparam := r.URL.Query().Get("guess")
	//Check if it is empty
	if guessparam != "" {
		//Try to parse it
		if guess, err := strconv.Atoi(guessparam); err == nil {
			//Return the number and no error
			return int(guess), nil
		}
		//Return the invalid input error
		return 0, errors.New("Invalid input")
	}
	//Return the no parameter present error
	return 0, errors.New("No parameter present")
}

func getRequestDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "r.Method:           ", r.Method)
	fmt.Fprintln(w, "r.URL:              ", r.URL)
	fmt.Fprintln(w, "r.Proto:            ", r.Proto)
	fmt.Fprintln(w, "r.ContentLength:    ", r.ContentLength)
	fmt.Fprintln(w, "r.TransferEncoding: ", r.TransferEncoding)
	fmt.Fprintln(w, "r.Close:            ", r.Close)
	fmt.Fprintln(w, "r.Host:             ", r.Host)
	fmt.Fprintln(w, "r.Form:             ", r.Form)
	fmt.Fprintln(w, "r.PostForm:         ", r.PostForm)
	fmt.Fprintln(w, "r.RemoteAddr:       ", r.RemoteAddr)
	fmt.Fprintln(w, "r.RequestURI:       ", r.RequestURI)

	fmt.Fprintln(w, "r.URL.Opaque:       ", r.URL.Opaque)
	fmt.Fprintln(w, "r.URL.Scheme:       ", r.URL.Scheme)
	fmt.Fprintln(w, "r.URL.Host:         ", r.URL.Host)
	fmt.Fprintln(w, "r.URL.Path:         ", r.URL.Path)
	fmt.Fprintln(w, "r.URL.RawPath:      ", r.URL.RawPath)
	fmt.Fprintln(w, "r.URL.RawQuery:     ", r.URL.RawQuery)
	fmt.Fprintln(w, "r.URL.Fragment:     ", r.URL.Fragment)

	fmt.Fprintln(w, "Header:")
	for key, value := range r.Header {
		fmt.Fprintln(w, "\t"+key+":", value)
	}

	body := new(bytes.Buffer)
	body.ReadFrom(r.Body)

	fmt.Fprintln(w, "r.Body:             ", body.String())
}
