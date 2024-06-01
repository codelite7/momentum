package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
)

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix("us"),
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().Annotations(entgql.OrderField("EMAIL")),
		field.String("workos_user_id").Unique().Immutable().NotEmpty().Annotations(entgql.Skip()),
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
		// a user can belong to many tenants
		edge.To("tenants", Tenant.Type).
			Required().
			Annotations(entgql.Skip()),
		// a user can only have one active tenant at a time
		edge.To("active_tenant", Tenant.Type).Unique().Required().Annotations(entgql.Skip()),
	}
}
