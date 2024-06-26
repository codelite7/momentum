package resolvers

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/codelite7/momentum/api/ent"
)

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client},
	})
}
