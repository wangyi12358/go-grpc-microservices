package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	graphql2 "microservices/internal/gateway/graphql"
	"microservices/internal/gateway/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service *service.Service
}

// NewSchema creates a graphql executable schema.
func NewSchema(svc *service.Service) graphql.ExecutableSchema {
	return graphql2.NewExecutableSchema(graphql2.Config{
		Resolvers: &Resolver{service: svc},
	})
}
