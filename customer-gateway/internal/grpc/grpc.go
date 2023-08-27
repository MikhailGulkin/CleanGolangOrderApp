package grpc

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer-gateway/internal/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(url string) (pb.CustomerServiceClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := pb.NewCustomerServiceClient(conn)
	return c, nil
}
