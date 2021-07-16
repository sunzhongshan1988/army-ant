package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sunzhongshan1988/army-ant/broker/config"
	"github.com/sunzhongshan1988/army-ant/broker/graph"
	"github.com/sunzhongshan1988/army-ant/broker/graph/generated"
	"log"
	"net/http"
)

func Graphql() {
	// Start server
	log.Printf("[system, graphql] info: Start Graphql Server")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("[system, graphql] info: connect to http://localhost:%s/ for GraphQL playground", config.GetGraphQLPort())
	log.Fatal(http.ListenAndServe(":"+config.GetGraphQLPort(), nil))
}
