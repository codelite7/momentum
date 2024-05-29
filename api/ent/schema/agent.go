package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
)

type Agent struct {
	ent.Schema
}

func (Agent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix("ag"),
	}
}

func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.String("provider").Annotations(entgql.OrderField("PROVIDER")),
		field.String("model").Annotations(entgql.OrderField("MODEL")),
		field.String("api_key"),
	}
}

func (Agent) Edges() []ent.Edge {
	return []ent.Edge{
		// an agent can respond to many messages
		edge.To("responses", Response.Type).Annotations(entgql.RelayConnection()),
	}
}
