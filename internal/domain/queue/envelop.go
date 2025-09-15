package queue

type Envelop struct {
	Type       string `json:"type"`
	RoutingKey string `json:"-"`
	Data       any    `json:"data"`
}
