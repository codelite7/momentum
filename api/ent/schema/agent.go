package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Agent struct {
	ent.Schema
}

func (Agent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(entgql.OrderField("NAME")),
		// TODO: model enum?
		field.String("model").Annotations(entgql.OrderField("MODEL")),
	}
}

func (Agent) Edges() []ent.Edge {
	return []ent.Edge{
		// an agent can send many messages
		edge.To("messages", Message.Type),
	}
}
