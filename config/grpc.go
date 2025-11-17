package config

import (
	"net"

	"google.golang.org/grpc"
)

func NewGrpc() (net.Listener, *grpc.Server) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	return lis, grpc.NewServer()
}
