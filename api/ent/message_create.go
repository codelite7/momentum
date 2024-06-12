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
	"github.com/codelite7/momentum/api/ent/tenant"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/ent/user"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MessageCreate) SetCreatedAt(t time.Time) *MessageCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreatedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MessageCreate) SetUpdatedAt(t time.Time) *MessageCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableUpdatedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetTenantID sets the "tenant_id" field.
func (mc *MessageCreate) SetTenantID(pu pulid.ID) *MessageCreate {
	mc.mutation.SetTenantID(pu)
	return mc
}

// SetContent sets the "content" field.
func (mc *MessageCreate) SetContent(s string) *MessageCreate {
	mc.mutation.SetContent(s)
	return mc
}

// SetMessageType sets the "message_type" field.
func (mc *MessageCreate) SetMessageType(mt message.MessageType) *MessageCreate {
	mc.mutation.SetMessageType(mt)
	return mc
}

// SetID sets the "id" field.
func (mc *MessageCreate) SetID(pu pulid.ID) *MessageCreate {
	mc.mutation.SetID(pu)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MessageCreate) SetNillableID(pu *pulid.ID) *MessageCreate {
	if pu != nil {
		mc.SetID(*pu)
	}
	return mc
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (mc *MessageCreate) SetTenant(t *Tenant) *MessageCreate {
	return mc.SetTenantID(t.ID)
}

// SetSentByID sets the "sent_by" edge to the User entity by ID.
func (mc *MessageCreate) SetSentByID(id pulid.ID) *MessageCreate {
	mc.mutation.SetSentByID(id)
	return mc
}

// SetSentBy sets the "sent_by" edge to the User entity.
func (mc *MessageCreate) SetSentBy(u *User) *MessageCreate {
	return mc.SetSentByID(u.ID)
}

// SetThreadID sets the "thread" edge to the Thread entity by ID.
func (mc *MessageCreate) SetThreadID(id pulid.ID) *MessageCreate {
	mc.mutation.SetThreadID(id)
	return mc
}

// SetThread sets the "thread" edge to the Thread entity.
func (mc *MessageCreate) SetThread(t *Thread) *MessageCreate {
	return mc.SetThreadID(t.ID)
}

// AddBookmarkIDs adds the "bookmarks" edge to the Bookmark entity by IDs.
func (mc *MessageCreate) AddBookmarkIDs(ids ...pulid.ID) *MessageCreate {
	mc.mutation.AddBookmarkIDs(ids...)
	return mc
}

// AddBookmarks adds the "bookmarks" edges to the Bookmark entity.
func (mc *MessageCreate) AddBookmarks(b ...*Bookmark) *MessageCreate {
	ids := make([]pulid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return mc.AddBookmarkIDs(ids...)
}

// SetChildID sets the "child" edge to the Thread entity by ID.
func (mc *MessageCreate) SetChildID(id pulid.ID) *MessageCreate {
	mc.mutation.SetChildID(id)
	return mc
}

// SetNillableChildID sets the "child" edge to the Thread entity by ID if the given value is not nil.
func (mc *MessageCreate) SetNillableChildID(id *pulid.ID) *MessageCreate {
	if id != nil {
		mc = mc.SetChildID(*id)
	}
	return mc
}

// SetChild sets the "child" edge to the Thread entity.
func (mc *MessageCreate) SetChild(t *Thread) *MessageCreate {
	return mc.SetChildID(t.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessageCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := message.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := message.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := message.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Message.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Message.updated_at"`)}
	}
	if _, ok := mc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "Message.tenant_id"`)}
	}
	if _, ok := mc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Message.content"`)}
	}
	if _, ok := mc.mutation.MessageType(); !ok {
		return &ValidationError{Name: "message_type", err: errors.New(`ent: missing required field "Message.message_type"`)}
	}
	if v, ok := mc.mutation.MessageType(); ok {
		if err := message.MessageTypeValidator(v); err != nil {
			return &ValidationError{Name: "message_type", err: fmt.Errorf(`ent: validator failed for field "Message.message_type": %w`, err)}
		}
	}
	if _, ok := mc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant", err: errors.New(`ent: missing required edge "Message.tenant"`)}
	}
	if _, ok := mc.mutation.SentByID(); !ok {
		return &ValidationError{Name: "sent_by", err: errors.New(`ent: missing required edge "Message.sent_by"`)}
	}
	if _, ok := mc.mutation.ThreadID(); !ok {
		return &ValidationError{Name: "thread", err: errors.New(`ent: missing required edge "Message.thread"`)}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
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
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(message.Table, sqlgraph.NewFieldSpec(message.FieldID, field.TypeString))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(message.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(message.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := mc.mutation.MessageType(); ok {
		_spec.SetField(message.FieldMessageType, field.TypeEnum, value)
		_node.MessageType = value
	}
	if nodes := mc.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   message.TenantTable,
			Columns: []string{message.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TenantID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.SentByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.SentByTable,
			Columns: []string{message.SentByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_messages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ThreadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ThreadTable,
			Columns: []string{message.ThreadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.thread_messages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.BookmarksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.BookmarksTable,
			Columns: []string{message.BookmarksColumn},
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
	if nodes := mc.mutation.ChildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   message.ChildTable,
			Columns: []string{message.ChildColumn},
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

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	err      error
	builders []*MessageCreate
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}