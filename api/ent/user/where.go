// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/codelite7/momentum/api/ent/predicate"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
)

// ID filters vertices based on their ID field.
func ID(id pulid.ID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.ID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.ID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.ID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.ID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.ID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.ID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.ID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.ID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v pulid.ID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldTenantID, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v pulid.ID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v pulid.ID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...pulid.ID) predicate.User {
	return predicate.User(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...pulid.ID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v pulid.ID) predicate.User {
	return predicate.User(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v pulid.ID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v pulid.ID) predicate.User {
	return predicate.User(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v pulid.ID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldTenantID, v))
}

// TenantIDContains applies the Contains predicate on the "tenant_id" field.
func TenantIDContains(v pulid.ID) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldContains(FieldTenantID, vc))
}

// TenantIDHasPrefix applies the HasPrefix predicate on the "tenant_id" field.
func TenantIDHasPrefix(v pulid.ID) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldHasPrefix(FieldTenantID, vc))
}

// TenantIDHasSuffix applies the HasSuffix predicate on the "tenant_id" field.
func TenantIDHasSuffix(v pulid.ID) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldHasSuffix(FieldTenantID, vc))
}

// TenantIDEqualFold applies the EqualFold predicate on the "tenant_id" field.
func TenantIDEqualFold(v pulid.ID) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldEqualFold(FieldTenantID, vc))
}

// TenantIDContainsFold applies the ContainsFold predicate on the "tenant_id" field.
func TenantIDContainsFold(v pulid.ID) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldContainsFold(FieldTenantID, vc))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// HasTenant applies the HasEdge predicate on the "tenant" edge.
func HasTenant() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TenantTable, TenantColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTenantWith applies the HasEdge predicate on the "tenant" edge with a given conditions (other predicates).
func HasTenantWith(preds ...predicate.Tenant) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newTenantStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBookmarks applies the HasEdge predicate on the "bookmarks" edge.
func HasBookmarks() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BookmarksTable, BookmarksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookmarksWith applies the HasEdge predicate on the "bookmarks" edge with a given conditions (other predicates).
func HasBookmarksWith(preds ...predicate.Bookmark) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newBookmarksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasThreads applies the HasEdge predicate on the "threads" edge.
func HasThreads() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ThreadsTable, ThreadsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasThreadsWith applies the HasEdge predicate on the "threads" edge with a given conditions (other predicates).
func HasThreadsWith(preds ...predicate.Thread) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newThreadsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMessages applies the HasEdge predicate on the "messages" edge.
func HasMessages() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessagesWith applies the HasEdge predicate on the "messages" edge with a given conditions (other predicates).
func HasMessagesWith(preds ...predicate.Message) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newMessagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
