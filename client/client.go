package client

import (
	"context"
	"log"

	"github.com/thebrubaker/colony/pb"
	"google.golang.org/grpc"
)

func CreateClient(address string) (pb.GameServiceClient, *grpc.ClientConn, context.Context, context.CancelFunc) {
	log.Printf("Connecting to %s", address)

	// Set up a connection to the server.
	connection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewGameServiceClient(connection)

	context, cancel := context.WithCancel(context.Background())

	return client, connection, context, cancel
}
