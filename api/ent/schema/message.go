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

type Message struct {
	ent.Schema
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
		field.String("content"),
	}
}

func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		// a message is sent by a user
		edge.From("sent_by", User.Type).Ref("messages").Unique(),
		// a message belongs to a thread
		edge.From("thread", Thread.Type).Ref("messages").Unique().Required(),
		// a message may have many bookmarks
		edge.To("bookmarks", Bookmark.Type),
	}
}

func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
