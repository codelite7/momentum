// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/response"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/ent/user"
	"github.com/google/uuid"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageQuery when eager-loading is set.
	Edges           MessageEdges `json:"edges"`
	thread_messages *uuid.UUID
	user_messages   *uuid.UUID
	selectValues    sql.SelectValues
}

// MessageEdges holds the relations/edges for other nodes in the graph.
type MessageEdges struct {
	// SentBy holds the value of the sent_by edge.
	SentBy *User `json:"sent_by,omitempty"`
	// Thread holds the value of the thread edge.
	Thread *Thread `json:"thread,omitempty"`
	// Bookmarks holds the value of the bookmarks edge.
	Bookmarks []*Bookmark `json:"bookmarks,omitempty"`
	// Response holds the value of the response edge.
	Response *Response `json:"response,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int

	namedBookmarks map[string][]*Bookmark
}

// SentByOrErr returns the SentBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) SentByOrErr() (*User, error) {
	if e.SentBy != nil {
		return e.SentBy, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "sent_by"}
}

// ThreadOrErr returns the Thread value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) ThreadOrErr() (*Thread, error) {
	if e.Thread != nil {
		return e.Thread, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: thread.Label}
	}
	return nil, &NotLoadedError{edge: "thread"}
}

// BookmarksOrErr returns the Bookmarks value or an error if the edge
// was not loaded in eager-loading.
func (e MessageEdges) BookmarksOrErr() ([]*Bookmark, error) {
	if e.loadedTypes[2] {
		return e.Bookmarks, nil
	}
	return nil, &NotLoadedError{edge: "bookmarks"}
}

// ResponseOrErr returns the Response value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) ResponseOrErr() (*Response, error) {
	if e.Response != nil {
		return e.Response, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: response.Label}
	}
	return nil, &NotLoadedError{edge: "response"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldContent:
			values[i] = new(sql.NullString)
		case message.FieldCreatedAt, message.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case message.FieldID:
			values[i] = new(uuid.UUID)
		case message.ForeignKeys[0]: // thread_messages
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case message.ForeignKeys[1]: // user_messages
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case message.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case message.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case message.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				m.Content = value.String
			}
		case message.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field thread_messages", values[i])
			} else if value.Valid {
				m.thread_messages = new(uuid.UUID)
				*m.thread_messages = *value.S.(*uuid.UUID)
			}
		case message.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_messages", values[i])
			} else if value.Valid {
				m.user_messages = new(uuid.UUID)
				*m.user_messages = *value.S.(*uuid.UUID)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Message.
// This includes values selected through modifiers, order, etc.
func (m *Message) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QuerySentBy queries the "sent_by" edge of the Message entity.
func (m *Message) QuerySentBy() *UserQuery {
	return NewMessageClient(m.config).QuerySentBy(m)
}

// QueryThread queries the "thread" edge of the Message entity.
func (m *Message) QueryThread() *ThreadQuery {
	return NewMessageClient(m.config).QueryThread(m)
}

// QueryBookmarks queries the "bookmarks" edge of the Message entity.
func (m *Message) QueryBookmarks() *BookmarkQuery {
	return NewMessageClient(m.config).QueryBookmarks(m)
}

// QueryResponse queries the "response" edge of the Message entity.
func (m *Message) QueryResponse() *ResponseQuery {
	return NewMessageClient(m.config).QueryResponse(m)
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return NewMessageClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(m.Content)
	builder.WriteByte(')')
	return builder.String()
}

// NamedBookmarks returns the Bookmarks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Message) NamedBookmarks(name string) ([]*Bookmark, error) {
	if m.Edges.namedBookmarks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedBookmarks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Message) appendNamedBookmarks(name string, edges ...*Bookmark) {
	if m.Edges.namedBookmarks == nil {
		m.Edges.namedBookmarks = make(map[string][]*Bookmark)
	}
	if len(edges) == 0 {
		m.Edges.namedBookmarks[name] = []*Bookmark{}
	} else {
		m.Edges.namedBookmarks[name] = append(m.Edges.namedBookmarks[name], edges...)
	}
}

// Messages is a parsable slice of Message.
type Messages []*Message
