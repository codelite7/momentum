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
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/ent/user"
	"github.com/google/uuid"
)

// ThreadUpdate is the builder for updating Thread entities.
type ThreadUpdate struct {
	config
	hooks    []Hook
	mutation *ThreadMutation
}

// Where appends a list predicates to the ThreadUpdate builder.
func (tu *ThreadUpdate) Where(ps ...predicate.Thread) *ThreadUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *ThreadUpdate) SetUpdatedAt(t time.Time) *ThreadUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tu *ThreadUpdate) SetNillableUpdatedAt(t *time.Time) *ThreadUpdate {
	if t != nil {
		tu.SetUpdatedAt(*t)
	}
	return tu
}

// SetName sets the "name" field.
func (tu *ThreadUpdate) SetName(s string) *ThreadUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *ThreadUpdate) SetNillableName(s *string) *ThreadUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (tu *ThreadUpdate) SetCreatedByID(id uuid.UUID) *ThreadUpdate {
	tu.mutation.SetCreatedByID(id)
	return tu
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (tu *ThreadUpdate) SetNillableCreatedByID(id *uuid.UUID) *ThreadUpdate {
	if id != nil {
		tu = tu.SetCreatedByID(*id)
	}
	return tu
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (tu *ThreadUpdate) SetCreatedBy(u *User) *ThreadUpdate {
	return tu.SetCreatedByID(u.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (tu *ThreadUpdate) AddMessageIDs(ids ...uuid.UUID) *ThreadUpdate {
	tu.mutation.AddMessageIDs(ids...)
	return tu
}

// AddMessages adds the "messages" edges to the Message entity.
func (tu *ThreadUpdate) AddMessages(m ...*Message) *ThreadUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tu.AddMessageIDs(ids...)
}

// AddBookmarkIDs adds the "bookmarks" edge to the Bookmark entity by IDs.
func (tu *ThreadUpdate) AddBookmarkIDs(ids ...uuid.UUID) *ThreadUpdate {
	tu.mutation.AddBookmarkIDs(ids...)
	return tu
}

// AddBookmarks adds the "bookmarks" edges to the Bookmark entity.
func (tu *ThreadUpdate) AddBookmarks(b ...*Bookmark) *ThreadUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tu.AddBookmarkIDs(ids...)
}

// SetChildID sets the "child" edge to the Thread entity by ID.
func (tu *ThreadUpdate) SetChildID(id uuid.UUID) *ThreadUpdate {
	tu.mutation.SetChildID(id)
	return tu
}

// SetNillableChildID sets the "child" edge to the Thread entity by ID if the given value is not nil.
func (tu *ThreadUpdate) SetNillableChildID(id *uuid.UUID) *ThreadUpdate {
	if id != nil {
		tu = tu.SetChildID(*id)
	}
	return tu
}

// SetChild sets the "child" edge to the Thread entity.
func (tu *ThreadUpdate) SetChild(t *Thread) *ThreadUpdate {
	return tu.SetChildID(t.ID)
}

// SetParentID sets the "parent" edge to the Thread entity by ID.
func (tu *ThreadUpdate) SetParentID(id uuid.UUID) *ThreadUpdate {
	tu.mutation.SetParentID(id)
	return tu
}

// SetNillableParentID sets the "parent" edge to the Thread entity by ID if the given value is not nil.
func (tu *ThreadUpdate) SetNillableParentID(id *uuid.UUID) *ThreadUpdate {
	if id != nil {
		tu = tu.SetParentID(*id)
	}
	return tu
}

// SetParent sets the "parent" edge to the Thread entity.
func (tu *ThreadUpdate) SetParent(t *Thread) *ThreadUpdate {
	return tu.SetParentID(t.ID)
}

// Mutation returns the ThreadMutation object of the builder.
func (tu *ThreadUpdate) Mutation() *ThreadMutation {
	return tu.mutation
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (tu *ThreadUpdate) ClearCreatedBy() *ThreadUpdate {
	tu.mutation.ClearCreatedBy()
	return tu
}

// ClearMessages clears all "messages" edges to the Message entity.
func (tu *ThreadUpdate) ClearMessages() *ThreadUpdate {
	tu.mutation.ClearMessages()
	return tu
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (tu *ThreadUpdate) RemoveMessageIDs(ids ...uuid.UUID) *ThreadUpdate {
	tu.mutation.RemoveMessageIDs(ids...)
	return tu
}

// RemoveMessages removes "messages" edges to Message entities.
func (tu *ThreadUpdate) RemoveMessages(m ...*Message) *ThreadUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tu.RemoveMessageIDs(ids...)
}

// ClearBookmarks clears all "bookmarks" edges to the Bookmark entity.
func (tu *ThreadUpdate) ClearBookmarks() *ThreadUpdate {
	tu.mutation.ClearBookmarks()
	return tu
}

// RemoveBookmarkIDs removes the "bookmarks" edge to Bookmark entities by IDs.
func (tu *ThreadUpdate) RemoveBookmarkIDs(ids ...uuid.UUID) *ThreadUpdate {
	tu.mutation.RemoveBookmarkIDs(ids...)
	return tu
}

// RemoveBookmarks removes "bookmarks" edges to Bookmark entities.
func (tu *ThreadUpdate) RemoveBookmarks(b ...*Bookmark) *ThreadUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tu.RemoveBookmarkIDs(ids...)
}

// ClearChild clears the "child" edge to the Thread entity.
func (tu *ThreadUpdate) ClearChild() *ThreadUpdate {
	tu.mutation.ClearChild()
	return tu
}

// ClearParent clears the "parent" edge to the Thread entity.
func (tu *ThreadUpdate) ClearParent() *ThreadUpdate {
	tu.mutation.ClearParent()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *ThreadUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *ThreadUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *ThreadUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *ThreadUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *ThreadUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(thread.Table, thread.Columns, sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(thread.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(thread.FieldName, field.TypeString, value)
	}
	if tu.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.CreatedByTable,
			Columns: []string{thread.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.CreatedByTable,
			Columns: []string{thread.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.MessagesTable,
			Columns: []string{thread.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !tu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.MessagesTable,
			Columns: []string{thread.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.MessagesTable,
			Columns: []string{thread.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.BookmarksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.BookmarksTable,
			Columns: []string{thread.BookmarksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedBookmarksIDs(); len(nodes) > 0 && !tu.mutation.BookmarksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.BookmarksTable,
			Columns: []string{thread.BookmarksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.BookmarksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.BookmarksTable,
			Columns: []string{thread.BookmarksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ChildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   thread.ChildTable,
			Columns: []string{thread.ChildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ChildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   thread.ChildTable,
			Columns: []string{thread.ChildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   thread.ParentTable,
			Columns: []string{thread.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   thread.ParentTable,
			Columns: []string{thread.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// ThreadUpdateOne is the builder for updating a single Thread entity.
type ThreadUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ThreadMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *ThreadUpdateOne) SetUpdatedAt(t time.Time) *ThreadUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuo *ThreadUpdateOne) SetNillableUpdatedAt(t *time.Time) *ThreadUpdateOne {
	if t != nil {
		tuo.SetUpdatedAt(*t)
	}
	return tuo
}

// SetName sets the "name" field.
func (tuo *ThreadUpdateOne) SetName(s string) *ThreadUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *ThreadUpdateOne) SetNillableName(s *string) *ThreadUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (tuo *ThreadUpdateOne) SetCreatedByID(id uuid.UUID) *ThreadUpdateOne {
	tuo.mutation.SetCreatedByID(id)
	return tuo
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (tuo *ThreadUpdateOne) SetNillableCreatedByID(id *uuid.UUID) *ThreadUpdateOne {
	if id != nil {
		tuo = tuo.SetCreatedByID(*id)
	}
	return tuo
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (tuo *ThreadUpdateOne) SetCreatedBy(u *User) *ThreadUpdateOne {
	return tuo.SetCreatedByID(u.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (tuo *ThreadUpdateOne) AddMessageIDs(ids ...uuid.UUID) *ThreadUpdateOne {
	tuo.mutation.AddMessageIDs(ids...)
	return tuo
}

// AddMessages adds the "messages" edges to the Message entity.
func (tuo *ThreadUpdateOne) AddMessages(m ...*Message) *ThreadUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tuo.AddMessageIDs(ids...)
}

// AddBookmarkIDs adds the "bookmarks" edge to the Bookmark entity by IDs.
func (tuo *ThreadUpdateOne) AddBookmarkIDs(ids ...uuid.UUID) *ThreadUpdateOne {
	tuo.mutation.AddBookmarkIDs(ids...)
	return tuo
}

// AddBookmarks adds the "bookmarks" edges to the Bookmark entity.
func (tuo *ThreadUpdateOne) AddBookmarks(b ...*Bookmark) *ThreadUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tuo.AddBookmarkIDs(ids...)
}

// SetChildID sets the "child" edge to the Thread entity by ID.
func (tuo *ThreadUpdateOne) SetChildID(id uuid.UUID) *ThreadUpdateOne {
	tuo.mutation.SetChildID(id)
	return tuo
}

// SetNillableChildID sets the "child" edge to the Thread entity by ID if the given value is not nil.
func (tuo *ThreadUpdateOne) SetNillableChildID(id *uuid.UUID) *ThreadUpdateOne {
	if id != nil {
		tuo = tuo.SetChildID(*id)
	}
	return tuo
}

// SetChild sets the "child" edge to the Thread entity.
func (tuo *ThreadUpdateOne) SetChild(t *Thread) *ThreadUpdateOne {
	return tuo.SetChildID(t.ID)
}

// SetParentID sets the "parent" edge to the Thread entity by ID.
func (tuo *ThreadUpdateOne) SetParentID(id uuid.UUID) *ThreadUpdateOne {
	tuo.mutation.SetParentID(id)
	return tuo
}

// SetNillableParentID sets the "parent" edge to the Thread entity by ID if the given value is not nil.
func (tuo *ThreadUpdateOne) SetNillableParentID(id *uuid.UUID) *ThreadUpdateOne {
	if id != nil {
		tuo = tuo.SetParentID(*id)
	}
	return tuo
}

// SetParent sets the "parent" edge to the Thread entity.
func (tuo *ThreadUpdateOne) SetParent(t *Thread) *ThreadUpdateOne {
	return tuo.SetParentID(t.ID)
}

// Mutation returns the ThreadMutation object of the builder.
func (tuo *ThreadUpdateOne) Mutation() *ThreadMutation {
	return tuo.mutation
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (tuo *ThreadUpdateOne) ClearCreatedBy() *ThreadUpdateOne {
	tuo.mutation.ClearCreatedBy()
	return tuo
}

// ClearMessages clears all "messages" edges to the Message entity.
func (tuo *ThreadUpdateOne) ClearMessages() *ThreadUpdateOne {
	tuo.mutation.ClearMessages()
	return tuo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (tuo *ThreadUpdateOne) RemoveMessageIDs(ids ...uuid.UUID) *ThreadUpdateOne {
	tuo.mutation.RemoveMessageIDs(ids...)
	return tuo
}

// RemoveMessages removes "messages" edges to Message entities.
func (tuo *ThreadUpdateOne) RemoveMessages(m ...*Message) *ThreadUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tuo.RemoveMessageIDs(ids...)
}

// ClearBookmarks clears all "bookmarks" edges to the Bookmark entity.
func (tuo *ThreadUpdateOne) ClearBookmarks() *ThreadUpdateOne {
	tuo.mutation.ClearBookmarks()
	return tuo
}

// RemoveBookmarkIDs removes the "bookmarks" edge to Bookmark entities by IDs.
func (tuo *ThreadUpdateOne) RemoveBookmarkIDs(ids ...uuid.UUID) *ThreadUpdateOne {
	tuo.mutation.RemoveBookmarkIDs(ids...)
	return tuo
}

// RemoveBookmarks removes "bookmarks" edges to Bookmark entities.
func (tuo *ThreadUpdateOne) RemoveBookmarks(b ...*Bookmark) *ThreadUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tuo.RemoveBookmarkIDs(ids...)
}

// ClearChild clears the "child" edge to the Thread entity.
func (tuo *ThreadUpdateOne) ClearChild() *ThreadUpdateOne {
	tuo.mutation.ClearChild()
	return tuo
}

// ClearParent clears the "parent" edge to the Thread entity.
func (tuo *ThreadUpdateOne) ClearParent() *ThreadUpdateOne {
	tuo.mutation.ClearParent()
	return tuo
}

// Where appends a list predicates to the ThreadUpdate builder.
func (tuo *ThreadUpdateOne) Where(ps ...predicate.Thread) *ThreadUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *ThreadUpdateOne) Select(field string, fields ...string) *ThreadUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Thread entity.
func (tuo *ThreadUpdateOne) Save(ctx context.Context) (*Thread, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *ThreadUpdateOne) SaveX(ctx context.Context) *Thread {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *ThreadUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *ThreadUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *ThreadUpdateOne) sqlSave(ctx context.Context) (_node *Thread, err error) {
	_spec := sqlgraph.NewUpdateSpec(thread.Table, thread.Columns, sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Thread.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, thread.FieldID)
		for _, f := range fields {
			if !thread.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != thread.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(thread.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(thread.FieldName, field.TypeString, value)
	}
	if tuo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.CreatedByTable,
			Columns: []string{thread.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.CreatedByTable,
			Columns: []string{thread.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.MessagesTable,
			Columns: []string{thread.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !tuo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.MessagesTable,
			Columns: []string{thread.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.MessagesTable,
			Columns: []string{thread.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.BookmarksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.BookmarksTable,
			Columns: []string{thread.BookmarksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedBookmarksIDs(); len(nodes) > 0 && !tuo.mutation.BookmarksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.BookmarksTable,
			Columns: []string{thread.BookmarksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.BookmarksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.BookmarksTable,
			Columns: []string{thread.BookmarksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ChildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   thread.ChildTable,
			Columns: []string{thread.ChildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ChildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   thread.ChildTable,
			Columns: []string{thread.ChildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   thread.ParentTable,
			Columns: []string{thread.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   thread.ParentTable,
			Columns: []string{thread.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Thread{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
