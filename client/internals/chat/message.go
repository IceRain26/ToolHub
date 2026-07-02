package chat

const (
	RoleSystem = "system"
	RoleUser = "user"
	RoleAssistant = "assistant"
)

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}