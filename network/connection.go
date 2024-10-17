package network

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"time"
)


const (
	castingAddr = "224.0.1.1:1234" // range: 224.0.0.1 -> 239.255.255.255
	MAXSIZE = 8192
)

var (
	PORT = 8080 
)


func InitAppServer() error {

	
	l, r := net.Listen( "tcp", fmt.Sprintf(":%v", PORT))

	if r != nil {

		PORT+=1
		// try another time with port 8081
		l, r = net.Listen( "tcp", fmt.Sprintf(":%v", PORT))

		if r != nil {
			return r
		}
		
	}

	for {

		conn, r := l.Accept()

		if r != nil {
			return r
		}

		b := make( []byte, MAXSIZE)

		n, r := conn.Read(b)

		if r != nil {
			return r
		}


		// TODO: need better error handling
		go func(n int, b []byte){

			// the data received from b[:n]
			message := string(b[:n])

			re, r := regexp.Compile(`tchat: received:(\d+)$`)

			if r != nil {
				return
			}

			if !re.MatchString(message) {
				return
			}

			// TODO:  we will establish a connection

		}(n, b)

				

	}

}



// this function will receive signals 
// and respond with a special message so that they know we can chat  
func ReceiveSignals(inter chan struct{}) error {

	addr, r := net.ResolveUDPAddr( "udp", castingAddr)

	if r != nil {
		return r
	}

	l, r := net.ListenMulticastUDP( "udp", nil, addr)

	l.SetReadBuffer(MAXSIZE)

	stop := false

	for !stop {

		// check if the user stopped the receiving of signals	
		select {
		case <- inter:
			stop = true
		default:
		}

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
	
		// using regex to check the message received 
		re, r := regexp.Compile(`tchat: sending:(\d+)$`)
		
		if r != nil {
			return r
		}

		// if the message is not correct we continue
		if !re.MatchString(string(b[:n])) {
			continue
		}

		// this will return in the index 1 the first capturing group which is the port number in our case
		match := re.FindStringSubmatch(string(b[:n]))


		// if the message is correct we send a response
		src.Port, r = strconv.Atoi(match[1])

		if r != nil {
			return r
		}

		conn, r := net.DialTCP( "tcp", nil, (*net.TCPAddr)(src))

		if r != nil {
			log.Fatal("error connecting with tcp ", r)
		}

		// TODO: Send a port in the response
		if _, r = conn.Write([]byte(fmt.Sprintf("tchat: received:%v", PORT))); r != nil {
			log.Fatal("error writing to tcp connection ", r)
		}

		conn.Close()

		break
	}

	return nil
}



func SendSignals() error {

	
	// the port for listenning
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


			buf := make([]byte, 1024)

			n, r := conn.Read(buf)

			if r != nil {
				log.Fatal("error reading data")
			}

			log.Println("received: ", string(buf[:n]))

			conn.Close()

			break
		}


		t.Close()

	}()



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


