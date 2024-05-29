// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/codelite7/momentum/api/ent/bookmark"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/predicate"
	"github.com/codelite7/momentum/api/ent/response"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/ent/user"
)

// BookmarkQuery is the builder for querying Bookmark entities.
type BookmarkQuery struct {
	config
	ctx          *QueryContext
	order        []bookmark.OrderOption
	inters       []Interceptor
	predicates   []predicate.Bookmark
	withUser     *UserQuery
	withThread   *ThreadQuery
	withMessage  *MessageQuery
	withResponse *ResponseQuery
	withFKs      bool
	modifiers    []func(*sql.Selector)
	loadTotal    []func(context.Context, []*Bookmark) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BookmarkQuery builder.
func (bq *BookmarkQuery) Where(ps ...predicate.Bookmark) *BookmarkQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BookmarkQuery) Limit(limit int) *BookmarkQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BookmarkQuery) Offset(offset int) *BookmarkQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BookmarkQuery) Unique(unique bool) *BookmarkQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BookmarkQuery) Order(o ...bookmark.OrderOption) *BookmarkQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryUser chains the current query on the "user" edge.
func (bq *BookmarkQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bookmark.Table, bookmark.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bookmark.UserTable, bookmark.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryThread chains the current query on the "thread" edge.
func (bq *BookmarkQuery) QueryThread() *ThreadQuery {
	query := (&ThreadClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bookmark.Table, bookmark.FieldID, selector),
			sqlgraph.To(thread.Table, thread.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bookmark.ThreadTable, bookmark.ThreadColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMessage chains the current query on the "message" edge.
func (bq *BookmarkQuery) QueryMessage() *MessageQuery {
	query := (&MessageClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bookmark.Table, bookmark.FieldID, selector),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bookmark.MessageTable, bookmark.MessageColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryResponse chains the current query on the "response" edge.
func (bq *BookmarkQuery) QueryResponse() *ResponseQuery {
	query := (&ResponseClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bookmark.Table, bookmark.FieldID, selector),
			sqlgraph.To(response.Table, response.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bookmark.ResponseTable, bookmark.ResponseColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Bookmark entity from the query.
// Returns a *NotFoundError when no Bookmark was found.
func (bq *BookmarkQuery) First(ctx context.Context) (*Bookmark, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bookmark.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BookmarkQuery) FirstX(ctx context.Context) *Bookmark {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Bookmark ID from the query.
// Returns a *NotFoundError when no Bookmark ID was found.
func (bq *BookmarkQuery) FirstID(ctx context.Context) (id pulid.ID, err error) {
	var ids []pulid.ID
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bookmark.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BookmarkQuery) FirstIDX(ctx context.Context) pulid.ID {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Bookmark entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Bookmark entity is found.
// Returns a *NotFoundError when no Bookmark entities are found.
func (bq *BookmarkQuery) Only(ctx context.Context) (*Bookmark, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bookmark.Label}
	default:
		return nil, &NotSingularError{bookmark.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BookmarkQuery) OnlyX(ctx context.Context) *Bookmark {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Bookmark ID in the query.
// Returns a *NotSingularError when more than one Bookmark ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BookmarkQuery) OnlyID(ctx context.Context) (id pulid.ID, err error) {
	var ids []pulid.ID
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bookmark.Label}
	default:
		err = &NotSingularError{bookmark.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BookmarkQuery) OnlyIDX(ctx context.Context) pulid.ID {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bookmarks.
func (bq *BookmarkQuery) All(ctx context.Context) ([]*Bookmark, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Bookmark, *BookmarkQuery]()
	return withInterceptors[[]*Bookmark](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BookmarkQuery) AllX(ctx context.Context) []*Bookmark {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Bookmark IDs.
func (bq *BookmarkQuery) IDs(ctx context.Context) (ids []pulid.ID, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(bookmark.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BookmarkQuery) IDsX(ctx context.Context) []pulid.ID {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BookmarkQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BookmarkQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BookmarkQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BookmarkQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BookmarkQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BookmarkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BookmarkQuery) Clone() *BookmarkQuery {
	if bq == nil {
		return nil
	}
	return &BookmarkQuery{
		config:       bq.config,
		ctx:          bq.ctx.Clone(),
		order:        append([]bookmark.OrderOption{}, bq.order...),
		inters:       append([]Interceptor{}, bq.inters...),
		predicates:   append([]predicate.Bookmark{}, bq.predicates...),
		withUser:     bq.withUser.Clone(),
		withThread:   bq.withThread.Clone(),
		withMessage:  bq.withMessage.Clone(),
		withResponse: bq.withResponse.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BookmarkQuery) WithUser(opts ...func(*UserQuery)) *BookmarkQuery {
	query := (&UserClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withUser = query
	return bq
}

// WithThread tells the query-builder to eager-load the nodes that are connected to
// the "thread" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BookmarkQuery) WithThread(opts ...func(*ThreadQuery)) *BookmarkQuery {
	query := (&ThreadClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withThread = query
	return bq
}

// WithMessage tells the query-builder to eager-load the nodes that are connected to
// the "message" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BookmarkQuery) WithMessage(opts ...func(*MessageQuery)) *BookmarkQuery {
	query := (&MessageClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withMessage = query
	return bq
}

// WithResponse tells the query-builder to eager-load the nodes that are connected to
// the "response" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BookmarkQuery) WithResponse(opts ...func(*ResponseQuery)) *BookmarkQuery {
	query := (&ResponseClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withResponse = query
	return bq
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
//	client.Bookmark.Query().
//		GroupBy(bookmark.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BookmarkQuery) GroupBy(field string, fields ...string) *BookmarkGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BookmarkGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = bookmark.Label
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
//	client.Bookmark.Query().
//		Select(bookmark.FieldCreatedAt).
//		Scan(ctx, &v)
func (bq *BookmarkQuery) Select(fields ...string) *BookmarkSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BookmarkSelect{BookmarkQuery: bq}
	sbuild.label = bookmark.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BookmarkSelect configured with the given aggregations.
func (bq *BookmarkQuery) Aggregate(fns ...AggregateFunc) *BookmarkSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BookmarkQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !bookmark.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BookmarkQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Bookmark, error) {
	var (
		nodes       = []*Bookmark{}
		withFKs     = bq.withFKs
		_spec       = bq.querySpec()
		loadedTypes = [4]bool{
			bq.withUser != nil,
			bq.withThread != nil,
			bq.withMessage != nil,
			bq.withResponse != nil,
		}
	)
	if bq.withUser != nil || bq.withThread != nil || bq.withMessage != nil || bq.withResponse != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, bookmark.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Bookmark).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Bookmark{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withUser; query != nil {
		if err := bq.loadUser(ctx, query, nodes, nil,
			func(n *Bookmark, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := bq.withThread; query != nil {
		if err := bq.loadThread(ctx, query, nodes, nil,
			func(n *Bookmark, e *Thread) { n.Edges.Thread = e }); err != nil {
			return nil, err
		}
	}
	if query := bq.withMessage; query != nil {
		if err := bq.loadMessage(ctx, query, nodes, nil,
			func(n *Bookmark, e *Message) { n.Edges.Message = e }); err != nil {
			return nil, err
		}
	}
	if query := bq.withResponse; query != nil {
		if err := bq.loadResponse(ctx, query, nodes, nil,
			func(n *Bookmark, e *Response) { n.Edges.Response = e }); err != nil {
			return nil, err
		}
	}
	for i := range bq.loadTotal {
		if err := bq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BookmarkQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Bookmark, init func(*Bookmark), assign func(*Bookmark, *User)) error {
	ids := make([]pulid.ID, 0, len(nodes))
	nodeids := make(map[pulid.ID][]*Bookmark)
	for i := range nodes {
		if nodes[i].user_bookmarks == nil {
			continue
		}
		fk := *nodes[i].user_bookmarks
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
			return fmt.Errorf(`unexpected foreign-key "user_bookmarks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bq *BookmarkQuery) loadThread(ctx context.Context, query *ThreadQuery, nodes []*Bookmark, init func(*Bookmark), assign func(*Bookmark, *Thread)) error {
	ids := make([]pulid.ID, 0, len(nodes))
	nodeids := make(map[pulid.ID][]*Bookmark)
	for i := range nodes {
		if nodes[i].thread_bookmarks == nil {
			continue
		}
		fk := *nodes[i].thread_bookmarks
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
			return fmt.Errorf(`unexpected foreign-key "thread_bookmarks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bq *BookmarkQuery) loadMessage(ctx context.Context, query *MessageQuery, nodes []*Bookmark, init func(*Bookmark), assign func(*Bookmark, *Message)) error {
	ids := make([]pulid.ID, 0, len(nodes))
	nodeids := make(map[pulid.ID][]*Bookmark)
	for i := range nodes {
		if nodes[i].message_bookmarks == nil {
			continue
		}
		fk := *nodes[i].message_bookmarks
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(message.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "message_bookmarks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bq *BookmarkQuery) loadResponse(ctx context.Context, query *ResponseQuery, nodes []*Bookmark, init func(*Bookmark), assign func(*Bookmark, *Response)) error {
	ids := make([]pulid.ID, 0, len(nodes))
	nodeids := make(map[pulid.ID][]*Bookmark)
	for i := range nodes {
		if nodes[i].response_bookmarks == nil {
			continue
		}
		fk := *nodes[i].response_bookmarks
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(response.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "response_bookmarks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (bq *BookmarkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BookmarkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bookmark.Table, bookmark.Columns, sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeString))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookmark.FieldID)
		for i := range fields {
			if fields[i] != bookmark.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BookmarkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(bookmark.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = bookmark.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BookmarkGroupBy is the group-by builder for Bookmark entities.
type BookmarkGroupBy struct {
	selector
	build *BookmarkQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BookmarkGroupBy) Aggregate(fns ...AggregateFunc) *BookmarkGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BookmarkGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookmarkQuery, *BookmarkGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BookmarkGroupBy) sqlScan(ctx context.Context, root *BookmarkQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BookmarkSelect is the builder for selecting fields of Bookmark entities.
type BookmarkSelect struct {
	*BookmarkQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BookmarkSelect) Aggregate(fns ...AggregateFunc) *BookmarkSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BookmarkSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookmarkQuery, *BookmarkSelect](ctx, bs.BookmarkQuery, bs, bs.inters, v)
}

func (bs *BookmarkSelect) sqlScan(ctx context.Context, root *BookmarkQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
