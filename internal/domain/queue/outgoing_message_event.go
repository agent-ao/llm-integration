package queue

import (
	"google.golang.org/genai"
)

type OutgoingMessageEvent struct {
	AuthToken         string           `json:"auth_token"`
	ClientContactId   string           `json:"client_contact_id"`
	BusinessContactId string           `json:"business_contact_id"`
	Content           []*genai.Content `json:"content"`
}
