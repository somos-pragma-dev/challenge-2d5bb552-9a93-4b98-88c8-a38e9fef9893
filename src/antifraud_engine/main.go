package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "github.com/fintech/antifraud_engine/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAntifraudEngineServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {}

func (s *server) CheckTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	// Implementar lógica de chequeo antifraude
	return &pb.TransactionResponse{}, nil
}