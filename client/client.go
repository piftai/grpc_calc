package main

import (
	"context"
	"log"
	"time"

	pb "github.com/piftai/grpc_calc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Add(ctx, &pb.AddRequest{A: 1, B: 4})
	if err != nil {
		log.Fatalf("Could not add: %v", err)
	}

	log.Printf("Addition Result: %d", res.GetResult())

	res1, err := c.Subtract(ctx, &pb.SubtractRequest{A: 5, B: 10})
	if err != nil {
		log.Fatalf("Could not subtract:%v", err)
	}

	log.Printf("Subtraction result: %v", res1.Result)
}
