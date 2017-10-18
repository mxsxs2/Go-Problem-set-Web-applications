# Go Problem set - Web applications
The application is a simple guessing game on a webserver
## Description
* The application serves at a given port which is 8080 at default and can be loaded at http://localhost:8080 or http://127.0.0.1:8080
* At first load it generates a random number between given boundaries which are 10 and 20 by default
* If the user clicks on the new game button on the landing page he/she will be redirected to the guess page where he/she can start guessing the number.

* Corresponding to the users guess the application will return with an appropriate answer which can be:
    * _Your guess was ```<guess>``` which is too low._
    * _Your guess was ```<guess>``` which is too high._
    * _Congratulations! You guessed the number._

* Once the user guessed the correct number, the app will offer a new game. 
* When a new game starts, a new random number is generated

## Extra features
### Locally load bootstrap
The bootstrap files from http://getbootstrap.com/ has been moved to ```css``` and ```js``` folders for faster loading. This way we do not have to worry about the availability of these files from outside the server.
### Serving favico.ico
I downloaded an ico file from https://m.veryicon.com/icons/application/ios7-style-metro-ui/metroui-folder-os-game-center.html and implemented the serving handler so the server does not return with 404
### Styling
* Used bootsrap's built in feauters to make the ```a``` tags (links) look like a button
* Set styling onto the input field and the guess button
* Set a padding on the whole body of the page
### Auto focus on input field. Form submission on enter
I wrote a small amout of JavaScript for better user experience:
* When the guess page is loaded the cursor will be set to the input
* The form can be sent with the press of the enter key
### Support for configuration
* The port and the random number generator can be configured by command line arguments
### Security
* The guessing field is validated at server side, therefore if the user inserted other than a whole number into the field, the app will ignore it.
* Every command line argument has their own default value set up. If a command line argument is not a whole number or missing, the default value will be used.
* If the given maximum value for the random number generator is less than the minimum value, the app will change the maximum value to the minimum value increased by 2, so the app will not crash.

## How to install and run GO

To install, simply go to GO's website and download the installer and run it: https://golang.org/

When the installation is done, the "go" command is going to be avaialable in terminal.(You have to restart an opened terminal)


## How to build the app

To build the application, clone this repository and in the folder run: 
```
go build
```

The previous command will compile the go file into a runnable.

Once the runnable is created then it can be run in terminal or as a standalone application, for example on windows: 
```
./Go-Problem-set-Web-applications.exe
```

## Setting the configuration flags
* Port number can be set in the command line with the ```-port``` flag followed by a whole number
* Minimum value for the random number generator can be set in the command line with the ```-min``` flag followed by a whole number
* Maximum value for the random number generator can be set in the command line with the ```-max``` flag followed by a whole number
For example:
```
./Go-Problem-set-Web-applications.exe -port 8081 -min 1 -max 100
```
## How to use curl for examining the response from this server
### Step 1: Installation and start
#### Windows 64
Download curl from here: https://data-representation.github.io/resources/curl.zip
Navigate to the downloads, unzip the file.
Where the file is unzipped start a command window and type in: ```curl.exe -v http://127.0.0.1:8080```

#### Linux/Mac os
Curl is available on mac and linux. Do not have to install manually.
Open up a treminal and type in ```curl -v http://127.0.0.1:8080```
### Step 2: Examining the response from the server
The above mentioned command with the parameters should return something like:
```
* Rebuilt URL to: http://127.0.0.1:8080/
* Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:8081
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Content-Length: 1173
< Content-Type: text/html
< Last-Modified: Wed, 18 Oct 2017 21:00:04 GMT
< Date: Wed, 18 Oct 2017 22:42:48 GMT
<
{ [1173 bytes data]
100  1173  100  1173    0     0   1173      0  0:00:01 --:--:--  0:00:01 1145k<!DOCTYPE html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>Guessing Game</title>
    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">
  </head>
  <body class="container-fluid">
    <h1>Guessing game</h1>
    <a href="/guess" class="btn btn-primary"  role="button">New Game</a>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="/js/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="/js/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
    <script src="/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
  </body>
</html>


* Connection #0 to host 127.0.0.1 left intact

```

* The first line _* Rebuilt URL to: http://127.0.0.1:8080/_ tells us, curl added a trailing _/_ to the url as we did not provide one.
* The second line _Trying 127.0.0.1...._ tells us, curl is trying to connect to this ip address, which is a special ip, the localhost.
* The third line _* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)_ tells us, curl is successfully connected to the given address.
* The next part is the request headers, which were sent by curl to the server:
```
> GET / HTTP/1.1          <--- Which method the curl uses on which protocol. Send a GET request with HTTP/1.1
> Host: 127.0.0.1:8081    <--- The server which host to look for as there could be more than one on the same server
> User-Agent: curl/7.55.1 <--- Who/what is looking for this server
> Accept: */*             <--- What can be accepted by curl. In this case it can be anything
```
* The next part is the response headers, which were sent back by the server to curl:
```
< HTTP/1.1 200 OK                              <--- Responds back on the same protocol with a status 200, so everything went ok
< Accept-Ranges: bytes                         <--- "Used by the server to advertise its support of partial requests."
< Content-Length: 1173                         <--- Send back the lenght of the page in bytes
< Content-Type: text/html                      <--- It is a HTML file
< Last-Modified: Wed, 18 Oct 2017 21:00:04 GMT <--- When the file was modified last time
< Date: Wed, 18 Oct 2017 22:42:48 GMT          <--- The current time
```
* The last part is the content of the file/page itself
* The very last line tells us the connection was closed between curl and the server

## References
* https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Ranges
* https://stackoverflow.com/questions/43601359/how-do-i-serve-css-and-js-in-go-lang
* https://m.veryicon.com/icons/application/ios7-style-metro-ui/metroui-folder-os-game-center.html
* http://www.thegeekstuff.com/2012/04/curl-examples/
* https://getbootstrap.com/docs/4.0/getting-started/introduction/
* https://golang.org/pkg/net/http/