package main

import (
	pb "grpc_dummy/Generate/Protoc"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "127.0.0.1:3001"

func main() {

	dummyClientConn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Connot connect to grpc server on %s. Error: %v.\n", addr, err)
	}

	dummyClient := pb.NewDummyServiceClient(dummyClientConn)

	tryDummyService(dummyClient)
}
