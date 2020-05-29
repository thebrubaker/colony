package server

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/thebrubaker/colony/game"
	"github.com/thebrubaker/colony/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type GameServer struct {
	pb.UnimplementedGameServerServer
	GameState *game.GameState
}

func (s *GameServer) StreamGameState(request *pb.EmptyRequest, stream pb.GameServer_StreamGameStateServer) error {
	for range time.Tick(time.Second) {
		json, _ := json.Marshal(s.GameState)
		stream.Send(&pb.GameState{
			Json: string(json),
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

func StartServer(gameState *game.GameState) {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := &GameServer{GameState: gameState}

	pb.RegisterGameServerServer(s, srv)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
