package schema

import (
	"entgo.io/contrib/entgql"
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
		// a message is sent by a user
		edge.From("sent_by", User.Type).Ref("messages").Unique().Required(),
		// a message belongs to a thread
		edge.From("thread", Thread.Type).Ref("messages").Unique().Required(),
		// a message may have many bookmarks
		edge.To("bookmarks", Bookmark.Type).Annotations(entgql.RelayConnection()),
		// a message has a response from an agent
		edge.To("response", Response.Type).Unique(),
	}
}
