package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	expanse "my-budgety-gRPC-server/expanse"
	expanse_pb "my-budgety-gRPC-server/expanse/protos"
)

const (
	port = ":50051"
)

type server struct {
	expanse_pb.ExpanseServiceServer
}

func (s *server) GetExpansesAmountInCurrentMonth(ctx context.Context,
	in *expanse_pb.GetExpansesAmountInCurrentMonthRequest) (*expanse_pb.GetExpansesAmountInCurrentMonthResponse, error) {
	expanse_amount := expanse.GetAllExpanse()
	return &expanse_pb.GetExpansesAmountInCurrentMonthResponse{Amount: expanse_amount}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	expanse_pb.RegisterExpanseServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Hello, World!")
}
