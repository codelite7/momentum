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
	"github.com/codelite7/momentum/api/ent/response"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/ent/user"
	"github.com/google/uuid"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MessageUpdate) SetUpdatedAt(t time.Time) *MessageUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableUpdatedAt(t *time.Time) *MessageUpdate {
	if t != nil {
		mu.SetUpdatedAt(*t)
	}
	return mu
}

// SetContent sets the "content" field.
func (mu *MessageUpdate) SetContent(s string) *MessageUpdate {
	mu.mutation.SetContent(s)
	return mu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableContent(s *string) *MessageUpdate {
	if s != nil {
		mu.SetContent(*s)
	}
	return mu
}

// SetSentByID sets the "sent_by" edge to the User entity by ID.
func (mu *MessageUpdate) SetSentByID(id uuid.UUID) *MessageUpdate {
	mu.mutation.SetSentByID(id)
	return mu
}

// SetSentBy sets the "sent_by" edge to the User entity.
func (mu *MessageUpdate) SetSentBy(u *User) *MessageUpdate {
	return mu.SetSentByID(u.ID)
}

// SetThreadID sets the "thread" edge to the Thread entity by ID.
func (mu *MessageUpdate) SetThreadID(id uuid.UUID) *MessageUpdate {
	mu.mutation.SetThreadID(id)
	return mu
}

// SetThread sets the "thread" edge to the Thread entity.
func (mu *MessageUpdate) SetThread(t *Thread) *MessageUpdate {
	return mu.SetThreadID(t.ID)
}

// AddBookmarkIDs adds the "bookmarks" edge to the Bookmark entity by IDs.
func (mu *MessageUpdate) AddBookmarkIDs(ids ...uuid.UUID) *MessageUpdate {
	mu.mutation.AddBookmarkIDs(ids...)
	return mu
}

// AddBookmarks adds the "bookmarks" edges to the Bookmark entity.
func (mu *MessageUpdate) AddBookmarks(b ...*Bookmark) *MessageUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return mu.AddBookmarkIDs(ids...)
}

// SetResponseID sets the "response" edge to the Response entity by ID.
func (mu *MessageUpdate) SetResponseID(id uuid.UUID) *MessageUpdate {
	mu.mutation.SetResponseID(id)
	return mu
}

// SetNillableResponseID sets the "response" edge to the Response entity by ID if the given value is not nil.
func (mu *MessageUpdate) SetNillableResponseID(id *uuid.UUID) *MessageUpdate {
	if id != nil {
		mu = mu.SetResponseID(*id)
	}
	return mu
}

