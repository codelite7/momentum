package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type Bookmark struct {
	ent.Schema
}

func (Bookmark) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Bookmark) Fields() []ent.Field {
	return []ent.Field{}
}

func (Bookmark) Edges() []ent.Edge {
	return []ent.Edge{
		// a bookmark is created by a user
		//edge.From("user", User.Type).Ref("bookmarks").Unique().Required(),
		// todo: make this required when we have better user stuff
		edge.From("user", User.Type).Ref("bookmarks").Unique(),
		// a bookmark may be associated with a thread
		edge.From("thread", Thread.Type).Ref("bookmarks").Unique(),
		// a bookmark may be associated with a message
		edge.From("message", Message.Type).Ref("bookmarks").Unique(),
	}
}
