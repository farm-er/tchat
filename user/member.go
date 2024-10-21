package user

import "net"








type Member struct {

	addr net.Addr

	username string

	conn net.Conn

	con Conversation

}



func NewMember( addr net.Addr, username string, conn net.Conn) *Member {

	return &Member{
		addr: addr,
		username: username,
		conn: conn,
	}

}


func (m Member) GetUsername() string {
	return m.username
}

func (m Member) GetAddr() net.Addr {
	
	return m.addr

}

func (m *Member) SendText( mes string, mem string) error {

	_, r := m.conn.Write([]byte(mes))

	if r != nil {
		return r
	}

	m.con.Messages = append(m.con.Messages, &Message{
		Sender: mem,
		Receiver: m.GetUsername(),
		content: mes,
	})

	return nil
}

func (m *Member) GetLastMessages( n int) []*Message {
	
	l := len(m.con.Messages)

	if l <= n {
		return m.con.Messages
	}

	return m.con.Messages[l-n:]
}

