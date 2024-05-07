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
		//edge.From("created_by", User.Type).Ref("threads").Unique().Required(),
		// todo make this required when we have better user stuff
		edge.From("created_by", User.Type).Ref("threads").Unique(),
		// a thread may have many messages
		edge.To("messages", Message.Type),
		// a thread may be bookmarked many times
		edge.To("bookmarks", Bookmark.Type),
		// a thread may have a parent thread and a child thread
		edge.To("parent", Thread.Type).Unique().From("child").Unique(),
	}
}
