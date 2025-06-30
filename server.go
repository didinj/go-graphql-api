package main

import (
	"log"
	"net/http"

	"github.com/didinj/go-graphql-api/db" // <-- replace with your actual module name

	"github.com/didinj/go-graphql-api/graph"
	"github.com/didinj/go-graphql-api/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("❌ DB connection failed: %v", err)
	}

	port := defaultPort
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 Server is running at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
