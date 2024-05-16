// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/codelite7/momentum/api/ent/predicate"
	"github.com/codelite7/momentum/api/ent/response"
)

// ResponseDelete is the builder for deleting a Response entity.
type ResponseDelete struct {
	config
	hooks    []Hook
	mutation *ResponseMutation
}

// Where appends a list predicates to the ResponseDelete builder.
func (rd *ResponseDelete) Where(ps ...predicate.Response) *ResponseDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *ResponseDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rd.sqlExec, rd.mutation, rd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *ResponseDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *ResponseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(response.Table, sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID))
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rd.mutation.done = true
	return affected, err
}

// ResponseDeleteOne is the builder for deleting a single Response entity.
type ResponseDeleteOne struct {
	rd *ResponseDelete
}

// Where appends a list predicates to the ResponseDelete builder.
func (rdo *ResponseDeleteOne) Where(ps ...predicate.Response) *ResponseDeleteOne {
	rdo.rd.mutation.Where(ps...)
	return rdo
}

// Exec executes the deletion query.
func (rdo *ResponseDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{response.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *ResponseDeleteOne) ExecX(ctx context.Context) {
	if err := rdo.Exec(ctx); err != nil {
		panic(err)
	}
}
