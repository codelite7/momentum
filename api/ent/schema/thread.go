package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"time"
)

type Thread struct {
	ent.Schema
}

func (Thread) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix("tr"),
		TenantMixin{},
	}
}

func (Thread) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(entgql.OrderField("NAME")),
		field.Time("last_viewed_at").Default(time.Now).Annotations(entgql.OrderField("LAST_VIEWED_AT")),
	}
}

func (Thread) Edges() []ent.Edge {
	return []ent.Edge{
		// a thread is created by a user
		edge.From("created_by", User.Type).Ref("threads").Unique().Required().Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		// a thread may have many messages
		edge.To("messages", Message.Type).Annotations(entgql.RelayConnection(), entsql.OnDelete(entsql.Cascade)),
		// a thread may be bookmarked many times
		edge.To("bookmarks", Bookmark.Type).Annotations(entgql.RelayConnection(), entsql.OnDelete(entsql.Cascade)),
		// a thread may have a parent thread and many child threads
		edge.To("children", Thread.Type).Annotations(entgql.RelayConnection(), entsql.OnDelete(entsql.Cascade)).From("parent").Unique(),
	}
}
