package net

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
	"go.uber.org/fx"
	"microservices/internal/gateway/graphql/resolver"
	"microservices/internal/gateway/service"
	"microservices/pkg/config"
	"microservices/pkg/middleware"
	"time"
)

// Defining the Graphql handler
func graphqlHandler(svc *service.Service) gin.HandlerFunc {
	srv := handler.New(resolver.NewSchema(svc))

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func NewHTTPServer(svc *service.Service) (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.GqlLoggerMiddleware(), middleware.GqlLogger(), middleware.Cors(), middleware.GinContext2Context())

	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/graphql").ServeHTTP(c.Writer, c.Request)
	})
	r.POST("/graphql", graphqlHandler(svc))
	return r, nil
}

func StartHTTPServer(lc fx.Lifecycle, r *gin.Engine, config *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				fmt.Printf("Server running at :%s \n", config.HTTP.Port)
				if err := r.Run(config.HTTP.Port); err != nil {
					fmt.Println("Server stopped with error:", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Server stopping...")
			return nil
		},
	})
}
