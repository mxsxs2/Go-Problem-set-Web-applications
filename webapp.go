//Go guessing game webapp by Krisztian Nagy
package main

import (
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//The boundaries of the random number generator
const RANDMAX = 20
const RANDMIN = 10

//Struct for the template
type Message struct {
	Message      string //The message
	GuessMessage string //The message after the guess
	Guessed      int    //If the number was guessed
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

//Handler for the landing page
func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	//Set the html hconent type
	w.Header().Set("Content-Type", "text/html")
	//Serve the index file
	http.ServeFile(w, r, "index.html")
}

//Handler for the guessing page
func guessPageHandler(w http.ResponseWriter, r *http.Request) {
	//Check if the cookie exist, if not then create it
	checkAndSetTargetCookie(w, r)
	//Reset the cookie if a new game was started
	if checkNewGameParam(r) == 1 {
		setTargetCookie(w)
	}
	//Set the html hconent type
	w.Header().Set("Content-Type", "text/html")
	//Parse the template file
	gTemplate := template.Must(template.ParseFiles("guess/guess.tmpl"))
	//Set the messages
	message := setMessage(w, r)

	//Merge the template with the message
	gTemplate.Execute(w, message)
}

//Function used to set the messages
func setMessage(w http.ResponseWriter, r *http.Request) Message {
	//Create the message
	message := Message{
		Message: fmt.Sprintf("Guess a number between %d and %d", RANDMIN, RANDMAX),
		Guessed: 0,
	}

	//Compare the guess and the stored number
	if result, guess, err := compareGuessToCookie(r); err == nil {
		//If they equal
		if result == 0 {
			message.GuessMessage = "Congratulations! You guessed the number."
			//Set the guessed number for the new game link
			message.Guessed = 1
		}
		//If the guess was lower
		if result == -1 {
			message.GuessMessage = fmt.Sprintf("Your guess was %d which is too low.", guess)
		}
		//If the guess was higher
		if result == 1 {
			message.GuessMessage = fmt.Sprintf("Your guess was %d which is too high.", guess)
		}
	}
	//Return the message
	return message
}

//Function used to compare the guess and the stored cookie
func compareGuessToCookie(r *http.Request) (int, int, error) {
	//Holder for the values
	var guess int
	var cookieValue int
	//Get the value of the cookie
	if value, err := getTargetCookieValue(r); err == nil {
		//Set the value
		cookieValue = value
	} else {
		return 0, 0, errors.New("No cookie set")
	}

	//Get the value of the guess
	if g, err := getGuessedNumberParameter(r); err == nil {
		//Set the guessed number if there was any
		guess = g
	} else {
		return 0, 0, errors.New("No guess parameter")
	}

	//If they equal
	if guess == cookieValue {
		return 0, guess, nil
	}
	//If the guess is lower
	if guess < cookieValue {
		return -1, guess, nil
	}
	//If the guess is higher
	if guess > cookieValue {
		return 1, guess, nil
	}

	//Return error
	return 0, 0, errors.New("Nothing to compare")
}

//Function used to get the target cookie
func getTargetCookieValue(r *http.Request) (int, error) {
	//Get the cookie
	cookie, err := r.Cookie("target")
	//Check if cookie exists
	if err == nil && cookie != nil {
		//Try to parse the cookie's value
		if value, err := strconv.ParseInt(cookie.Value, 10, 0); err == nil {
			//Return the cookie
			return int(value), nil
		}
		//Retrun not exists error
		return 0, errors.New("Invalid value in cookie")
	}
	//Retrun not exists error
	return 0, errors.New("Not exists")
}

//Function used to check if the target cookies is set or not. If not set it will set it.
func checkAndSetTargetCookie(w http.ResponseWriter, r *http.Request) {
	//Check if the cookie exists. If not then set one
	if _, err := getTargetCookieValue(r); err != nil {
		setTargetCookie(w)
	}
}

//Function used to set the target cookie
func setTargetCookie(w http.ResponseWriter) {
	//Create the random number
	randomNumber := rand.Intn(RANDMAX-RANDMIN) + RANDMIN
	//Create the cookie
	cookie := http.Cookie{Name: "target", Value: strconv.Itoa(randomNumber), Expires: time.Now().Add(365 * 24 * time.Hour)}
	//Set the cookie
	http.SetCookie(w, &cookie)
}

//Function used to check for new game
func checkNewGameParam(r *http.Request) int {
	//Get the guess parameters. It might be more than one
	newgame := r.URL.Query().Get("newgame")
	//Check if it is empty
	if newgame != "" {
		//Return one if it exists
		return 1
	}
	return 0
}

//Function used to get the guessed number form the url parameters
func getGuessedNumberParameter(r *http.Request) (int, error) {
	//POST adapted from https://astaxie.gitbooks.io/build-web-application-with-golang/de/04.1.html

	//Parse the form
	r.ParseForm()
	//Check if the parameter exists
	if r.Form["guess"] != nil {
		//Try to parse it
		if guess, err := strconv.Atoi(strings.Join(r.Form["guess"], "")); err == nil {
			//Return the number and no error
			return int(guess), nil
		}
		//Return the invalid input error
		return 0, errors.New("Invalid input")
	}

	/*//Get the guess parameters. It might be more than one
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
	*/
	//Return the no parameter present error
	return 0, errors.New("No parameter present")
}
