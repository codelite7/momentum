// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldWorkosUserID holds the string denoting the workos_user_id field in the database.
	FieldWorkosUserID = "workos_user_id"
	// EdgeBookmarks holds the string denoting the bookmarks edge name in mutations.
	EdgeBookmarks = "bookmarks"
	// EdgeThreads holds the string denoting the threads edge name in mutations.
	EdgeThreads = "threads"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// EdgeTenants holds the string denoting the tenants edge name in mutations.
	EdgeTenants = "tenants"
	// EdgeActiveTenant holds the string denoting the active_tenant edge name in mutations.
	EdgeActiveTenant = "active_tenant"
	// Table holds the table name of the user in the database.
	Table = "users"
	// BookmarksTable is the table that holds the bookmarks relation/edge.
	BookmarksTable = "bookmarks"
	// BookmarksInverseTable is the table name for the Bookmark entity.
	// It exists in this package in order to avoid circular dependency with the "bookmark" package.
	BookmarksInverseTable = "bookmarks"
	// BookmarksColumn is the table column denoting the bookmarks relation/edge.
	BookmarksColumn = "user_bookmarks"
	// ThreadsTable is the table that holds the threads relation/edge.
	ThreadsTable = "threads"
	// ThreadsInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	ThreadsInverseTable = "threads"
	// ThreadsColumn is the table column denoting the threads relation/edge.
	ThreadsColumn = "user_threads"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "user_messages"
	// TenantsTable is the table that holds the tenants relation/edge. The primary key declared below.
	TenantsTable = "user_tenants"
	// TenantsInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantsInverseTable = "tenants"
	// ActiveTenantTable is the table that holds the active_tenant relation/edge.
	ActiveTenantTable = "users"
	// ActiveTenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	ActiveTenantInverseTable = "tenants"
	// ActiveTenantColumn is the table column denoting the active_tenant relation/edge.
	ActiveTenantColumn = "user_active_tenant"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEmail,
	FieldWorkosUserID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_active_tenant",
}

var (
	// TenantsPrimaryKey and TenantsColumn2 are the table columns denoting the
	// primary key for the tenants relation (M2M).
	TenantsPrimaryKey = []string{"user_id", "tenant_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// WorkosUserIDValidator is a validator for the "workos_user_id" field. It is called by the builders before save.
	WorkosUserIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.ID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByWorkosUserID orders the results by the workos_user_id field.
func ByWorkosUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkosUserID, opts...).ToFunc()
}

// ByBookmarksCount orders the results by bookmarks count.
func ByBookmarksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBookmarksStep(), opts...)
	}
}

// ByBookmarks orders the results by bookmarks terms.
func ByBookmarks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBookmarksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByThreadsCount orders the results by threads count.
func ByThreadsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newThreadsStep(), opts...)
	}
}

// ByThreads orders the results by threads terms.
func ByThreads(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newThreadsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMessagesCount orders the results by messages count.
func ByMessagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMessagesStep(), opts...)
	}
}

// ByMessages orders the results by messages terms.
func ByMessages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMessagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTenantsCount orders the results by tenants count.
func ByTenantsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTenantsStep(), opts...)
	}
}

// ByTenants orders the results by tenants terms.
func ByTenants(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTenantsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByActiveTenantField orders the results by active_tenant field.
func ByActiveTenantField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActiveTenantStep(), sql.OrderByField(field, opts...))
	}
}
func newBookmarksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookmarksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BookmarksTable, BookmarksColumn),
	)
}
func newThreadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ThreadsTable, ThreadsColumn),
	)
}
func newMessagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MessagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
	)
}
func newTenantsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TenantsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TenantsTable, TenantsPrimaryKey...),
	)
}
func newActiveTenantStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActiveTenantInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ActiveTenantTable, ActiveTenantColumn),
	)
}