// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/ent/follow"
	"github.com/lysand-org/versia-go/ent/user"
	"github.com/lysand-org/versia-go/pkg/versia"
)

// Follow is the model entity for the Follow schema.
type Follow struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// IsRemote holds the value of the "isRemote" field.
	IsRemote bool `json:"isRemote,omitempty"`
	// URI holds the value of the "uri" field.
	URI string `json:"uri,omitempty"`
	// Extensions holds the value of the "extensions" field.
	Extensions versia.Extensions `json:"extensions,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status holds the value of the "status" field.
	Status follow.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FollowQuery when eager-loading is set.
	Edges           FollowEdges `json:"edges"`
	follow_follower *uuid.UUID
	follow_followee *uuid.UUID
	selectValues    sql.SelectValues
}

// FollowEdges holds the relations/edges for other nodes in the graph.
type FollowEdges struct {
	// Follower holds the value of the follower edge.
	Follower *User `json:"follower,omitempty"`
	// Followee holds the value of the followee edge.
	Followee *User `json:"followee,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FollowerOrErr returns the Follower value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FollowEdges) FollowerOrErr() (*User, error) {
	if e.Follower != nil {
		return e.Follower, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "follower"}
}

// FolloweeOrErr returns the Followee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FollowEdges) FolloweeOrErr() (*User, error) {
	if e.Followee != nil {
		return e.Followee, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "followee"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Follow) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case follow.FieldExtensions:
			values[i] = new([]byte)
		case follow.FieldIsRemote:
			values[i] = new(sql.NullBool)
		case follow.FieldURI, follow.FieldStatus:
			values[i] = new(sql.NullString)
		case follow.FieldCreatedAt, follow.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case follow.FieldID:
			values[i] = new(uuid.UUID)
		case follow.ForeignKeys[0]: // follow_follower
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case follow.ForeignKeys[1]: // follow_followee
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Follow fields.
func (f *Follow) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case follow.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case follow.FieldIsRemote:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isRemote", values[i])
			} else if value.Valid {
				f.IsRemote = value.Bool
			}
		case follow.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				f.URI = value.String
			}
		case follow.FieldExtensions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field extensions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &f.Extensions); err != nil {
					return fmt.Errorf("unmarshal field extensions: %w", err)
				}
			}
		case follow.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case follow.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case follow.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				f.Status = follow.Status(value.String)
			}
		case follow.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field follow_follower", values[i])
			} else if value.Valid {
				f.follow_follower = new(uuid.UUID)
				*f.follow_follower = *value.S.(*uuid.UUID)
			}
		case follow.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field follow_followee", values[i])
			} else if value.Valid {
				f.follow_followee = new(uuid.UUID)
				*f.follow_followee = *value.S.(*uuid.UUID)
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Follow.
// This includes values selected through modifiers, order, etc.
func (f *Follow) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryFollower queries the "follower" edge of the Follow entity.
func (f *Follow) QueryFollower() *UserQuery {
	return NewFollowClient(f.config).QueryFollower(f)
}

// QueryFollowee queries the "followee" edge of the Follow entity.
func (f *Follow) QueryFollowee() *UserQuery {
	return NewFollowClient(f.config).QueryFollowee(f)
}

// Update returns a builder for updating this Follow.
// Note that you need to call Follow.Unwrap() before calling this method if this Follow
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Follow) Update() *FollowUpdateOne {
	return NewFollowClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Follow entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Follow) Unwrap() *Follow {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Follow is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Follow) String() string {
	var builder strings.Builder
	builder.WriteString("Follow(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("isRemote=")
	builder.WriteString(fmt.Sprintf("%v", f.IsRemote))
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(f.URI)
	builder.WriteString(", ")
	builder.WriteString("extensions=")
	builder.WriteString(fmt.Sprintf("%v", f.Extensions))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", f.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Follows is a parsable slice of Follow.
type Follows []*Follow
