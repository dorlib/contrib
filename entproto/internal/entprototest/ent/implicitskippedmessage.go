// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/contrib/entproto/internal/entprototest/ent/implicitskippedmessage"
	"entgo.io/ent/dialect/sql"
)

// ImplicitSkippedMessage is the model entity for the ImplicitSkippedMessage schema.
type ImplicitSkippedMessage struct {
	config
	// ID of the ent.
	ID                         int `json:"id,omitempty"`
	depends_on_skipped_skipped *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ImplicitSkippedMessage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case implicitskippedmessage.FieldID:
			values[i] = &sql.NullInt64{}
		case implicitskippedmessage.ForeignKeys[0]: // depends_on_skipped_skipped
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ImplicitSkippedMessage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ImplicitSkippedMessage fields.
func (ism *ImplicitSkippedMessage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case implicitskippedmessage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ism.ID = int(value.Int64)
		case implicitskippedmessage.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field depends_on_skipped_skipped", value)
			} else if value.Valid {
				ism.depends_on_skipped_skipped = new(int)
				*ism.depends_on_skipped_skipped = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ImplicitSkippedMessage.
// Note that you need to call ImplicitSkippedMessage.Unwrap() before calling this method if this ImplicitSkippedMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (ism *ImplicitSkippedMessage) Update() *ImplicitSkippedMessageUpdateOne {
	return (&ImplicitSkippedMessageClient{config: ism.config}).UpdateOne(ism)
}

// Unwrap unwraps the ImplicitSkippedMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ism *ImplicitSkippedMessage) Unwrap() *ImplicitSkippedMessage {
	tx, ok := ism.config.driver.(*txDriver)
	if !ok {
		panic("ent: ImplicitSkippedMessage is not a transactional entity")
	}
	ism.config.driver = tx.drv
	return ism
}

// String implements the fmt.Stringer.
func (ism *ImplicitSkippedMessage) String() string {
	var builder strings.Builder
	builder.WriteString("ImplicitSkippedMessage(")
	builder.WriteString(fmt.Sprintf("id=%v", ism.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ImplicitSkippedMessages is a parsable slice of ImplicitSkippedMessage.
type ImplicitSkippedMessages []*ImplicitSkippedMessage

func (ism ImplicitSkippedMessages) config(cfg config) {
	for _i := range ism {
		ism[_i].config = cfg
	}
}
