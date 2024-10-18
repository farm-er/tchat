package network

import (
	"fmt"
	"net"
	"regexp"
)


const (
	castingAddr = "224.0.1.1:1234" // range: 224.0.0.1 -> 239.255.255.255
	MAXSIZE = 8192
)

func InitAppServer(port int) error {

	
	l, r := net.Listen( "tcp", fmt.Sprintf(":%v", port))

	if r != nil {
		return r
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





