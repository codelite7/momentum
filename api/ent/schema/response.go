package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Response struct {
	ent.Schema
}

func (Response) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Response) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").Optional(),
	}
}

func (Response) Edges() []ent.Edge {
	return []ent.Edge{
		// a response is sent by an agent
		edge.From("sent_by", Agent.Type).Ref("responses").Unique().Required(),
		// a response is connected to a message
		edge.From("message", Message.Type).Ref("response").Unique().Required(),
		// a response may have many bookmarks
		edge.To("bookmarks", Bookmark.Type).Annotations(entgql.RelayConnection()),
	}
}
