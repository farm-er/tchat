package user





type Message struct {

	Sender string

	Receiver string

	content string

}


func (m *Message) GetContent() string {
	return m.content
}


type Conversation struct {

	Messages []*Message 

}




