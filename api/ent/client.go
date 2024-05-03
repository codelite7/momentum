// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/codelite7/momentum/api/ent/migrate"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/codelite7/momentum/api/ent/agent"
	"github.com/codelite7/momentum/api/ent/bookmark"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Agent is the client for interacting with the Agent builders.
	Agent *AgentClient
	// Bookmark is the client for interacting with the Bookmark builders.
	Bookmark *BookmarkClient
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
	// Thread is the client for interacting with the Thread builders.
	Thread *ThreadClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Agent = NewAgentClient(c.config)
	c.Bookmark = NewBookmarkClient(c.config)
	c.Message = NewMessageClient(c.config)
	c.Thread = NewThreadClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Agent:    NewAgentClient(cfg),
		Bookmark: NewBookmarkClient(cfg),
		Message:  NewMessageClient(cfg),
		Thread:   NewThreadClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Agent:    NewAgentClient(cfg),
		Bookmark: NewBookmarkClient(cfg),
		Message:  NewMessageClient(cfg),
		Thread:   NewThreadClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Agent.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Agent.Use(hooks...)
	c.Bookmark.Use(hooks...)
	c.Message.Use(hooks...)
	c.Thread.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Agent.Intercept(interceptors...)
	c.Bookmark.Intercept(interceptors...)
	c.Message.Intercept(interceptors...)
	c.Thread.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AgentMutation:
		return c.Agent.mutate(ctx, m)
	case *BookmarkMutation:
		return c.Bookmark.mutate(ctx, m)
	case *MessageMutation:
		return c.Message.mutate(ctx, m)
	case *ThreadMutation:
		return c.Thread.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AgentClient is a client for the Agent schema.
type AgentClient struct {
	config
}

