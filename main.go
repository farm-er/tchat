package main

import (
	"flag"
	"log"
	"os"

	"github.com/farm-er/tchat/frontend"
	"github.com/farm-er/tchat/network"
	users "github.com/farm-er/tchat/user"
)

// TODO: New plan
// --- when starting the app will listen on port 8080 by default or port chosen by the user
// --- But joinning the multicast will be by the user
// --- Accepting any request will be by the user
// --- Sending any signals to find other members will be by the user


func main() {


	out, r := os.OpenFile( "log_output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	
	if r != nil {
		log.Fatal("Error opening log file ", r)
	}

	log.SetOutput(out)

	userName := flag.String( "username", "", "your username that will be shown to other users")
	port := flag.Int( "port", 8080, "the port to use for communication")

	flag.Parse()
	
	// initialize the user
	user := users.NewUser(*userName)	

	// initializing the server 
	go func ( port int) {

		if r := network.InitAppServer( port); r != nil {
			log.Fatal("Error initializing the server", r)
		}

	}(*port)

	// initialize the screen
	s, r := frontend.NewScreen(user, *port)

	if r != nil {
		log.Fatal("Error occured while initializing the screen")
	}

	s.Start()

}




