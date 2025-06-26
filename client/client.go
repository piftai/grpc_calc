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
	n := time.Now()

	res, err := c.Add(ctx, &pb.AddRequest{A: 1, B: 4})
	if err != nil {
		log.Fatalf("Could not add: %v", err)
	}

	log.Printf("Addition Result: %d", res.GetResult())

	res1, err := c.Subtract(ctx, &pb.SubtractRequest{A: 500, B: 150})
	if err != nil {
		log.Fatalf("Could not subtract:%v", err)
	}

	log.Printf("Subtraction result: %v", res1.Result)

	res2, err := c.Multiple(ctx, &pb.MultipleRequest{A: 2, B: 2})
	if err != nil {
		log.Fatalf("Could not subtract:%v", err)
	}

	log.Printf("Multiplication result: %v", res2.Result)

	res3, err := c.Division(ctx, &pb.DivisionRequest{A: 851381583, B: 2651.2})
	if err != nil {
		log.Fatalf("Could not subtract:%v", err)
	}

	log.Printf("Divison result: %v", res3.Result)
	log.Printf("")

	log.Printf("wasted time of division request: %v", time.Since(n))

}
