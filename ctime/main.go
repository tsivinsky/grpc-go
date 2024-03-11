package main

import (
	"context"
	"data"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	data.UnimplementedTimeServer
}

func (s *server) Now(ctx context.Context, in *data.NowRequest) (*data.NowResponse, error) {
	t := time.Now().Unix()
	return &data.NowResponse{Message: &t}, nil
}

func main() {
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	data.RegisterTimeServer(s, &server{})

	if err := s.Serve(ln); err != nil {
		panic(err)
	}
}