// SetResponse sets the "response" edge to the Response entity.
func (mu *MessageUpdate) SetResponse(r *Response) *MessageUpdate {
	return mu.SetResponseID(r.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearSentBy clears the "sent_by" edge to the User entity.
func (mu *MessageUpdate) ClearSentBy() *MessageUpdate {
	mu.mutation.ClearSentBy()
	return mu
}

// ClearThread clears the "thread" edge to the Thread entity.
func (mu *MessageUpdate) ClearThread() *MessageUpdate {
	mu.mutation.ClearThread()
	return mu
}

// ClearBookmarks clears all "bookmarks" edges to the Bookmark entity.
func (mu *MessageUpdate) ClearBookmarks() *MessageUpdate {
	mu.mutation.ClearBookmarks()
	return mu
}

// RemoveBookmarkIDs removes the "bookmarks" edge to Bookmark entities by IDs.
func (mu *MessageUpdate) RemoveBookmarkIDs(ids ...uuid.UUID) *MessageUpdate {
	mu.mutation.RemoveBookmarkIDs(ids...)
	return mu
}

// RemoveBookmarks removes "bookmarks" edges to Bookmark entities.
func (mu *MessageUpdate) RemoveBookmarks(b ...*Bookmark) *MessageUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return mu.RemoveBookmarkIDs(ids...)
}

// ClearResponse clears the "response" edge to the Response entity.
func (mu *MessageUpdate) ClearResponse() *MessageUpdate {
	mu.mutation.ClearResponse()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MessageUpdate) check() error {
	if _, ok := mu.mutation.SentByID(); mu.mutation.SentByCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Message.sent_by"`)
	}
	if _, ok := mu.mutation.ThreadID(); mu.mutation.ThreadCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Message.thread"`)
	}
	return nil
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(message.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
	}
	if mu.mutation.SentByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.SentByTable,
			Columns: []string{message.SentByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.SentByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.SentByTable,
			Columns: []string{message.SentByColumn},
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
	if mu.mutation.ThreadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ThreadTable,
			Columns: []string{message.ThreadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ThreadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ThreadTable,
			Columns: []string{message.ThreadColumn},
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
	if mu.mutation.BookmarksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.BookmarksTable,
			Columns: []string{message.BookmarksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedBookmarksIDs(); len(nodes) > 0 && !mu.mutation.BookmarksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.BookmarksTable,
			Columns: []string{message.BookmarksColumn},
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
	if nodes := mu.mutation.BookmarksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.BookmarksTable,
			Columns: []string{message.BookmarksColumn},
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
	if mu.mutation.ResponseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   message.ResponseTable,
			Columns: []string{message.ResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ResponseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   message.ResponseTable,
			Columns: []string{message.ResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MessageUpdateOne) SetUpdatedAt(t time.Time) *MessageUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableUpdatedAt(t *time.Time) *MessageUpdateOne {
	if t != nil {
		muo.SetUpdatedAt(*t)
	}
	return muo
}

// SetContent sets the "content" field.
func (muo *MessageUpdateOne) SetContent(s string) *MessageUpdateOne {
	muo.mutation.SetContent(s)
	return muo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableContent(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetContent(*s)
	}
	return muo
}

// SetSentByID sets the "sent_by" edge to the User entity by ID.
func (muo *MessageUpdateOne) SetSentByID(id uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetSentByID(id)
	return muo
}

// SetSentBy sets the "sent_by" edge to the User entity.
func (muo *MessageUpdateOne) SetSentBy(u *User) *MessageUpdateOne {
	return muo.SetSentByID(u.ID)
}

// SetThreadID sets the "thread" edge to the Thread entity by ID.
func (muo *MessageUpdateOne) SetThreadID(id uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetThreadID(id)
	return muo
}

// SetThread sets the "thread" edge to the Thread entity.
func (muo *MessageUpdateOne) SetThread(t *Thread) *MessageUpdateOne {
	return muo.SetThreadID(t.ID)
}

// AddBookmarkIDs adds the "bookmarks" edge to the Bookmark entity by IDs.
func (muo *MessageUpdateOne) AddBookmarkIDs(ids ...uuid.UUID) *MessageUpdateOne {
	muo.mutation.AddBookmarkIDs(ids...)
	return muo
}

// AddBookmarks adds the "bookmarks" edges to the Bookmark entity.
func (muo *MessageUpdateOne) AddBookmarks(b ...*Bookmark) *MessageUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return muo.AddBookmarkIDs(ids...)
}

// SetResponseID sets the "response" edge to the Response entity by ID.
func (muo *MessageUpdateOne) SetResponseID(id uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetResponseID(id)
	return muo
}

// SetNillableResponseID sets the "response" edge to the Response entity by ID if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableResponseID(id *uuid.UUID) *MessageUpdateOne {
	if id != nil {
		muo = muo.SetResponseID(*id)
	}
	return muo
}

// SetResponse sets the "response" edge to the Response entity.
func (muo *MessageUpdateOne) SetResponse(r *Response) *MessageUpdateOne {
	return muo.SetResponseID(r.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearSentBy clears the "sent_by" edge to the User entity.
func (muo *MessageUpdateOne) ClearSentBy() *MessageUpdateOne {
	muo.mutation.ClearSentBy()
	return muo
}

// ClearThread clears the "thread" edge to the Thread entity.
func (muo *MessageUpdateOne) ClearThread() *MessageUpdateOne {
	muo.mutation.ClearThread()
	return muo
}

// ClearBookmarks clears all "bookmarks" edges to the Bookmark entity.
func (muo *MessageUpdateOne) ClearBookmarks() *MessageUpdateOne {
	muo.mutation.ClearBookmarks()
	return muo
}

// RemoveBookmarkIDs removes the "bookmarks" edge to Bookmark entities by IDs.
func (muo *MessageUpdateOne) RemoveBookmarkIDs(ids ...uuid.UUID) *MessageUpdateOne {
	muo.mutation.RemoveBookmarkIDs(ids...)
	return muo
}

// RemoveBookmarks removes "bookmarks" edges to Bookmark entities.
func (muo *MessageUpdateOne) RemoveBookmarks(b ...*Bookmark) *MessageUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return muo.RemoveBookmarkIDs(ids...)
}

// ClearResponse clears the "response" edge to the Response entity.
func (muo *MessageUpdateOne) ClearResponse() *MessageUpdateOne {
	muo.mutation.ClearResponse()
	return muo
}

// Where appends a list predicates to the MessageUpdate builder.
func (muo *MessageUpdateOne) Where(ps ...predicate.Message) *MessageUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MessageUpdateOne) check() error {
	if _, ok := muo.mutation.SentByID(); muo.mutation.SentByCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Message.sent_by"`)
	}
	if _, ok := muo.mutation.ThreadID(); muo.mutation.ThreadCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Message.thread"`)
	}
	return nil
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(message.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
	}
	if muo.mutation.SentByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.SentByTable,
			Columns: []string{message.SentByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.SentByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.SentByTable,
			Columns: []string{message.SentByColumn},
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
	if muo.mutation.ThreadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ThreadTable,
			Columns: []string{message.ThreadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ThreadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ThreadTable,
			Columns: []string{message.ThreadColumn},
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
	if muo.mutation.BookmarksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.BookmarksTable,
			Columns: []string{message.BookmarksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedBookmarksIDs(); len(nodes) > 0 && !muo.mutation.BookmarksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.BookmarksTable,
			Columns: []string{message.BookmarksColumn},
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
	if nodes := muo.mutation.BookmarksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.BookmarksTable,
			Columns: []string{message.BookmarksColumn},
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
	if muo.mutation.ResponseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   message.ResponseTable,
			Columns: []string{message.ResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ResponseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   message.ResponseTable,
			Columns: []string{message.ResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
