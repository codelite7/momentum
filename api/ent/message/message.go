// Code generated by ent, DO NOT EDIT.

package message

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// EdgeSentBy holds the string denoting the sent_by edge name in mutations.
	EdgeSentBy = "sent_by"
	// EdgeThread holds the string denoting the thread edge name in mutations.
	EdgeThread = "thread"
	// EdgeBookmarks holds the string denoting the bookmarks edge name in mutations.
	EdgeBookmarks = "bookmarks"
	// Table holds the table name of the message in the database.
	Table = "messages"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "messages"
	// TenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantInverseTable = "tenants"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "tenant_id"
	// SentByTable is the table that holds the sent_by relation/edge.
	SentByTable = "messages"
	// SentByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SentByInverseTable = "users"
	// SentByColumn is the table column denoting the sent_by relation/edge.
	SentByColumn = "user_messages"
	// ThreadTable is the table that holds the thread relation/edge.
	ThreadTable = "messages"
	// ThreadInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	ThreadInverseTable = "threads"
	// ThreadColumn is the table column denoting the thread relation/edge.
	ThreadColumn = "thread_messages"
	// BookmarksTable is the table that holds the bookmarks relation/edge.
	BookmarksTable = "bookmarks"
	// BookmarksInverseTable is the table name for the Bookmark entity.
	// It exists in this package in order to avoid circular dependency with the "bookmark" package.
	BookmarksInverseTable = "bookmarks"
	// BookmarksColumn is the table column denoting the bookmarks relation/edge.
	BookmarksColumn = "message_bookmarks"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTenantID,
	FieldContent,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "messages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"thread_messages",
	"user_messages",
}

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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.ID
)

// OrderOption defines the ordering options for the Message queries.
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

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByTenantField orders the results by tenant field.
func ByTenantField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTenantStep(), sql.OrderByField(field, opts...))
	}
}

// BySentByField orders the results by sent_by field.
func BySentByField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSentByStep(), sql.OrderByField(field, opts...))
	}
}

// ByThreadField orders the results by thread field.
func ByThreadField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newThreadStep(), sql.OrderByField(field, opts...))
	}
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
func newTenantStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TenantInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, TenantTable, TenantColumn),
	)
}
func newSentByStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SentByInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SentByTable, SentByColumn),
	)
}
func newThreadStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ThreadTable, ThreadColumn),
	)
}
func newBookmarksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookmarksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BookmarksTable, BookmarksColumn),
	)
}
