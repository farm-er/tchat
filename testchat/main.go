package main

import (
	"log"
	"net"
	"time"
)

const (
	castingAddr = "224.0.1.1:1234" // range: 224.0.0.1 -> 239.255.255.255
	MAXSIZE = 8192
)


func main() {

	// listen locally for tcp
	go func() {

		t, r := net.Listen( "tcp", ":8080")

		if r != nil {
			log.Fatal("error listenning to TCP on port 8080: ", r)
		}

		log.Println("Listenning to port :8080")

		for {

			conn, r := t.Accept()

			if r != nil {
				log.Fatal("error accepting the connection on port 8080: ", r)
			}


			sender := conn.RemoteAddr().String()

			buf := make([]byte, 1024)

			n, r := conn.Read(buf)

			if r != nil {
				log.Fatal("error reading data")
			}

			log.Println("received: ", string(buf[:n]), " from ", sender)

			conn.Close()

			break
		}


		t.Close()

	}()

	


	// multicasting to castingAddr
	addr, r := net.ResolveUDPAddr( "udp", castingAddr)

	if r != nil {
		log.Fatal("Error resolving Casting address ", r)
	}

	// making a connection
	s, r := net.DialUDP("udp", nil, addr)

	if r != nil {
		log.Fatal("Error setting up UDP connection ", r)
	}

	for {

		s.Write([]byte("tchat: sending:8080"))

		time.Sleep(time.Second)
	}

}
