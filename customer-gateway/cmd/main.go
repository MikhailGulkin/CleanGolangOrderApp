package main

import (
	"fmt"
	"github.com/99designs/gqlgen/handler"
	"github.com/MikhailGulkin/CleanGolangOrderApp/pkg/env"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer-gateway/internal/graphql"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer-gateway/internal/grpc"
	"log"
	"net/http"
)

func main() {
	grpcClient, err := grpc.NewClient(env.GetEnv("GRPC_CLIENT", "localhost:50052"))
	if err != nil {
		log.Fatal(err)
	}
	s, err := graphql.NewGraphQLServer(grpcClient)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.GetEnv("GRAPH_PORT", "8001")), nil))
}
