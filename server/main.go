package main

import (
	"context"
	"log"
	"net"

	"../pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *pb.CalcRequest) (*pb.CalcReply, error) {
	result := in.A + in.B

	return &pb.CalcReply{Sum: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}

	grcpServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grcpServer, &server{})

	reflection.Register(grcpServer)

	log.Println("Server is listenting on port 50000.")
	if err := grcpServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve Calculator service: %v", err)
	}
}
