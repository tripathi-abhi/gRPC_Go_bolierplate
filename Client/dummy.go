package main

import (
	"context"
	"fmt"
	pb "grpc_dummy/Generate/Protoc"
	"log"
)

func tryDummyService(dummyClient pb.DummyServiceClient) {
	res, err := dummyClient.Dummy(context.Background(), &pb.DummyRequest{})

	if err != nil {
		log.Fatalf("Cannot get response. Error: %v.\n", err)
	}

	fmt.Printf("Dummy Response: %v.\n", res)
}
