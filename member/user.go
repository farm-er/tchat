package member









type User struct {

	// the member we're trying to send the message to 
	MemFocus string

	// the text we want to send 
	Message string

	// all open conversation 
	Members []*member

}





