// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/codelite7/momentum/api/ent/bookmark"
	"github.com/google/uuid"
)

// Bookmark is the model entity for the Bookmark schema.
type Bookmark struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bookmark) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bookmark.FieldCreatedAt, bookmark.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case bookmark.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bookmark fields.
func (b *Bookmark) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bookmark.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case bookmark.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case bookmark.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Bookmark.
// This includes values selected through modifiers, order, etc.
func (b *Bookmark) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Bookmark.
// Note that you need to call Bookmark.Unwrap() before calling this method if this Bookmark
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bookmark) Update() *BookmarkUpdateOne {
	return NewBookmarkClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Bookmark entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Bookmark) Unwrap() *Bookmark {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bookmark is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bookmark) String() string {
	var builder strings.Builder
	builder.WriteString("Bookmark(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Bookmarks is a parsable slice of Bookmark.
type Bookmarks []*Bookmark
