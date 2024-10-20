package network

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"

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
			log.Println("Error: accepting the connection")
			continue
		}
		
		b := make( []byte, MAXSIZE)
		
		n, r := conn.Read(b)

		if r != nil {
			log.Println("Error: reading from the connection")
			continue 
		}


		// the data received from b[:n]
		message := string(b[:n])

		log.Println("we got ", message)

		// We take a format with the port and username of the sender
		re, r := regexp.Compile(`tchat:received:(\d+):(\S+)`)

		if r != nil {
			log.Println("Error: compiling regex expression")
			continue 
		}

		re2, r := regexp.Compile(`tchat:established:(\S+)`)
		
		if r != nil {
			log.Println("Error: compiling regex expression")
			continue 
		}

		// checkk if the message a connection establishing message 
		if re.MatchString(message) {

			log.Println("and it's accepted")

			parts := strings.Split(message, ":")

			port := parts[2]
			username := parts[3]


			// TODO: complete the handle connection  
			go handleConn( conn.RemoteAddr().String(), mainUser, port, username)
			
			conn.Close()
		
		} else if re2.MatchString(message) {
			
			parts := strings.Split(message, ":")

			go handleEstConn( conn, mainUser, parts[2])

		}
	}

}

// TODO: better error handling
func handleConn ( addr string, mainUser *user.User, port string, username string) {

	// changing the port to the one sent in the message 
	host, _, _ := net.SplitHostPort(addr)

	// the full address 
	addr = net.JoinHostPort(host, port)

	fAddr, r := net.ResolveTCPAddr( "tcp", addr)

	if r != nil {
		log.Fatalf("Error creating member's address from %s with error %s", addr, r.Error())
	}

	// create new member
	newMem := user.NewMember( fAddr, username)

	// adding the member 
	index := mainUser.AppendMembers( newMem)

	log.Printf("Added Member: %s, address: %v", newMem.GetUsername(), newMem.GetAddr().String())
	// buffer for reading 
	b := make( []byte, MAXSIZE)

	// TODO: create new connection to the port 

	conn, r := net.Dial( "tcp", fAddr.String())

	if r != nil {
		log.Fatalf("Error connecting with the member %s on %s", mainUser.Members[index].GetUsername(), mainUser.Members[index].GetAddr().String())
	}

	defer conn.Close()

	// writing to the sender so he can add us as a member 
	if _, r = conn.Write([]byte(fmt.Sprintf("tchat:established:%s", mainUser.Username))); r != nil {
		log.Fatalf("Error sending vital information to the member with %s", r.Error())
	}

	for {
		n, r := conn.Read(b)

		if r != nil {
			log.Fatalf("Error reading from the last connection with %s", r.Error())
		}

		// TODO: Receive messages and pass them using the index 
	
		log.Println(string(b[:n]))

	}

}


func handleEstConn( conn net.Conn, User *user.User, username string) {

	// add the member and get his index 
	newMem := user.NewMember( conn.RemoteAddr(), username)

	_ = User.AppendMembers( newMem)

	
	b := make( []byte, MAXSIZE)

	for {

		n, r := conn.Read(b)

		if r != nil {
			log.Fatalf("Error reading from established connection with %s", r.Error())
		}


		log.Println(string(b[:n]))

	}

}




