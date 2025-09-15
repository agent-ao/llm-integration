package queue

import "google.golang.org/genai"

type IncomingMessageEvent struct {
	AuthToken           string           `json:"auth_token"`
	ClientContactId     string           `json:"client_contact_id"`
	BusinessContactId   string           `json:"business_contact_id"`
	ConversationHistory []*genai.Content `json:"conversation_history"`
	SystemInstruction   []*genai.Part    `json:"system_instruction"`
}
