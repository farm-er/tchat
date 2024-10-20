package network

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"

	"github.com/farm-er/tchat/user"
)

// this function will receive signals
// and respond with a special message so that they know we can chat
func ReceiveSignals(inter chan struct{}, mainUser *user.User) error {

	log.Println("starting receiving")
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
			return r
		}

		log.Println(string(b[:n]))

		// TODO: Send a port in the response
		if _, r = conn.Write([]byte(fmt.Sprintf("tchat:received:%v:%s", mainUser.Port, mainUser.Username))); r != nil {
			return r
		}

		conn.Close()

		break
	}

	log.Println("stopped receiving")
	return nil
}









