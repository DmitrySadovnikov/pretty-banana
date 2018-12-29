package app

import (
	"context"
	"net"
	"os"
	"log"
	p "direction_service/app/proto"
	s "direction_service/app/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type gRPCServer struct{}

func (_ *gRPCServer) Calculate(c context.Context, r *p.Calculate_Request) (*p.Calculate_Response, error) {
	service := s.DirectionsCalculateService{}
	result := service.Call(r)
	return result, nil
}

func StartGRPCServer() {
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("gRPC server running on port " + port + "...")

	server := grpc.NewServer()
	p.RegisterDirectionServer(server, &gRPCServer{})

	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
