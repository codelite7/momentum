// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/codelite7/momentum/api/ent/agent"
	"github.com/codelite7/momentum/api/ent/bookmark"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/response"
	"github.com/google/uuid"
)

// ResponseCreate is the builder for creating a Response entity.
type ResponseCreate struct {
	config
	mutation *ResponseMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rc *ResponseCreate) SetCreatedAt(t time.Time) *ResponseCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *ResponseCreate) SetNillableCreatedAt(t *time.Time) *ResponseCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *ResponseCreate) SetUpdatedAt(t time.Time) *ResponseCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *ResponseCreate) SetNillableUpdatedAt(t *time.Time) *ResponseCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetContent sets the "content" field.
func (rc *ResponseCreate) SetContent(s string) *ResponseCreate {
	rc.mutation.SetContent(s)
	return rc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (rc *ResponseCreate) SetNillableContent(s *string) *ResponseCreate {
	if s != nil {
		rc.SetContent(*s)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *ResponseCreate) SetID(u uuid.UUID) *ResponseCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *ResponseCreate) SetNillableID(u *uuid.UUID) *ResponseCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// SetSentByID sets the "sent_by" edge to the Agent entity by ID.
func (rc *ResponseCreate) SetSentByID(id uuid.UUID) *ResponseCreate {
	rc.mutation.SetSentByID(id)
	return rc
}

// SetSentBy sets the "sent_by" edge to the Agent entity.
func (rc *ResponseCreate) SetSentBy(a *Agent) *ResponseCreate {
	return rc.SetSentByID(a.ID)
}

// SetMessageID sets the "message" edge to the Message entity by ID.
func (rc *ResponseCreate) SetMessageID(id uuid.UUID) *ResponseCreate {
	rc.mutation.SetMessageID(id)
	return rc
}

// SetMessage sets the "message" edge to the Message entity.
func (rc *ResponseCreate) SetMessage(m *Message) *ResponseCreate {
	return rc.SetMessageID(m.ID)
}

// AddBookmarkIDs adds the "bookmarks" edge to the Bookmark entity by IDs.
func (rc *ResponseCreate) AddBookmarkIDs(ids ...uuid.UUID) *ResponseCreate {
	rc.mutation.AddBookmarkIDs(ids...)
	return rc
}

// AddBookmarks adds the "bookmarks" edges to the Bookmark entity.
func (rc *ResponseCreate) AddBookmarks(b ...*Bookmark) *ResponseCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return rc.AddBookmarkIDs(ids...)
}

// Mutation returns the ResponseMutation object of the builder.
func (rc *ResponseCreate) Mutation() *ResponseMutation {
	return rc.mutation
}

// Save creates the Response in the database.
func (rc *ResponseCreate) Save(ctx context.Context) (*Response, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ResponseCreate) SaveX(ctx context.Context) *Response {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ResponseCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ResponseCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ResponseCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := response.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := response.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := response.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ResponseCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Response.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Response.updated_at"`)}
	}
	if _, ok := rc.mutation.SentByID(); !ok {
		return &ValidationError{Name: "sent_by", err: errors.New(`ent: missing required edge "Response.sent_by"`)}
	}
	if _, ok := rc.mutation.MessageID(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required edge "Response.message"`)}
	}
	return nil
}

func (rc *ResponseCreate) sqlSave(ctx context.Context) (*Response, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ResponseCreate) createSpec() (*Response, *sqlgraph.CreateSpec) {
	var (
		_node = &Response{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(response.Table, sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(response.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(response.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.Content(); ok {
		_spec.SetField(response.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if nodes := rc.mutation.SentByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   response.SentByTable,
			Columns: []string{response.SentByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.agent_responses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   response.MessageTable,
			Columns: []string{response.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.message_response = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.BookmarksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   response.BookmarksTable,
			Columns: []string{response.BookmarksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ResponseCreateBulk is the builder for creating many Response entities in bulk.
type ResponseCreateBulk struct {
	config
	err      error
	builders []*ResponseCreate
}

// Save creates the Response entities in the database.
func (rcb *ResponseCreateBulk) Save(ctx context.Context) ([]*Response, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Response, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResponseMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ResponseCreateBulk) SaveX(ctx context.Context) []*Response {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ResponseCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ResponseCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
