package server

import (
	"context"
	"errors"

	"github.com/thebrubaker/colony/game"
	"github.com/thebrubaker/colony/pb"
	"github.com/thebrubaker/colony/streams"
)

type GameController interface {
	CreateGame() game.GameKey
	SendCommand(game.GameKey, string) bool
	SetSpeed(game.GameKey, game.TickRate) bool
}

type StreamController interface {
	CreateStream(game.GameKey, pb.GameService_StreamGameServer) *streams.Stream
	RemoveStream(game.GameKey, *streams.Stream) bool
}

type GameService struct {
	pb.UnimplementedGameServiceServer
	streamController StreamController
	gameController   GameController
}

func NewGameService(gc GameController, sc StreamController) *GameService {
	return &GameService{
		gameController:   gc,
		streamController: sc,
	}
}

func (gs *GameService) CreateGame(c context.Context, request *pb.CreateGameRequest) (*pb.GameState, error) {
	key := gs.gameController.CreateGame()

	return &pb.GameState{GameKey: string(key), Json: "{}"}, nil
}

func (gs *GameService) StreamGame(request *pb.StreamGameRequest, stream pb.GameService_StreamGameServer) error {
	key := game.GameKey(request.GameKey)
	s := gs.streamController.CreateStream(key, stream)
	defer gs.streamController.RemoveStream(key, s)
	<-stream.Context().Done()
	return stream.Context().Err()
}

func (gs *GameService) SendCommand(c context.Context, request *pb.SendCommandRequest) (*pb.SendCommandResponse, error) {
	key := game.GameKey(request.GameKey)
	gs.gameController.SendCommand(key, request.CommandType)
	return &pb.SendCommandResponse{CommandKey: ""}, nil
}

func (gs *GameService) CancelCommand(c context.Context, request *pb.CancelCommandRequest) (*pb.CancelCommandResponse, error) {
	return &pb.CancelCommandResponse{}, errors.New("Not Implemented")
}

func (gs *GameService) SetSpeed(c context.Context, request *pb.SetSpeedRequest) (*pb.SetSpeedResponse, error) {
	key := game.GameKey(request.GameKey)
	gs.gameController.SetSpeed(key, game.TickRate(request.Speed))
	return &pb.SetSpeedResponse{Err: ""}, nil
}
