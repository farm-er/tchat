package network

import (
	"fmt"
	"log"
	"net"
	"regexp"

	"github.com/farm-er/tchat/user"
)


const (
	castingAddr = "224.0.1.1:1234" // range: 224.0.0.1 -> 239.255.255.255
	MAXSIZE = 8192
)

func InitAppServer(port int, mainUser *user.User) error {

	
	l, r := net.Listen( "tcp", fmt.Sprintf(":%v", port))

	if r != nil {
		return r
	}

	defer l.Close()

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

		addr := conn.RemoteAddr()

		// TODO: need better error handling
		go func(n int, b []byte){

			// the data received from b[:n]
			message := string(b[:n])

			re, r := regexp.Compile(`tchat: received:(\d+)$`)

			if r != nil {
				return
			}
			
			// checkk if the message a connection establishing message 
			if re.MatchString(message) {
			
				// TODO: add username in the message

				newMember := user.NewMember( addr, "username until we make one")

				mainUser.Members = append(mainUser.Members, newMember)

				log.Println("added new member")

				return 
			}

		}(n, b)

	}

}





