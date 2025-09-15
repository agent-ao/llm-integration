package queue

import "google.golang.org/genai"

type Prompt struct {
	History     []*genai.Content
	Instruction *genai.Content
}
