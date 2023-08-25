//go:generate go run github.com/99designs/gqlgen generate
package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer-gateway/internal/grpc"
)

type Server struct {
	customerClient *grpc.Client
}

func NewGraphQLServer(customerUrl string) (*Server, error) {
	customerClient, err := grpc.NewClient(customerUrl)
	if err != nil {
		return nil, err
	}
	return &Server{
		customerClient: customerClient,
	}, nil
}
func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		client: s.customerClient.Client,
	}
}
func (s *Server) Query() QueryResolver {
	return &queryResolver{}
}
func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
