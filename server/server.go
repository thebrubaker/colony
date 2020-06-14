package server

import (
	"log"
	"net"

	"github.com/thebrubaker/colony/pb"
	"google.golang.org/grpc"
)

func NewServer(lis net.Listener, service *GameService) *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterGameServiceServer(server, service)
	go func() {
		err := server.Serve(lis)

		if err != nil {
			log.Println(err)
		}

		return
	}()
	return server
}
