package grpc

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/presentation/grpc/servicespb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

type GrpcServer struct {
	service servicespb.CustomerServiceServer
}

func NewGrpcServer(service servicespb.CustomerServiceServer) *GrpcServer {
	return &GrpcServer{
		service: service,
	}
}

func (g *GrpcServer) Run() error {
	s := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: maxConnectionIdle * time.Minute,
		Timeout:           gRPCTimeout * time.Second,
		MaxConnectionAge:  maxConnectionAge * time.Minute,
		Time:              gRPCTime * time.Minute,
	}))
	servicespb.RegisterCustomerServiceServer(s, g.service)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50052))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
