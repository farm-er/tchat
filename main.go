package main

import (
	"encoding/hex"
	"log"
	"net"
	"time"

	"github.com/farm-er/tchat/frontend"
)


const (
	castingAddr = "224.0.1.1:1234" // range: 224.0.0.1 -> 239.255.255.255
	MAXSIZE = 8192
)



func sendSignals() error {

	addr, r := net.ResolveUDPAddr( "udp", castingAddr)

	if r != nil {
		return r
	}

	// making a connection
	s, r := net.DialUDP("udp", nil, addr)

	if r != nil {
		return r
	}

	for {

		s.Write([]byte("tchat: sending"))

		time.Sleep(time.Second)
	}

}



// this function will receive signals 
// and respond with a special message so that they know we can chat  
func receiveSignals() error {

	addr, r := net.ResolveUDPAddr( "udp", castingAddr)

	if r != nil {
		return r
	}

	l, r := net.ListenMulticastUDP( "udp", nil, addr)

	l.SetReadBuffer(MAXSIZE)

	for {

		b := make( []byte, MAXSIZE)

		// receive message
		// n is the number of bytes read
		// b the buffer where the bytes read are stored
		// src is the *net.UDPAddr of the sender
		n, src, r := l.ReadFromUDP(b)
		
		if r != nil {
			return r
		}

		// TODO: check the message if it's a tchat sending or received signal 

		// if string(b[:n]) != "tchat: sending" {
		// 	continue
		// }

		log.Println(hex.Dump(b[:n]))

		conn, r :=  net.DialUDP( "udp", nil, src)

		if r != nil {
			log.Println("error")
			return r
		}

		conn.Write([]byte("tchat: received"))

		conn.Close()

	}

}



func main() {

	// when starting the app will listen for any signal 
	go func ()  {
		
		if r := receiveSignals(); r != nil {
			log.Fatal("Error occured while listening to other members")
		}
	
	}()

	
	// manipulating the screen
	s, r := frontend.NewScreen()

	if r != nil {
		log.Fatal("Error occured while initializing the screen")
	}

	s.Start()

}




