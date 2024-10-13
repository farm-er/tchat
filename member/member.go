package member

import (
	"fmt"
	"math/rand"
	"net"
)




type member struct {

	addr *net.UDPAddr 

	// the member's name and it's initialized with a random name 
	name string 

}



func NewMember( addr *net.UDPAddr) *member {

	return &member{
		addr: addr,
		name: fmt.Sprintf("member%v", rand.Intn(999999)),
	}
	
}




