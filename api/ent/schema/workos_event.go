package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type WorkosEventCursor struct {
	ent.Schema
}

func (WorkosEventCursor) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_created_cursor").Optional(),
	}
}

func (WorkosEventCursor) Annotations() []schema.Annotation {
	return []schema.Annotation{entgql.Skip()}
}
