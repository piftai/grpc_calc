package main

import (
	"context"
	"log"
	"net"

	"github.com/piftai/grpc_calc/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.GetA() + req.GetB()
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Server did not wake:%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})

	log.Println("Server is running.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve:%v", err)
	}

}
