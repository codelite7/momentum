package api

// run ent generation to generate ent code.
//go:generate go run -mod=mod ./ent/entc.go
// run gqlgen to generate graphql code from ent code
//go:generate go run -mod=mod github.com/99designs/gqlgen
