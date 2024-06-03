package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
)

type Agent struct {
	ent.Schema
}

// agents are global, they don't belong to a tenant
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
		field.String("api_key").Annotations(entgql.Skip()),
	}
}

func (Agent) Edges() []ent.Edge {
	return []ent.Edge{}
}
