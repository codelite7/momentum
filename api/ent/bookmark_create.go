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
	"github.com/codelite7/momentum/api/ent/response"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/ent/user"
)

// BookmarkCreate is the builder for creating a Bookmark entity.
type BookmarkCreate struct {
	config
	mutation *BookmarkMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bc *BookmarkCreate) SetCreatedAt(t time.Time) *BookmarkCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BookmarkCreate) SetNillableCreatedAt(t *time.Time) *BookmarkCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BookmarkCreate) SetUpdatedAt(t time.Time) *BookmarkCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BookmarkCreate) SetNillableUpdatedAt(t *time.Time) *BookmarkCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BookmarkCreate) SetID(pu pulid.ID) *BookmarkCreate {
	bc.mutation.SetID(pu)
	return bc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bc *BookmarkCreate) SetNillableID(pu *pulid.ID) *BookmarkCreate {
	if pu != nil {
		bc.SetID(*pu)
	}
	return bc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bc *BookmarkCreate) SetUserID(id pulid.ID) *BookmarkCreate {
	bc.mutation.SetUserID(id)
	return bc
}

// SetUser sets the "user" edge to the User entity.
func (bc *BookmarkCreate) SetUser(u *User) *BookmarkCreate {
	return bc.SetUserID(u.ID)
}

// SetThreadID sets the "thread" edge to the Thread entity by ID.
func (bc *BookmarkCreate) SetThreadID(id pulid.ID) *BookmarkCreate {
	bc.mutation.SetThreadID(id)
	return bc
}

// SetNillableThreadID sets the "thread" edge to the Thread entity by ID if the given value is not nil.
func (bc *BookmarkCreate) SetNillableThreadID(id *pulid.ID) *BookmarkCreate {
	if id != nil {
		bc = bc.SetThreadID(*id)
	}
	return bc
}

// SetThread sets the "thread" edge to the Thread entity.
func (bc *BookmarkCreate) SetThread(t *Thread) *BookmarkCreate {
	return bc.SetThreadID(t.ID)
}

// SetMessageID sets the "message" edge to the Message entity by ID.
func (bc *BookmarkCreate) SetMessageID(id pulid.ID) *BookmarkCreate {
	bc.mutation.SetMessageID(id)
	return bc
}

// SetNillableMessageID sets the "message" edge to the Message entity by ID if the given value is not nil.
func (bc *BookmarkCreate) SetNillableMessageID(id *pulid.ID) *BookmarkCreate {
	if id != nil {
		bc = bc.SetMessageID(*id)
	}
	return bc
}

// SetMessage sets the "message" edge to the Message entity.
func (bc *BookmarkCreate) SetMessage(m *Message) *BookmarkCreate {
	return bc.SetMessageID(m.ID)
}

// SetResponseID sets the "response" edge to the Response entity by ID.
func (bc *BookmarkCreate) SetResponseID(id pulid.ID) *BookmarkCreate {
	bc.mutation.SetResponseID(id)
	return bc
}

// SetNillableResponseID sets the "response" edge to the Response entity by ID if the given value is not nil.
func (bc *BookmarkCreate) SetNillableResponseID(id *pulid.ID) *BookmarkCreate {
	if id != nil {
		bc = bc.SetResponseID(*id)
	}
	return bc
}

// SetResponse sets the "response" edge to the Response entity.
func (bc *BookmarkCreate) SetResponse(r *Response) *BookmarkCreate {
	return bc.SetResponseID(r.ID)
}

// Mutation returns the BookmarkMutation object of the builder.
func (bc *BookmarkCreate) Mutation() *BookmarkMutation {
	return bc.mutation
}

// Save creates the Bookmark in the database.
func (bc *BookmarkCreate) Save(ctx context.Context) (*Bookmark, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookmarkCreate) SaveX(ctx context.Context) *Bookmark {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookmarkCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookmarkCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BookmarkCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := bookmark.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := bookmark.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
	if _, ok := bc.mutation.ID(); !ok {
		v := bookmark.DefaultID()
		bc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookmarkCreate) check() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Bookmark.created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Bookmark.updated_at"`)}
	}
	if _, ok := bc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Bookmark.user"`)}
	}
	return nil
}

func (bc *BookmarkCreate) sqlSave(ctx context.Context) (*Bookmark, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
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
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BookmarkCreate) createSpec() (*Bookmark, *sqlgraph.CreateSpec) {
	var (
		_node = &Bookmark{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(bookmark.Table, sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeString))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(bookmark.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(bookmark.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := bc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.user_bookmarks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.ThreadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.ThreadTable,
			Columns: []string{bookmark.ThreadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.thread_bookmarks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.MessageIDs(); len(nodes) > 0 {
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
		_node.message_bookmarks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.ResponseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.ResponseTable,
			Columns: []string{bookmark.ResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.response_bookmarks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BookmarkCreateBulk is the builder for creating many Bookmark entities in bulk.
type BookmarkCreateBulk struct {
	config
	err      error
	builders []*BookmarkCreate
}

// Save creates the Bookmark entities in the database.
func (bcb *BookmarkCreateBulk) Save(ctx context.Context) ([]*Bookmark, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Bookmark, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookmarkMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookmarkCreateBulk) SaveX(ctx context.Context) []*Bookmark {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookmarkCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookmarkCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
