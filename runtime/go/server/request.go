package server

var _ Request = (*request)(nil)

type Request interface {
	Topic() string
	Event() string
	Payload() []byte
}

type request struct {
	topic   string
	event   string
	payload []byte
}

func (r *request) Topic() string {
	return r.topic
}

func (r *request) Event() string {
	return r.event
}

func (r *request) Payload() []byte {
	return r.payload
}

func newRequest(topic, event string, payload []byte) *request {
	return &request{
		topic:   topic,
		event:   event,
		payload: payload,
	}
}
