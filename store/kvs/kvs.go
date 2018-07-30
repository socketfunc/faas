package kvs

import (
	"context"

	"github.com/socketfunc/faas/store/hashes"
	pb "github.com/socketfunc/faas/store/proto"
)

type server struct {
	hashes *hashes.Hashes
}

var _ pb.KvsServer = (*server)(nil)

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	entity, err := s.hashes.Get(req.Namespace, req.Key)
	if err != nil {
		return nil, err
	}
	res := &pb.GetResponse{
		Namespace: req.Namespace,
		Key:       req.Key,
		Entity:    entity,
	}
	return res, nil
}

func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	if err := s.hashes.Put(req.Namespace, req.Key, req.Entity); err != nil {
		return nil, err
	}
	res := &pb.PutResponse{
		Successful: true,
	}
	return res, nil
}

func (s *server) Modify(ctx context.Context, req *pb.ModifyRequest) (*pb.ModifyResponse, error) {
	entity, successful, err := s.hashes.Modify(req.Namespace, req.Key, req.Filters, req.Updates)
	if err != nil {
		return nil, err
	}
	res := &pb.ModifyResponse{
		Namespace:  req.Namespace,
		Key:        req.Key,
		Successful: successful,
		Entity:     entity,
	}
	return res, nil
}

func (s *server) Del(ctx context.Context, req *pb.DelRequest) (*pb.DelResponse, error) {
	s.hashes.Del(req.Namespace, req.Key)
	res := &pb.DelResponse{}
	return res, nil
}

func (s *server) DelAll(ctx context.Context, req *pb.DelRequest) (*pb.DelResponse, error) {
	s.hashes.DelAll(req.Namespace)
	res := &pb.DelResponse{}
	return res, nil
}

func (s *server) Keys(ctx context.Context, req *pb.KeysRequest) (*pb.KeysResponse, error) {
	keys := s.hashes.Keys()
	res := &pb.KeysResponse{
		Keys: keys,
	}
	return res, nil
}

func New() *server {
	return &server{
		hashes: hashes.New(),
	}
}
