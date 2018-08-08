package manager

import (
	pb "github.com/socketfunc/faas/runtime/proto"
)

type server struct{}

var _ pb.RuntimeServer = (*server)(nil)

func (s *server) Stream(stream pb.Runtime_StreamServer) error {
	return nil
}

func New() *server {
	return &server{}
}
