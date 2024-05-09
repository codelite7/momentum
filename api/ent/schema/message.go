package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Message struct {
	ent.Schema
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
	}
}

func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		// a message can be sent by a user
		// todo: constraint on one or the other
		edge.From("sent_by_agent", Agent.Type).Ref("messages").Unique(),
		// a message can be sent by an agent
		edge.From("sent_by_user", User.Type).Ref("messages").Unique(),
		// a message belongs to a thread
		edge.From("thread", Thread.Type).Ref("messages").Unique().Required(),
		// a message may have many bookmarks
		edge.To("bookmarks", Bookmark.Type),
	}
}
