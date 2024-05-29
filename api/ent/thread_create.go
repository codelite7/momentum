// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/codelite7/momentum/api/ent/bookmark"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/ent/user"
)

// ThreadCreate is the builder for creating a Thread entity.
type ThreadCreate struct {
	config
	mutation *ThreadMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *ThreadCreate) SetCreatedAt(t time.Time) *ThreadCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *ThreadCreate) SetNillableCreatedAt(t *time.Time) *ThreadCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *ThreadCreate) SetUpdatedAt(t time.Time) *ThreadCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *ThreadCreate) SetNillableUpdatedAt(t *time.Time) *ThreadCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *ThreadCreate) SetName(s string) *ThreadCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetID sets the "id" field.
func (tc *ThreadCreate) SetID(pu pulid.ID) *ThreadCreate {
	tc.mutation.SetID(pu)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *ThreadCreate) SetNillableID(pu *pulid.ID) *ThreadCreate {
	if pu != nil {
		tc.SetID(*pu)
	}
	return tc
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (tc *ThreadCreate) SetCreatedByID(id pulid.ID) *ThreadCreate {
	tc.mutation.SetCreatedByID(id)
	return tc
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (tc *ThreadCreate) SetCreatedBy(u *User) *ThreadCreate {
	return tc.SetCreatedByID(u.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (tc *ThreadCreate) AddMessageIDs(ids ...pulid.ID) *ThreadCreate {
	tc.mutation.AddMessageIDs(ids...)
	return tc
}

// AddMessages adds the "messages" edges to the Message entity.
func (tc *ThreadCreate) AddMessages(m ...*Message) *ThreadCreate {
	ids := make([]pulid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tc.AddMessageIDs(ids...)
}

// AddBookmarkIDs adds the "bookmarks" edge to the Bookmark entity by IDs.
func (tc *ThreadCreate) AddBookmarkIDs(ids ...pulid.ID) *ThreadCreate {
	tc.mutation.AddBookmarkIDs(ids...)
	return tc
}

// AddBookmarks adds the "bookmarks" edges to the Bookmark entity.
func (tc *ThreadCreate) AddBookmarks(b ...*Bookmark) *ThreadCreate {
	ids := make([]pulid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tc.AddBookmarkIDs(ids...)
}

// SetParentID sets the "parent" edge to the Thread entity by ID.
func (tc *ThreadCreate) SetParentID(id pulid.ID) *ThreadCreate {
	tc.mutation.SetParentID(id)
	return tc
}

// SetNillableParentID sets the "parent" edge to the Thread entity by ID if the given value is not nil.
func (tc *ThreadCreate) SetNillableParentID(id *pulid.ID) *ThreadCreate {
	if id != nil {
		tc = tc.SetParentID(*id)
	}
	return tc
}

// SetParent sets the "parent" edge to the Thread entity.
func (tc *ThreadCreate) SetParent(t *Thread) *ThreadCreate {
	return tc.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Thread entity by IDs.
func (tc *ThreadCreate) AddChildIDs(ids ...pulid.ID) *ThreadCreate {
	tc.mutation.AddChildIDs(ids...)
	return tc
}

// AddChildren adds the "children" edges to the Thread entity.
func (tc *ThreadCreate) AddChildren(t ...*Thread) *ThreadCreate {
	ids := make([]pulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddChildIDs(ids...)
}

// Mutation returns the ThreadMutation object of the builder.
func (tc *ThreadCreate) Mutation() *ThreadMutation {
	return tc.mutation
}

// Save creates the Thread in the database.
func (tc *ThreadCreate) Save(ctx context.Context) (*Thread, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *ThreadCreate) SaveX(ctx context.Context) *Thread {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *ThreadCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *ThreadCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *ThreadCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := thread.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := thread.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := thread.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *ThreadCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Thread.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Thread.updated_at"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Thread.name"`)}
	}
	if _, ok := tc.mutation.CreatedByID(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required edge "Thread.created_by"`)}
	}
	return nil
}

func (tc *ThreadCreate) sqlSave(ctx context.Context) (*Thread, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*pulid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *ThreadCreate) createSpec() (*Thread, *sqlgraph.CreateSpec) {
	var (
		_node = &Thread{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(thread.Table, sqlgraph.NewFieldSpec(thread.FieldID, field.TypeString))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(thread.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(thread.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(thread.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := tc.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.CreatedByTable,
			Columns: []string{thread.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_threads = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.MessagesTable,
			Columns: []string{thread.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.BookmarksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.BookmarksTable,
			Columns: []string{thread.BookmarksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.ParentTable,
			Columns: []string{thread.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.thread_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.ChildrenTable,
			Columns: []string{thread.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ThreadCreateBulk is the builder for creating many Thread entities in bulk.
type ThreadCreateBulk struct {
	config
	err      error
	builders []*ThreadCreate
}

// Save creates the Thread entities in the database.
func (tcb *ThreadCreateBulk) Save(ctx context.Context) ([]*Thread, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Thread, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ThreadMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *ThreadCreateBulk) SaveX(ctx context.Context) []*Thread {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *ThreadCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *ThreadCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
