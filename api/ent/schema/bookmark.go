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

type Bookmark struct {
	ent.Schema
}

func (Bookmark) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
	}
}

func (Bookmark) Edges() []ent.Edge {
	return []ent.Edge{
		// a bookmark is created by a user
		edge.From("user", User.Type).Ref("bookmarks").Unique().Required(),
		// a bookmark may be associated with a thread
		edge.From("thread", Thread.Type).Ref("bookmarks").Unique(),
		// a bookmark may be associated with a message
		edge.From("message", Message.Type).Ref("bookmarks").Unique(),
	}
}

func (Bookmark) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
