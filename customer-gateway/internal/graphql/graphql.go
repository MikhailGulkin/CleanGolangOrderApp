//go:generate go run github.com/99designs/gqlgen generate
package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer-gateway/internal/grpc/pb"
)

type Server struct {
	customerClient pb.CustomerServiceClient
}

func NewGraphQLServer(client pb.CustomerServiceClient) (*Server, error) {
	return &Server{
		customerClient: client,
	}, nil
}
func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		client: s.customerClient,
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
