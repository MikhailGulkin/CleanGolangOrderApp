package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer-gateway/internal/graphql"
	"log"
	"net/http"
)

func main() {
	s, err := graphql.NewGraphQLServer("localhost:50052")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
