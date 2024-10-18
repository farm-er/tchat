package user








type User struct {

	// Visible name to other users
	Username string

	// the member we're trying to send the message to 
	MemFocus string

	// the text we want to send 
	Message string

	// all open conversation 
	Members []*User

}





func NewUser( name string) *User {

	return &User{
		Username: name,
	}

} 








