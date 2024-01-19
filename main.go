package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"go-graphql-cli/adapters/resolvers"
	"go-graphql-cli/cmd"
	"go-graphql-cli/infra/db"
	"go-graphql-cli/infra/graph"
)

const defaultPort = "8080"

func main() {
	// cli
	cmd.Execute()

	// db connection
	gormDB, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// run GraphQL server
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolvers.NewResolver(gormDB),
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
