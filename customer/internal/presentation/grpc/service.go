package grpc

import (
	"context"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/commands"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/presentation/grpc/pb"
	"github.com/google/uuid"
)

type CustomerGrpcService struct {
	cs *application.CustomerServices
}

func NewCustomerGrpcService(cs *application.CustomerServices) pb.CustomerServiceServer {
	return &CustomerGrpcService{
		cs: cs,
	}
}
func (s *CustomerGrpcService) CreateCustomer(
	_ context.Context, request *pb.CreateCustomerRequest,
) (*pb.CreateCustomerResponse, error) {
	customerID, customerErr := uuid.Parse(request.CustomerID)
	if customerErr != nil {
		return nil, customerErr
	}
	addressID, addressErr := uuid.Parse(request.AddressID)
	if addressErr != nil {
		return nil, addressErr
	}
	c := commands.CreateCustomerCommand{
		CustomerID: customerID,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		MiddleName: request.MiddleName,
		AddressID:  addressID,
		Email:      request.Email,
	}

	response, err := s.cs.Commands.CreateCustomer.Handle(c)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCustomerResponse{
		Id:      response.CustomerID,
		EventID: response.EventID,
	}, nil
}
