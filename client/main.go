package main

import (
	"log"

	"../pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf(
			"Failed to create connection: %v",
			err,
		)
	}

	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	r, err := c.Add(context.Background(), &pb.CalcRequest{A: 1000, B: 2000})

	if err != nil {
		log.Fatalf(
			"Cannot execute function: %v",
			err,
		)
	}
	log.Printf(
		"Calculator: %d + %d = %d",
		1000, 2000, r.Sum,
	)
}
