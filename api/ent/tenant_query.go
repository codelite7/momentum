// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/codelite7/momentum/api/ent/predicate"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"github.com/codelite7/momentum/api/ent/tenant"
	"github.com/codelite7/momentum/api/ent/user"
)

// TenantQuery is the builder for querying Tenant entities.
type TenantQuery struct {
	config
	ctx            *QueryContext
	order          []tenant.OrderOption
	inters         []Interceptor
	predicates     []predicate.Tenant
	withUsers      *UserQuery
	loadTotal      []func(context.Context, []*Tenant) error
	modifiers      []func(*sql.Selector)
	withNamedUsers map[string]*UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TenantQuery builder.
func (tq *TenantQuery) Where(ps ...predicate.Tenant) *TenantQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TenantQuery) Limit(limit int) *TenantQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TenantQuery) Offset(offset int) *TenantQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TenantQuery) Unique(unique bool) *TenantQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TenantQuery) Order(o ...tenant.OrderOption) *TenantQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryUsers chains the current query on the "users" edge.
func (tq *TenantQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tenant.Table, tenant.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tenant.UsersTable, tenant.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Tenant entity from the query.
// Returns a *NotFoundError when no Tenant was found.
func (tq *TenantQuery) First(ctx context.Context) (*Tenant, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tenant.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TenantQuery) FirstX(ctx context.Context) *Tenant {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Tenant ID from the query.
// Returns a *NotFoundError when no Tenant ID was found.
func (tq *TenantQuery) FirstID(ctx context.Context) (id pulid.ID, err error) {
	var ids []pulid.ID
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tenant.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TenantQuery) FirstIDX(ctx context.Context) pulid.ID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Tenant entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Tenant entity is found.
// Returns a *NotFoundError when no Tenant entities are found.
func (tq *TenantQuery) Only(ctx context.Context) (*Tenant, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tenant.Label}
	default:
		return nil, &NotSingularError{tenant.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TenantQuery) OnlyX(ctx context.Context) *Tenant {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Tenant ID in the query.
// Returns a *NotSingularError when more than one Tenant ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TenantQuery) OnlyID(ctx context.Context) (id pulid.ID, err error) {
	var ids []pulid.ID
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tenant.Label}
	default:
		err = &NotSingularError{tenant.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TenantQuery) OnlyIDX(ctx context.Context) pulid.ID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tenants.
func (tq *TenantQuery) All(ctx context.Context) ([]*Tenant, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Tenant, *TenantQuery]()
	return withInterceptors[[]*Tenant](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TenantQuery) AllX(ctx context.Context) []*Tenant {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Tenant IDs.
func (tq *TenantQuery) IDs(ctx context.Context) (ids []pulid.ID, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(tenant.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TenantQuery) IDsX(ctx context.Context) []pulid.ID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TenantQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TenantQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TenantQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TenantQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TenantQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TenantQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TenantQuery) Clone() *TenantQuery {
	if tq == nil {
		return nil
	}
	return &TenantQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]tenant.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Tenant{}, tq.predicates...),
		withUsers:  tq.withUsers.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TenantQuery) WithUsers(opts ...func(*UserQuery)) *TenantQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUsers = query
	return tq
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
//	client.Tenant.Query().
//		GroupBy(tenant.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TenantQuery) GroupBy(field string, fields ...string) *TenantGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TenantGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = tenant.Label
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
//	client.Tenant.Query().
//		Select(tenant.FieldCreatedAt).
//		Scan(ctx, &v)
func (tq *TenantQuery) Select(fields ...string) *TenantSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TenantSelect{TenantQuery: tq}
	sbuild.label = tenant.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TenantSelect configured with the given aggregations.
func (tq *TenantQuery) Aggregate(fns ...AggregateFunc) *TenantSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TenantQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !tenant.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TenantQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Tenant, error) {
	var (
		nodes       = []*Tenant{}
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Tenant).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Tenant{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withUsers; query != nil {
		if err := tq.loadUsers(ctx, query, nodes,
			func(n *Tenant) { n.Edges.Users = []*User{} },
			func(n *Tenant, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedUsers {
		if err := tq.loadUsers(ctx, query, nodes,
			func(n *Tenant) { n.appendNamedUsers(name) },
			func(n *Tenant, e *User) { n.appendNamedUsers(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range tq.loadTotal {
		if err := tq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TenantQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Tenant, init func(*Tenant), assign func(*Tenant, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[pulid.ID]*Tenant)
	nids := make(map[pulid.ID]map[*Tenant]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(tenant.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(tenant.UsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(tenant.UsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(tenant.UsersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(pulid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*pulid.ID)
				inValue := *values[1].(*pulid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Tenant]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (tq *TenantQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TenantQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tenant.Table, tenant.Columns, sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeString))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tenant.FieldID)
		for i := range fields {
			if fields[i] != tenant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TenantQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(tenant.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = tenant.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range tq.modifiers {
		m(selector)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tq *TenantQuery) ForUpdate(opts ...sql.LockOption) *TenantQuery {
	if tq.driver.Dialect() == dialect.Postgres {
		tq.Unique(false)
	}
	tq.modifiers = append(tq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tq *TenantQuery) ForShare(opts ...sql.LockOption) *TenantQuery {
	if tq.driver.Dialect() == dialect.Postgres {
		tq.Unique(false)
	}
	tq.modifiers = append(tq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tq
}

// WithNamedUsers tells the query-builder to eager-load the nodes that are connected to the "users"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TenantQuery) WithNamedUsers(name string, opts ...func(*UserQuery)) *TenantQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedUsers == nil {
		tq.withNamedUsers = make(map[string]*UserQuery)
	}
	tq.withNamedUsers[name] = query
	return tq
}

// TenantGroupBy is the group-by builder for Tenant entities.
type TenantGroupBy struct {
	selector
	build *TenantQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TenantGroupBy) Aggregate(fns ...AggregateFunc) *TenantGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TenantGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TenantQuery, *TenantGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TenantGroupBy) sqlScan(ctx context.Context, root *TenantQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TenantSelect is the builder for selecting fields of Tenant entities.
type TenantSelect struct {
	*TenantQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TenantSelect) Aggregate(fns ...AggregateFunc) *TenantSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TenantSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TenantQuery, *TenantSelect](ctx, ts.TenantQuery, ts, ts.inters, v)
}

func (ts *TenantSelect) sqlScan(ctx context.Context, root *TenantQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
