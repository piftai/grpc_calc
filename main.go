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

func (s *server) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	result := req.A - req.B
	return &pb.SubtractResponse{Result: result}, nil
}

func (s *server) Multiple(ctx context.Context, req *pb.MultipleRequest) (*pb.MultipleResponse, error) {
	result := req.A * req.B
	return &pb.MultipleResponse{Result: result}, nil
}

func (s *server) Division(ctx context.Context, req *pb.DivisionRequest) (*pb.DivisionResponse, error) {
	result := req.A / req.B
	return &pb.DivisionResponse{Result: result}, nil
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
