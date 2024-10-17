package main

import (
	"flag"
	"log"

	"github.com/farm-er/tchat/frontend"
	"github.com/farm-er/tchat/network"
	users "github.com/farm-er/tchat/user"
)

// TODO: New plan
// --- when starting the app will listen on port 8080
// --- But joinning the multicast will be by the user
// --- Accepting any request will be by the user
// --- Sending any signals to find other members will be by the user


func main() {

	
	userName := flag.String( "username", "", "your username that will be shown to other users")

	// initialize the user
	user := users.NewUser(*userName)	

	log.Println(user)

	go network.InitAppServer()

	// initialize the screen
	s, r := frontend.NewScreen()

	if r != nil {
		log.Fatal("Error occured while initializing the screen")
	}

	s.Start()

}




