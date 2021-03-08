package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sunzhongshan1988/army-ant/broker/graph"
	"github.com/sunzhongshan1988/army-ant/broker/graph/generated"
)

const defaultPort = "8080"

func Graphql() {
	// Start server
	log.Printf("--Start Graphql Server")

	port := os.Getenv("GRAPHQL_PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
