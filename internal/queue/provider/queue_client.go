package provider

type QueueClient interface {
	PublishMessage(queue string, message []byte) error
	ConsumeMessages(queue string, handler func([]byte) error) error
	Close()
}
