package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Agent struct {
	ent.Schema
}

func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
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

func (Agent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
