package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "github.com/fintech/antifraud_engine/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err!= nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAntifraudEngineClient(conn)

	req := &pb.TransactionRequest{}
	res, err := c.CheckTransaction(context.Background(), req)
	if err!= nil {
		log.Fatalf("could not check transaction: %v", err)
	}
	log.Printf("Transaction response: %v", res)
}