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
	"github.com/codelite7/momentum/api/ent/response"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
)

// AgentCreate is the builder for creating a Agent entity.
type AgentCreate struct {
	config
	mutation *AgentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AgentCreate) SetCreatedAt(t time.Time) *AgentCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AgentCreate) SetNillableCreatedAt(t *time.Time) *AgentCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AgentCreate) SetUpdatedAt(t time.Time) *AgentCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AgentCreate) SetNillableUpdatedAt(t *time.Time) *AgentCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetProvider sets the "provider" field.
func (ac *AgentCreate) SetProvider(s string) *AgentCreate {
	ac.mutation.SetProvider(s)
	return ac
}

// SetModel sets the "model" field.
func (ac *AgentCreate) SetModel(s string) *AgentCreate {
	ac.mutation.SetModel(s)
	return ac
}

// SetAPIKey sets the "api_key" field.
func (ac *AgentCreate) SetAPIKey(s string) *AgentCreate {
	ac.mutation.SetAPIKey(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AgentCreate) SetID(pu pulid.ID) *AgentCreate {
	ac.mutation.SetID(pu)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AgentCreate) SetNillableID(pu *pulid.ID) *AgentCreate {
	if pu != nil {
		ac.SetID(*pu)
	}
	return ac
}

// AddResponseIDs adds the "responses" edge to the Response entity by IDs.
func (ac *AgentCreate) AddResponseIDs(ids ...pulid.ID) *AgentCreate {
	ac.mutation.AddResponseIDs(ids...)
	return ac
}

// AddResponses adds the "responses" edges to the Response entity.
func (ac *AgentCreate) AddResponses(r ...*Response) *AgentCreate {
	ids := make([]pulid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ac.AddResponseIDs(ids...)
}

// Mutation returns the AgentMutation object of the builder.
func (ac *AgentCreate) Mutation() *AgentMutation {
	return ac.mutation
}

// Save creates the Agent in the database.
func (ac *AgentCreate) Save(ctx context.Context) (*Agent, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AgentCreate) SaveX(ctx context.Context) *Agent {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AgentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AgentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AgentCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := agent.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := agent.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := agent.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AgentCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Agent.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Agent.updated_at"`)}
	}
	if _, ok := ac.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`ent: missing required field "Agent.provider"`)}
	}
	if _, ok := ac.mutation.Model(); !ok {
		return &ValidationError{Name: "model", err: errors.New(`ent: missing required field "Agent.model"`)}
	}
	if _, ok := ac.mutation.APIKey(); !ok {
		return &ValidationError{Name: "api_key", err: errors.New(`ent: missing required field "Agent.api_key"`)}
	}
	return nil
}

func (ac *AgentCreate) sqlSave(ctx context.Context) (*Agent, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AgentCreate) createSpec() (*Agent, *sqlgraph.CreateSpec) {
	var (
		_node = &Agent{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(agent.Table, sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(agent.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(agent.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Provider(); ok {
		_spec.SetField(agent.FieldProvider, field.TypeString, value)
		_node.Provider = value
	}
	if value, ok := ac.mutation.Model(); ok {
		_spec.SetField(agent.FieldModel, field.TypeString, value)
		_node.Model = value
	}
	if value, ok := ac.mutation.APIKey(); ok {
		_spec.SetField(agent.FieldAPIKey, field.TypeString, value)
		_node.APIKey = value
	}
	if nodes := ac.mutation.ResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.ResponsesTable,
			Columns: []string{agent.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AgentCreateBulk is the builder for creating many Agent entities in bulk.
type AgentCreateBulk struct {
	config
	err      error
	builders []*AgentCreate
}

// Save creates the Agent entities in the database.
func (acb *AgentCreateBulk) Save(ctx context.Context) ([]*Agent, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Agent, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AgentMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AgentCreateBulk) SaveX(ctx context.Context) []*Agent {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AgentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AgentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
