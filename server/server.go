package server

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/thebrubaker/colony/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type GameServer struct {
	pb.UnimplementedGameServerServer
}

func (s *GameServer) StreamGameState(request *pb.EmptyRequest, stream pb.GameServer_StreamGameStateServer) error {
	for range time.Tick(16 * time.Millisecond) {
		stream.Send(&pb.GameState{
			Json: "",
		})
	}

	return nil
}

func (s *GameServer) GetCommandTypes(c context.Context, request *pb.EmptyRequest) (*pb.CommandList, error) {
	return &pb.CommandList{}, nil
}

func (s *GameServer) AddCommand(c context.Context, request *pb.AddCommandRequest) (*pb.Command, error) {
	return &pb.Command{}, nil
}

func (s *GameServer) RemoveCommand(c context.Context, request *pb.Command) (*pb.EmptyResponse, error) {
	return &pb.EmptyResponse{}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", port)

	log.Printf("Listening on %s", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	service := &GameServer{}

	log.Println("Registering service.")
	pb.RegisterGameServerServer(server, service)

	log.Println("Starting server.")
	server.Serve(lis)
}
