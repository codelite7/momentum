// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"github.com/codelite7/momentum/api/ent/tenant"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/ent/user"
)

// Thread is the model entity for the Thread schema.
type Thread struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID pulid.ID `json:"tenant_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LastViewedAt holds the value of the "last_viewed_at" field.
	LastViewedAt time.Time `json:"last_viewed_at,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider thread.Provider `json:"provider,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ThreadQuery when eager-loading is set.
	Edges         ThreadEdges `json:"edges"`
	message_child *pulid.ID
	user_threads  *pulid.ID
	selectValues  sql.SelectValues
}

// ThreadEdges holds the relations/edges for other nodes in the graph.
type ThreadEdges struct {
	// Tenant holds the value of the tenant edge.
	Tenant *Tenant `json:"tenant,omitempty"`
	// CreatedBy holds the value of the created_by edge.
	CreatedBy *User `json:"created_by,omitempty"`
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Message `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedMessages map[string][]*Message
}

// TenantOrErr returns the Tenant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ThreadEdges) TenantOrErr() (*Tenant, error) {
	if e.Tenant != nil {
		return e.Tenant, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: tenant.Label}
	}
	return nil, &NotLoadedError{edge: "tenant"}
}

// CreatedByOrErr returns the CreatedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ThreadEdges) CreatedByOrErr() (*User, error) {
	if e.CreatedBy != nil {
		return e.CreatedBy, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "created_by"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e ThreadEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[2] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ThreadEdges) ParentOrErr() (*Message, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: message.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Thread) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case thread.FieldID, thread.FieldTenantID:
			values[i] = new(pulid.ID)
		case thread.FieldName, thread.FieldProvider:
			values[i] = new(sql.NullString)
		case thread.FieldCreatedAt, thread.FieldUpdatedAt, thread.FieldLastViewedAt:
			values[i] = new(sql.NullTime)
		case thread.ForeignKeys[0]: // message_child
			values[i] = &sql.NullScanner{S: new(pulid.ID)}
		case thread.ForeignKeys[1]: // user_threads
			values[i] = &sql.NullScanner{S: new(pulid.ID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Thread fields.
func (t *Thread) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case thread.FieldID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case thread.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case thread.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case thread.FieldTenantID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value != nil {
				t.TenantID = *value
			}
		case thread.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case thread.FieldLastViewedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_viewed_at", values[i])
			} else if value.Valid {
				t.LastViewedAt = value.Time
			}
		case thread.FieldProvider:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider", values[i])
			} else if value.Valid {
				t.Provider = thread.Provider(value.String)
			}
		case thread.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field message_child", values[i])
			} else if value.Valid {
				t.message_child = new(pulid.ID)
				*t.message_child = *value.S.(*pulid.ID)
			}
		case thread.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_threads", values[i])
			} else if value.Valid {
				t.user_threads = new(pulid.ID)
				*t.user_threads = *value.S.(*pulid.ID)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Thread.
// This includes values selected through modifiers, order, etc.
func (t *Thread) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryTenant queries the "tenant" edge of the Thread entity.
func (t *Thread) QueryTenant() *TenantQuery {
	return NewThreadClient(t.config).QueryTenant(t)
}

// QueryCreatedBy queries the "created_by" edge of the Thread entity.
func (t *Thread) QueryCreatedBy() *UserQuery {
	return NewThreadClient(t.config).QueryCreatedBy(t)
}

// QueryMessages queries the "messages" edge of the Thread entity.
func (t *Thread) QueryMessages() *MessageQuery {
	return NewThreadClient(t.config).QueryMessages(t)
}

// QueryParent queries the "parent" edge of the Thread entity.
func (t *Thread) QueryParent() *MessageQuery {
	return NewThreadClient(t.config).QueryParent(t)
}

// Update returns a builder for updating this Thread.
// Note that you need to call Thread.Unwrap() before calling this method if this Thread
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Thread) Update() *ThreadUpdateOne {
	return NewThreadClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Thread entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Thread) Unwrap() *Thread {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Thread is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Thread) String() string {
	var builder strings.Builder
	builder.WriteString("Thread(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(fmt.Sprintf("%v", t.TenantID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("last_viewed_at=")
	builder.WriteString(t.LastViewedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("provider=")
	builder.WriteString(fmt.Sprintf("%v", t.Provider))
	builder.WriteByte(')')
	return builder.String()
}

// NamedMessages returns the Messages named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Thread) NamedMessages(name string) ([]*Message, error) {
	if t.Edges.namedMessages == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedMessages[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Thread) appendNamedMessages(name string, edges ...*Message) {
	if t.Edges.namedMessages == nil {
		t.Edges.namedMessages = make(map[string][]*Message)
	}
	if len(edges) == 0 {
		t.Edges.namedMessages[name] = []*Message{}
	} else {
		t.Edges.namedMessages[name] = append(t.Edges.namedMessages[name], edges...)
	}
}

// Threads is a parsable slice of Thread.
type Threads []*Thread