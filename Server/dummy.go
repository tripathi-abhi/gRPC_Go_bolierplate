package main

import (
	"context"
	pb "grpc_dummy/Generate/Protoc"
)

func (s *Server) Dummy(context.Context, *pb.DummyRequest) (*pb.DummyResponse, error) {
	return &pb.DummyResponse{
		DummyString: "You'll get only this for dummy!!",
	}, nil
}
