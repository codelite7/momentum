// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/codelite7/momentum/api/ent/bookmark"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/predicate"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"github.com/codelite7/momentum/api/ent/user"
)

// BookmarkUpdate is the builder for updating Bookmark entities.
type BookmarkUpdate struct {
	config
	hooks    []Hook
	mutation *BookmarkMutation
}

// Where appends a list predicates to the BookmarkUpdate builder.
func (bu *BookmarkUpdate) Where(ps ...predicate.Bookmark) *BookmarkUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BookmarkUpdate) SetUpdatedAt(t time.Time) *BookmarkUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bu *BookmarkUpdate) SetNillableUpdatedAt(t *time.Time) *BookmarkUpdate {
	if t != nil {
		bu.SetUpdatedAt(*t)
	}
	return bu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bu *BookmarkUpdate) SetUserID(id pulid.ID) *BookmarkUpdate {
	bu.mutation.SetUserID(id)
	return bu
}

// SetUser sets the "user" edge to the User entity.
func (bu *BookmarkUpdate) SetUser(u *User) *BookmarkUpdate {
	return bu.SetUserID(u.ID)
}

// SetMessageID sets the "message" edge to the Message entity by ID.
func (bu *BookmarkUpdate) SetMessageID(id pulid.ID) *BookmarkUpdate {
	bu.mutation.SetMessageID(id)
	return bu
}

// SetNillableMessageID sets the "message" edge to the Message entity by ID if the given value is not nil.
func (bu *BookmarkUpdate) SetNillableMessageID(id *pulid.ID) *BookmarkUpdate {
	if id != nil {
		bu = bu.SetMessageID(*id)
	}
	return bu
}

// SetMessage sets the "message" edge to the Message entity.
func (bu *BookmarkUpdate) SetMessage(m *Message) *BookmarkUpdate {
	return bu.SetMessageID(m.ID)
}

// Mutation returns the BookmarkMutation object of the builder.
func (bu *BookmarkUpdate) Mutation() *BookmarkMutation {
	return bu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (bu *BookmarkUpdate) ClearUser() *BookmarkUpdate {
	bu.mutation.ClearUser()
	return bu
}

// ClearMessage clears the "message" edge to the Message entity.
func (bu *BookmarkUpdate) ClearMessage() *BookmarkUpdate {
	bu.mutation.ClearMessage()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookmarkUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookmarkUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookmarkUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookmarkUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookmarkUpdate) check() error {
	if _, ok := bu.mutation.TenantID(); bu.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Bookmark.tenant"`)
	}
	if _, ok := bu.mutation.UserID(); bu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Bookmark.user"`)
	}
	return nil
}

func (bu *BookmarkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(bookmark.Table, bookmark.Columns, sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeString))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(bookmark.FieldUpdatedAt, field.TypeTime, value)
	}
	if bu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.UserTable,
			Columns: []string{bookmark.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.UserTable,
			Columns: []string{bookmark.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.MessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.MessageTable,
			Columns: []string{bookmark.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.MessageTable,
			Columns: []string{bookmark.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookmark.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookmarkUpdateOne is the builder for updating a single Bookmark entity.
type BookmarkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookmarkMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BookmarkUpdateOne) SetUpdatedAt(t time.Time) *BookmarkUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (buo *BookmarkUpdateOne) SetNillableUpdatedAt(t *time.Time) *BookmarkUpdateOne {
	if t != nil {
		buo.SetUpdatedAt(*t)
	}
	return buo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (buo *BookmarkUpdateOne) SetUserID(id pulid.ID) *BookmarkUpdateOne {
	buo.mutation.SetUserID(id)
	return buo
}

// SetUser sets the "user" edge to the User entity.
func (buo *BookmarkUpdateOne) SetUser(u *User) *BookmarkUpdateOne {
	return buo.SetUserID(u.ID)
}

// SetMessageID sets the "message" edge to the Message entity by ID.
func (buo *BookmarkUpdateOne) SetMessageID(id pulid.ID) *BookmarkUpdateOne {
	buo.mutation.SetMessageID(id)
	return buo
}

// SetNillableMessageID sets the "message" edge to the Message entity by ID if the given value is not nil.
func (buo *BookmarkUpdateOne) SetNillableMessageID(id *pulid.ID) *BookmarkUpdateOne {
	if id != nil {
		buo = buo.SetMessageID(*id)
	}
	return buo
}

// SetMessage sets the "message" edge to the Message entity.
func (buo *BookmarkUpdateOne) SetMessage(m *Message) *BookmarkUpdateOne {
	return buo.SetMessageID(m.ID)
}

// Mutation returns the BookmarkMutation object of the builder.
func (buo *BookmarkUpdateOne) Mutation() *BookmarkMutation {
	return buo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (buo *BookmarkUpdateOne) ClearUser() *BookmarkUpdateOne {
	buo.mutation.ClearUser()
	return buo
}

// ClearMessage clears the "message" edge to the Message entity.
func (buo *BookmarkUpdateOne) ClearMessage() *BookmarkUpdateOne {
	buo.mutation.ClearMessage()
	return buo
}

// Where appends a list predicates to the BookmarkUpdate builder.
func (buo *BookmarkUpdateOne) Where(ps ...predicate.Bookmark) *BookmarkUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookmarkUpdateOne) Select(field string, fields ...string) *BookmarkUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Bookmark entity.
func (buo *BookmarkUpdateOne) Save(ctx context.Context) (*Bookmark, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookmarkUpdateOne) SaveX(ctx context.Context) *Bookmark {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookmarkUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookmarkUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookmarkUpdateOne) check() error {
	if _, ok := buo.mutation.TenantID(); buo.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Bookmark.tenant"`)
	}
	if _, ok := buo.mutation.UserID(); buo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Bookmark.user"`)
	}
	return nil
}

func (buo *BookmarkUpdateOne) sqlSave(ctx context.Context) (_node *Bookmark, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(bookmark.Table, bookmark.Columns, sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeString))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Bookmark.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookmark.FieldID)
		for _, f := range fields {
			if !bookmark.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bookmark.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(bookmark.FieldUpdatedAt, field.TypeTime, value)
	}
	if buo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.UserTable,
			Columns: []string{bookmark.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.UserTable,
			Columns: []string{bookmark.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.MessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.MessageTable,
			Columns: []string{bookmark.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.MessageTable,
			Columns: []string{bookmark.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Bookmark{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookmark.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
