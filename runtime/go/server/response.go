package server

import (
	"log"

	pb "github.com/socketfunc/faas/runtime/proto"
)

var _ Response = (*response)(nil)

type Response interface {
	Send(topic, event string, payload []byte)
	Reply(payload []byte)
	Broadcast(event string, payload []byte)
}

type response struct {
	topic  string
	event  string
	stream pb.Runtime_StreamServer
}

func (r *response) Send(topic, event string, payload []byte) {
	send := &pb.Send{
		Cmd: pb.Cmd_STREAM,
		StreamSend: &pb.StreamSend{
			Topic:   topic,
			Event:   event,
			Payload: payload,
		},
	}
	if err := r.stream.Send(send); err != nil {
		log.Printf("%+v\n", err)
	}
}

func (r *response) Reply(payload []byte) {
	send := &pb.Send{
		Cmd: pb.Cmd_STREAM,
		StreamSend: &pb.StreamSend{
			Topic:   r.topic,
			Event:   r.event,
			Payload: payload,
		},
	}
	if err := r.stream.Send(send); err != nil {
		log.Printf("%+v\n", err)
	}
}

func (r *response) Broadcast(event string, payload []byte) {

}

func newResponse(topic, event string, stream pb.Runtime_StreamServer) *response {
	return &response{
		topic:  topic,
		event:  event,
		stream: stream,
	}
}
