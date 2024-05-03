// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

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

// MessageQuery is the builder for querying Message entities.
type MessageQuery struct {
	config
	ctx                *QueryContext
	order              []message.OrderOption
	inters             []Interceptor
	predicates         []predicate.Message
	withSentBy         *UserQuery
	withThread         *ThreadQuery
	withBookmarks      *BookmarkQuery
	withFKs            bool
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*Message) error
	withNamedBookmarks map[string]*BookmarkQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageQuery builder.
func (mq *MessageQuery) Where(ps ...predicate.Message) *MessageQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MessageQuery) Limit(limit int) *MessageQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MessageQuery) Offset(offset int) *MessageQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MessageQuery) Unique(unique bool) *MessageQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MessageQuery) Order(o ...message.OrderOption) *MessageQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QuerySentBy chains the current query on the "sent_by" edge.
func (mq *MessageQuery) QuerySentBy() *UserQuery {
	query := (&UserClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.SentByTable, message.SentByColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryThread chains the current query on the "thread" edge.
func (mq *MessageQuery) QueryThread() *ThreadQuery {
	query := (&ThreadClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, selector),
			sqlgraph.To(thread.Table, thread.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.ThreadTable, message.ThreadColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBookmarks chains the current query on the "bookmarks" edge.
func (mq *MessageQuery) QueryBookmarks() *BookmarkQuery {
	query := (&BookmarkClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, selector),
			sqlgraph.To(bookmark.Table, bookmark.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, message.BookmarksTable, message.BookmarksColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Message entity from the query.
// Returns a *NotFoundError when no Message was found.
func (mq *MessageQuery) First(ctx context.Context) (*Message, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{message.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MessageQuery) FirstX(ctx context.Context) *Message {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Message ID from the query.
// Returns a *NotFoundError when no Message ID was found.
func (mq *MessageQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{message.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MessageQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Message entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Message entity is found.
// Returns a *NotFoundError when no Message entities are found.
func (mq *MessageQuery) Only(ctx context.Context) (*Message, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{message.Label}
	default:
		return nil, &NotSingularError{message.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MessageQuery) OnlyX(ctx context.Context) *Message {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Message ID in the query.
// Returns a *NotSingularError when more than one Message ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MessageQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{message.Label}
	default:
		err = &NotSingularError{message.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MessageQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Messages.
func (mq *MessageQuery) All(ctx context.Context) ([]*Message, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Message, *MessageQuery]()
	return withInterceptors[[]*Message](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MessageQuery) AllX(ctx context.Context) []*Message {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Message IDs.
func (mq *MessageQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(message.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MessageQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MessageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MessageQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MessageQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MessageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MessageQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MessageQuery) Clone() *MessageQuery {
	if mq == nil {
		return nil
	}
	return &MessageQuery{
		config:        mq.config,
		ctx:           mq.ctx.Clone(),
		order:         append([]message.OrderOption{}, mq.order...),
		inters:        append([]Interceptor{}, mq.inters...),
		predicates:    append([]predicate.Message{}, mq.predicates...),
		withSentBy:    mq.withSentBy.Clone(),
		withThread:    mq.withThread.Clone(),
		withBookmarks: mq.withBookmarks.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithSentBy tells the query-builder to eager-load the nodes that are connected to
// the "sent_by" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MessageQuery) WithSentBy(opts ...func(*UserQuery)) *MessageQuery {
	query := (&UserClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withSentBy = query
	return mq
}

// WithThread tells the query-builder to eager-load the nodes that are connected to
// the "thread" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MessageQuery) WithThread(opts ...func(*ThreadQuery)) *MessageQuery {
	query := (&ThreadClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withThread = query
	return mq
}

// WithBookmarks tells the query-builder to eager-load the nodes that are connected to
// the "bookmarks" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MessageQuery) WithBookmarks(opts ...func(*BookmarkQuery)) *MessageQuery {
	query := (&BookmarkClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withBookmarks = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Message.Query().
//		GroupBy(message.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MessageQuery) GroupBy(field string, fields ...string) *MessageGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MessageGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = message.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Message.Query().
//		Select(message.FieldCreatedAt).
//		Scan(ctx, &v)
func (mq *MessageQuery) Select(fields ...string) *MessageSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MessageSelect{MessageQuery: mq}
	sbuild.label = message.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MessageSelect configured with the given aggregations.
func (mq *MessageQuery) Aggregate(fns ...AggregateFunc) *MessageSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MessageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !message.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Message, error) {
	var (
		nodes       = []*Message{}
		withFKs     = mq.withFKs
		_spec       = mq.querySpec()
		loadedTypes = [3]bool{
			mq.withSentBy != nil,
			mq.withThread != nil,
			mq.withBookmarks != nil,
		}
	)
	if mq.withSentBy != nil || mq.withThread != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, message.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Message).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Message{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withSentBy; query != nil {
		if err := mq.loadSentBy(ctx, query, nodes, nil,
			func(n *Message, e *User) { n.Edges.SentBy = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withThread; query != nil {
		if err := mq.loadThread(ctx, query, nodes, nil,
			func(n *Message, e *Thread) { n.Edges.Thread = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withBookmarks; query != nil {
		if err := mq.loadBookmarks(ctx, query, nodes,
			func(n *Message) { n.Edges.Bookmarks = []*Bookmark{} },
			func(n *Message, e *Bookmark) { n.Edges.Bookmarks = append(n.Edges.Bookmarks, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range mq.withNamedBookmarks {
		if err := mq.loadBookmarks(ctx, query, nodes,
			func(n *Message) { n.appendNamedBookmarks(name) },
			func(n *Message, e *Bookmark) { n.appendNamedBookmarks(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range mq.loadTotal {
		if err := mq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MessageQuery) loadSentBy(ctx context.Context, query *UserQuery, nodes []*Message, init func(*Message), assign func(*Message, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Message)
	for i := range nodes {
		if nodes[i].user_messages == nil {
			continue
		}
		fk := *nodes[i].user_messages
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_messages" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MessageQuery) loadThread(ctx context.Context, query *ThreadQuery, nodes []*Message, init func(*Message), assign func(*Message, *Thread)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Message)
	for i := range nodes {
		if nodes[i].thread_messages == nil {
			continue
		}
		fk := *nodes[i].thread_messages
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(thread.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "thread_messages" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MessageQuery) loadBookmarks(ctx context.Context, query *BookmarkQuery, nodes []*Message, init func(*Message), assign func(*Message, *Bookmark)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Message)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Bookmark(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(message.BookmarksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.message_bookmarks
		if fk == nil {
			return fmt.Errorf(`foreign-key "message_bookmarks" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "message_bookmarks" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for i := range fields {
			if fields[i] != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(message.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = message.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedBookmarks tells the query-builder to eager-load the nodes that are connected to the "bookmarks"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (mq *MessageQuery) WithNamedBookmarks(name string, opts ...func(*BookmarkQuery)) *MessageQuery {
	query := (&BookmarkClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if mq.withNamedBookmarks == nil {
		mq.withNamedBookmarks = make(map[string]*BookmarkQuery)
	}
	mq.withNamedBookmarks[name] = query
	return mq
}

// MessageGroupBy is the group-by builder for Message entities.
type MessageGroupBy struct {
	selector
	build *MessageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MessageGroupBy) Aggregate(fns ...AggregateFunc) *MessageGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MessageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MessageQuery, *MessageGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MessageGroupBy) sqlScan(ctx context.Context, root *MessageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MessageSelect is the builder for selecting fields of Message entities.
type MessageSelect struct {
	*MessageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MessageSelect) Aggregate(fns ...AggregateFunc) *MessageSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MessageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MessageQuery, *MessageSelect](ctx, ms.MessageQuery, ms, ms.inters, v)
}

func (ms *MessageSelect) sqlScan(ctx context.Context, root *MessageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
