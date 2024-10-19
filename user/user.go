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

}





func NewUser( name string) *User {

	return &User{
		Username: name,
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




