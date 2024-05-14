package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Thread struct {
	ent.Schema
}

func (Thread) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Thread) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(entgql.OrderField("NAME")),
	}
}

func (Thread) Edges() []ent.Edge {
	return []ent.Edge{
		// a thread is created by a user
		edge.From("created_by", User.Type).Ref("threads").Unique().Required(),
		// a thread may have many messages
		edge.To("messages", Message.Type).Annotations(entgql.RelayConnection()),
		// a thread may be bookmarked many times
		edge.To("bookmarks", Bookmark.Type).Annotations(entgql.RelayConnection()),
		// a thread may have a parent thread and many child threads
		edge.To("children", Thread.Type).Annotations(entgql.RelayConnection()).From("parent").Unique(),
	}
}
