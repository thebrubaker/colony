PROTO_PATH=/Users/joel-brubaker/go/src/github.com/thebrubaker/colony/pb
GO_OUT=/Users/joel-brubaker/go/src

protoc --go_out=plugins=grpc:. game.proto