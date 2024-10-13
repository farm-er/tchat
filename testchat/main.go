
package main

import (
	"log"
	"net"
	"time"
)

const (
	sendingPort   = 12345
	receivingPort = 12345 // The same port as sending
	serverAddr    = "224.0.0.1:9999"
)

func main() {
	// Create a UDP connection for sending and receiving
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"), // Listen on all available interfaces
		Port: sendingPort,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Set the buffer size
	conn.SetReadBuffer(8192)

	// Send a message
	sendAddr := &net.UDPAddr{
		IP:   net.ParseIP("224.0.0.1"),
		Port: 9999, // Send to multicast address
	}

	_, err = conn.WriteToUDP([]byte("Hello, multicast!"), sendAddr)
	if err != nil {
		log.Fatal("Error sending message: ", err)
	}

	log.Println("Message sent, waiting for a response...")

	// Listen for response on the same port
	buffer := make([]byte, 8192)
	conn.SetDeadline(time.Now().Add(5 * time.Second)) // Optional: Set a timeout for response

	// Reading from the same connection
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatal("Error receiving response: ", err)
	}

	log.Printf("Received response from %v: %s", addr, string(buffer[:n]))
}
