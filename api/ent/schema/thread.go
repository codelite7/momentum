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

type Thread struct {
	ent.Schema
}

func (Thread) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
		field.String("name").Annotations(entgql.OrderField("NAME")),
	}
}

func (Thread) Edges() []ent.Edge {
	return []ent.Edge{
		// a thread is created by a user
		edge.From("created_by", User.Type).Ref("threads").Unique().Required(),
		// a thread may have many messages
		edge.To("messages", Message.Type),
		// a thread may be bookmarked many times
		edge.To("bookmarks", Bookmark.Type),
		// a thread may have a parent thread and a child thread
		edge.To("parent", Thread.Type).Unique().From("child").Unique(),
	}
}

func (Thread) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
