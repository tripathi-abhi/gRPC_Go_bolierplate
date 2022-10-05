package main

import (
	pb "grpc_dummy/Generate/Protoc"
	"log"
	"net"

	"google.golang.org/grpc"
)

const addr = "127.0.0.1:3001"

type Server struct {
	pb.DummyServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %s. Error: %v.\n", addr, err)
	}

	s := grpc.NewServer()

	pb.RegisterDummyServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve. Error: %v\n.", err)
	}
}
