package user


// TODO: maybe add a mutex for concurrently accessed data
type User struct {

	// Visible name to other users
	Username string

	// the member we're trying to send the message to 
	MemFocus int

	// the text we want to send 
	Message string

	// all open conversation 
	Members []*Member

	// the port used 
	Port int

}





func NewUser( name string, port int) *User {

	return &User{
		Username: name,
		Port: port,
		MemFocus: 0,
		Message: "",
		Members: []*Member{},
	}

} 


func (u *User) ShiftFocusN() {
	if u.MemFocus == len(u.Members) - 1 {
		u.MemFocus = 0
		return
	}
	u.MemFocus += 1
}

func (u *User) ShiftFocusP() {

	if u.MemFocus == 0 {
		u.MemFocus = len(u.Members) - 1
		return
	}
	u.MemFocus -= 1
	return
}

// TODO: mutex for members
func (u *User) AppendMembers(mem *Member) int {

	u.Members = append(u.Members, mem)

	return len(u.Members)-1
}


func (u *User) ReceiveText( mes string, index int) {

	u.Members[index].con.Messages = append(u.Members[index].con.Messages, &Message{
		Sender: u.Members[index].GetUsername(),
		Receiver: u.Username,
		content: mes,
	})

}



