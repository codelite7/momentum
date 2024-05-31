package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
)

type Bookmark struct {
	ent.Schema
}

func (Bookmark) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix("bm"),
		TenantMixin{},
	}
}

func (Bookmark) Fields() []ent.Field {
	return []ent.Field{}
}

func (Bookmark) Edges() []ent.Edge {
	return []ent.Edge{
		// a bookmark is created by a user
		//edge.From("user", User.Type).Ref("bookmarks").Unique().Required(),
		edge.From("user", User.Type).Ref("bookmarks").Unique().Required(),
		// a bookmark may be associated with a thread
		edge.From("thread", Thread.Type).Ref("bookmarks").Unique(),
		// a bookmark may be associated with a message
		edge.From("message", Message.Type).Ref("bookmarks").Unique(),
		// a bookmark may be associated with a response
		edge.From("response", Response.Type).Ref("bookmarks").Unique(),
	}
}
