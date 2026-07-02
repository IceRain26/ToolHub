package chat

type Session struct {
	Messages []Message
}

func New() *Session{
	return &Session{
		Messages: []Message{},
	}
}

func (s *Session) Add(role, content string){
	s.Messages = append(s.Messages, Message{
		Role: role,
		Content: content,
	})
}