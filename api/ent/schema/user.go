package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().Annotations(entgql.OrderField("EMAIL")),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// a user can have many bookmarks
		edge.To("bookmarks", Bookmark.Type).Annotations(entgql.RelayConnection()),
		// a user can have many threads
		edge.To("threads", Thread.Type).Annotations(entgql.RelayConnection()),
		// a user can send many messages
		edge.To("messages", Message.Type).Annotations(entgql.RelayConnection()),
	}
}
