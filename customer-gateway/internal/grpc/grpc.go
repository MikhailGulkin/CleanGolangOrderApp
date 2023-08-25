package grpc

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer-gateway/internal/grpc/servicespb"
	"google.golang.org/grpc"
)

type Client struct {
	Client servicespb.CustomerServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := servicespb.NewCustomerServiceClient(conn)
	return &Client{c}, nil
}
