package user

import "net"








type Member struct {

	addr net.Addr

	username string

}



func NewMember( addr net.Addr, username string) *Member {

	return &Member{
		addr: addr,
		username: username,
	}

}


func (m Member) GetUsername() string {
	return m.username
}

func (m Member) GetAddr() net.Addr {
	
	return m.addr

}

