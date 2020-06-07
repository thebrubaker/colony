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
	Games []*game.Game
}

func (s *GameServer) CreateGame(c context.Context, request *pb.CreateGameRequest) (*pb.CreateGameResponse, error) {
	name := request.Name

	log.Printf("new request: %v", request)

	// for _, game := range s.Games {
	// 	if game.Name == name {
	// 		return &pb.CreateGameResponse{Name: name, Error: "Game name already active."}, nil
	// 	}
	// }

	// s.Games = append(s.Games, game.CreateGame(name))

	return &pb.CreateGameResponse{Name: name}, nil
}

func (s *GameServer) StreamGame(request *pb.StreamGameRequest, stream pb.GameServer_StreamGameServer) error {
	var game *game.Game

	name := request.Name

	for _, g := range s.Games {
		if g.Name == name {
			game = g
		}
	}

	for range time.Tick(16 * time.Millisecond) {
		data, err := json.Marshal(game)

		if err != nil {
			continue
		}

		stream.Send(&pb.Game{
			Json: string(data),
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
	var games []*game.Game

	lis := CreateListener()
	server := CreateServer(games)

	server.Serve(lis)
}

func CreateServer(games []*game.Game) *grpc.Server {
	server := grpc.NewServer()
	service := &GameServer{
		Games: games,
	}

	log.Println("Registering service.")
	pb.RegisterGameServerServer(server, service)

	log.Println("Starting server.")

	return server
}

func CreateListener() net.Listener {
	lis, err := net.Listen("tcp", port)

	log.Printf("Listening on %s", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return lis
}