// NewAgentClient returns a client for the Agent from the given config.
func NewAgentClient(c config) *AgentClient {
	return &AgentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `agent.Hooks(f(g(h())))`.
func (c *AgentClient) Use(hooks ...Hook) {
	c.hooks.Agent = append(c.hooks.Agent, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `agent.Intercept(f(g(h())))`.
func (c *AgentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Agent = append(c.inters.Agent, interceptors...)
}

// Create returns a builder for creating a Agent entity.
func (c *AgentClient) Create() *AgentCreate {
	mutation := newAgentMutation(c.config, OpCreate)
	return &AgentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Agent entities.
func (c *AgentClient) CreateBulk(builders ...*AgentCreate) *AgentCreateBulk {
	return &AgentCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AgentClient) MapCreateBulk(slice any, setFunc func(*AgentCreate, int)) *AgentCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AgentCreateBulk{err: fmt.Errorf("calling to AgentClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AgentCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AgentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Agent.
func (c *AgentClient) Update() *AgentUpdate {
	mutation := newAgentMutation(c.config, OpUpdate)
	return &AgentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AgentClient) UpdateOne(a *Agent) *AgentUpdateOne {
	mutation := newAgentMutation(c.config, OpUpdateOne, withAgent(a))
	return &AgentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AgentClient) UpdateOneID(id uuid.UUID) *AgentUpdateOne {
	mutation := newAgentMutation(c.config, OpUpdateOne, withAgentID(id))
	return &AgentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Agent.
func (c *AgentClient) Delete() *AgentDelete {
	mutation := newAgentMutation(c.config, OpDelete)
	return &AgentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AgentClient) DeleteOne(a *Agent) *AgentDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AgentClient) DeleteOneID(id uuid.UUID) *AgentDeleteOne {
	builder := c.Delete().Where(agent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AgentDeleteOne{builder}
}

// Query returns a query builder for Agent.
func (c *AgentClient) Query() *AgentQuery {
	return &AgentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAgent},
		inters: c.Interceptors(),
	}
}

// Get returns a Agent entity by its id.
func (c *AgentClient) Get(ctx context.Context, id uuid.UUID) (*Agent, error) {
	return c.Query().Where(agent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AgentClient) GetX(ctx context.Context, id uuid.UUID) *Agent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AgentClient) Hooks() []Hook {
	return c.hooks.Agent
}

// Interceptors returns the client interceptors.
func (c *AgentClient) Interceptors() []Interceptor {
	return c.inters.Agent
}

func (c *AgentClient) mutate(ctx context.Context, m *AgentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AgentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AgentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AgentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AgentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Agent mutation op: %q", m.Op())
	}
}

// BookmarkClient is a client for the Bookmark schema.
type BookmarkClient struct {
	config
}

// NewBookmarkClient returns a client for the Bookmark from the given config.
func NewBookmarkClient(c config) *BookmarkClient {
	return &BookmarkClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bookmark.Hooks(f(g(h())))`.
func (c *BookmarkClient) Use(hooks ...Hook) {
	c.hooks.Bookmark = append(c.hooks.Bookmark, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `bookmark.Intercept(f(g(h())))`.
func (c *BookmarkClient) Intercept(interceptors ...Interceptor) {
	c.inters.Bookmark = append(c.inters.Bookmark, interceptors...)
}

// Create returns a builder for creating a Bookmark entity.
func (c *BookmarkClient) Create() *BookmarkCreate {
	mutation := newBookmarkMutation(c.config, OpCreate)
	return &BookmarkCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Bookmark entities.
func (c *BookmarkClient) CreateBulk(builders ...*BookmarkCreate) *BookmarkCreateBulk {
	return &BookmarkCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *BookmarkClient) MapCreateBulk(slice any, setFunc func(*BookmarkCreate, int)) *BookmarkCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &BookmarkCreateBulk{err: fmt.Errorf("calling to BookmarkClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*BookmarkCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &BookmarkCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Bookmark.
func (c *BookmarkClient) Update() *BookmarkUpdate {
	mutation := newBookmarkMutation(c.config, OpUpdate)
	return &BookmarkUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BookmarkClient) UpdateOne(b *Bookmark) *BookmarkUpdateOne {
	mutation := newBookmarkMutation(c.config, OpUpdateOne, withBookmark(b))
	return &BookmarkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BookmarkClient) UpdateOneID(id uuid.UUID) *BookmarkUpdateOne {
	mutation := newBookmarkMutation(c.config, OpUpdateOne, withBookmarkID(id))
	return &BookmarkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Bookmark.
func (c *BookmarkClient) Delete() *BookmarkDelete {
	mutation := newBookmarkMutation(c.config, OpDelete)
	return &BookmarkDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BookmarkClient) DeleteOne(b *Bookmark) *BookmarkDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BookmarkClient) DeleteOneID(id uuid.UUID) *BookmarkDeleteOne {
	builder := c.Delete().Where(bookmark.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BookmarkDeleteOne{builder}
}

// Query returns a query builder for Bookmark.
func (c *BookmarkClient) Query() *BookmarkQuery {
	return &BookmarkQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBookmark},
		inters: c.Interceptors(),
	}
}

// Get returns a Bookmark entity by its id.
func (c *BookmarkClient) Get(ctx context.Context, id uuid.UUID) (*Bookmark, error) {
	return c.Query().Where(bookmark.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BookmarkClient) GetX(ctx context.Context, id uuid.UUID) *Bookmark {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BookmarkClient) Hooks() []Hook {
	return c.hooks.Bookmark
}

// Interceptors returns the client interceptors.
func (c *BookmarkClient) Interceptors() []Interceptor {
	return c.inters.Bookmark
}

func (c *BookmarkClient) mutate(ctx context.Context, m *BookmarkMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BookmarkCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BookmarkUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BookmarkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BookmarkDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Bookmark mutation op: %q", m.Op())
	}
}

// MessageClient is a client for the Message schema.
type MessageClient struct {
	config
}

// NewMessageClient returns a client for the Message from the given config.
func NewMessageClient(c config) *MessageClient {
	return &MessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `message.Hooks(f(g(h())))`.
func (c *MessageClient) Use(hooks ...Hook) {
	c.hooks.Message = append(c.hooks.Message, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `message.Intercept(f(g(h())))`.
func (c *MessageClient) Intercept(interceptors ...Interceptor) {
	c.inters.Message = append(c.inters.Message, interceptors...)
}

// Create returns a builder for creating a Message entity.
func (c *MessageClient) Create() *MessageCreate {
	mutation := newMessageMutation(c.config, OpCreate)
	return &MessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Message entities.
func (c *MessageClient) CreateBulk(builders ...*MessageCreate) *MessageCreateBulk {
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MessageClient) MapCreateBulk(slice any, setFunc func(*MessageCreate, int)) *MessageCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MessageCreateBulk{err: fmt.Errorf("calling to MessageClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MessageCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Message.
func (c *MessageClient) Update() *MessageUpdate {
	mutation := newMessageMutation(c.config, OpUpdate)
	return &MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageClient) UpdateOne(m *Message) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessage(m))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageClient) UpdateOneID(id uuid.UUID) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessageID(id))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Message.
func (c *MessageClient) Delete() *MessageDelete {
	mutation := newMessageMutation(c.config, OpDelete)
	return &MessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MessageClient) DeleteOne(m *Message) *MessageDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MessageClient) DeleteOneID(id uuid.UUID) *MessageDeleteOne {
	builder := c.Delete().Where(message.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageDeleteOne{builder}
}

// Query returns a query builder for Message.
func (c *MessageClient) Query() *MessageQuery {
	return &MessageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMessage},
		inters: c.Interceptors(),
	}
}

// Get returns a Message entity by its id.
func (c *MessageClient) Get(ctx context.Context, id uuid.UUID) (*Message, error) {
	return c.Query().Where(message.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageClient) GetX(ctx context.Context, id uuid.UUID) *Message {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MessageClient) Hooks() []Hook {
	return c.hooks.Message
}

// Interceptors returns the client interceptors.
func (c *MessageClient) Interceptors() []Interceptor {
	return c.inters.Message
}

func (c *MessageClient) mutate(ctx context.Context, m *MessageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MessageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MessageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Message mutation op: %q", m.Op())
	}
}

// ThreadClient is a client for the Thread schema.
type ThreadClient struct {
	config
}

// NewThreadClient returns a client for the Thread from the given config.
func NewThreadClient(c config) *ThreadClient {
	return &ThreadClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `thread.Hooks(f(g(h())))`.
func (c *ThreadClient) Use(hooks ...Hook) {
	c.hooks.Thread = append(c.hooks.Thread, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `thread.Intercept(f(g(h())))`.
func (c *ThreadClient) Intercept(interceptors ...Interceptor) {
	c.inters.Thread = append(c.inters.Thread, interceptors...)
}

// Create returns a builder for creating a Thread entity.
func (c *ThreadClient) Create() *ThreadCreate {
	mutation := newThreadMutation(c.config, OpCreate)
	return &ThreadCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Thread entities.
func (c *ThreadClient) CreateBulk(builders ...*ThreadCreate) *ThreadCreateBulk {
	return &ThreadCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ThreadClient) MapCreateBulk(slice any, setFunc func(*ThreadCreate, int)) *ThreadCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ThreadCreateBulk{err: fmt.Errorf("calling to ThreadClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ThreadCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ThreadCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Thread.
func (c *ThreadClient) Update() *ThreadUpdate {
	mutation := newThreadMutation(c.config, OpUpdate)
	return &ThreadUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ThreadClient) UpdateOne(t *Thread) *ThreadUpdateOne {
	mutation := newThreadMutation(c.config, OpUpdateOne, withThread(t))
	return &ThreadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ThreadClient) UpdateOneID(id uuid.UUID) *ThreadUpdateOne {
	mutation := newThreadMutation(c.config, OpUpdateOne, withThreadID(id))
	return &ThreadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Thread.
func (c *ThreadClient) Delete() *ThreadDelete {
	mutation := newThreadMutation(c.config, OpDelete)
	return &ThreadDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ThreadClient) DeleteOne(t *Thread) *ThreadDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ThreadClient) DeleteOneID(id uuid.UUID) *ThreadDeleteOne {
	builder := c.Delete().Where(thread.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ThreadDeleteOne{builder}
}

// Query returns a query builder for Thread.
func (c *ThreadClient) Query() *ThreadQuery {
	return &ThreadQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeThread},
		inters: c.Interceptors(),
	}
}

// Get returns a Thread entity by its id.
func (c *ThreadClient) Get(ctx context.Context, id uuid.UUID) (*Thread, error) {
	return c.Query().Where(thread.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ThreadClient) GetX(ctx context.Context, id uuid.UUID) *Thread {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ThreadClient) Hooks() []Hook {
	return c.hooks.Thread
}

// Interceptors returns the client interceptors.
func (c *ThreadClient) Interceptors() []Interceptor {
	return c.inters.Thread
}

func (c *ThreadClient) mutate(ctx context.Context, m *ThreadMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ThreadCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ThreadUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ThreadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ThreadDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Thread mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Agent, Bookmark, Message, Thread, User []ent.Hook
	}
	inters struct {
		Agent, Bookmark, Message, Thread, User []ent.Interceptor
	}
)
