package kvs

import (
	"io"

	"log"

	"github.com/socketfunc/faas/store/hashes"
	pb "github.com/socketfunc/faas/store/proto"
)

type server struct {
	hashes *hashes.Hashes
}

var _ pb.KvsServer = (*server)(nil)

func (s *server) Command(stream pb.Kvs_CommandServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		switch in.Command {
		case pb.Command_Get:
			go s.get(in.Get, stream)
		case pb.Command_Put:
			go s.put(in.Put, stream)
		case pb.Command_Modify:
			go s.modify(in.Modify, stream)
		case pb.Command_Del:
			go s.del(in.Del, stream)
		case pb.Command_DellAll:
			go s.dellAll(in.Del, stream)
		case pb.Command_Keys:
			go s.keys(in.Keys, stream)
		default:
			return nil
		}
	}
}

func (s *server) get(req *pb.GetRequest, stream pb.Kvs_CommandServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	entity, err := s.hashes.Get(req.Namespace, req.Key)
	if err != nil {
		return err
	}

	resp := &pb.CommandResponse{
		Command: pb.Command_Get,
		Get: &pb.GetResponse{
			Namespace: req.Namespace,
			Key:       req.Key,
			Entity:    entity,
		},
	}
	return stream.Send(resp)
}

func (s *server) put(req *pb.PutRequest, stream pb.Kvs_CommandServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	if err := s.hashes.Put(req.Namespace, req.Key, req.Entity); err != nil {
		return err
	}

	resp := &pb.CommandResponse{
		Command: pb.Command_Put,
		Put: &pb.PutResponse{
			Successful: true,
		},
	}
	return stream.Send(resp)
}

func (s *server) modify(req *pb.ModifyRequest, stream pb.Kvs_CommandServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	entity, successful, err := s.hashes.Modify(req.Namespace, req.Key, req.Filters, req.Updates)
	if err != nil {
		return err
	}

	resp := &pb.CommandResponse{
		Command: pb.Command_Modify,
		Modify: &pb.ModifyResponse{
			Namespace:  req.Namespace,
			Key:        req.Key,
			Successful: successful,
			Entity:     entity,
		},
	}
	return stream.Send(resp)
}

func (s *server) del(req *pb.DelRequest, stream pb.Kvs_CommandServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	s.hashes.Del(req.Namespace, req.Key)

	resp := &pb.CommandResponse{
		Command: pb.Command_Del,
		Del:     &pb.DelResponse{},
	}
	return stream.Send(resp)
}

func (s *server) dellAll(req *pb.DelRequest, stream pb.Kvs_CommandServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	s.hashes.DelAll(req.Namespace)

	resp := &pb.CommandResponse{
		Command: pb.Command_DellAll,
		Del:     &pb.DelResponse{},
	}
	return stream.Send(resp)
}

func (s *server) keys(req *pb.KeysRequest, stream pb.Kvs_CommandServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	keys := s.hashes.Keys()

	resp := &pb.CommandResponse{
		Command: pb.Command_Keys,
		Keys: &pb.KeysResponse{
			Keys: keys,
		},
	}
	return stream.Send(resp)
}

func New() *server {
	return &server{
		hashes: hashes.New(),
	}
}
