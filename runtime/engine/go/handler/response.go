package handler

import (
	pb "github.com/socketfunc/faas/runtime/proto"
)

type Response struct {
	conn pb.Internal_PipeServer
}

func (r *Response) Send(topic, event string, id int, payload []byte) error {
	pipeRes := &pb.PipeResponse{
		Version: 1,
		Cmd:     pb.Cmd_Send,
		Packet: &pb.Packet{
			Topic:   topic,
			Event:   event,
			Id:      int32(id),
			Payload: payload,
		},
	}
	if err := r.conn.Send(pipeRes); err != nil {
		return err
	}
	return nil
}

func NewResponse(conn pb.Internal_PipeServer) *Response {
	return &Response{
		conn: conn,
	}
}
