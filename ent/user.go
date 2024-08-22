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
	"github.com/lysand-org/versia-go/ent/image"
	"github.com/lysand-org/versia-go/ent/user"
	"github.com/lysand-org/versia-go/pkg/versia"
)

// User is the model entity for the User schema.
type User struct {
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
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// PasswordHash holds the value of the "passwordHash" field.
	PasswordHash *[]byte `json:"passwordHash,omitempty"`
	// DisplayName holds the value of the "displayName" field.
	DisplayName *string `json:"displayName,omitempty"`
	// Biography holds the value of the "biography" field.
	Biography *string `json:"biography,omitempty"`
	// PublicKey holds the value of the "publicKey" field.
	PublicKey []byte `json:"publicKey,omitempty"`
	// PublicKeyActor holds the value of the "publicKeyActor" field.
	PublicKeyActor string `json:"publicKeyActor,omitempty"`
	// PublicKeyAlgorithm holds the value of the "publicKeyAlgorithm" field.
	PublicKeyAlgorithm string `json:"publicKeyAlgorithm,omitempty"`
	// PrivateKey holds the value of the "privateKey" field.
	PrivateKey []byte `json:"privateKey,omitempty"`
	// Indexable holds the value of the "indexable" field.
	Indexable bool `json:"indexable,omitempty"`
	// PrivacyLevel holds the value of the "privacyLevel" field.
	PrivacyLevel user.PrivacyLevel `json:"privacyLevel,omitempty"`
	// Fields holds the value of the "fields" field.
	Fields []versia.UserField `json:"fields,omitempty"`
	// Inbox holds the value of the "inbox" field.
	Inbox string `json:"inbox,omitempty"`
	// Featured holds the value of the "featured" field.
	Featured string `json:"featured,omitempty"`
	// Followers holds the value of the "followers" field.
	Followers string `json:"followers,omitempty"`
	// Following holds the value of the "following" field.
	Following string `json:"following,omitempty"`
	// Outbox holds the value of the "outbox" field.
	Outbox string `json:"outbox,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges             UserEdges `json:"edges"`
	user_avatar_image *int
	user_header_image *int
	selectValues      sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// AvatarImage holds the value of the avatarImage edge.
	AvatarImage *Image `json:"avatarImage,omitempty"`
	// HeaderImage holds the value of the headerImage edge.
	HeaderImage *Image `json:"headerImage,omitempty"`
	// AuthoredNotes holds the value of the authoredNotes edge.
	AuthoredNotes []*Note `json:"authoredNotes,omitempty"`
	// MentionedNotes holds the value of the mentionedNotes edge.
	MentionedNotes []*Note `json:"mentionedNotes,omitempty"`
	// Servers holds the value of the servers edge.
	Servers []*InstanceMetadata `json:"servers,omitempty"`
	// ModeratedServers holds the value of the moderatedServers edge.
	ModeratedServers []*InstanceMetadata `json:"moderatedServers,omitempty"`
	// AdministeredServers holds the value of the administeredServers edge.
	AdministeredServers []*InstanceMetadata `json:"administeredServers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// AvatarImageOrErr returns the AvatarImage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) AvatarImageOrErr() (*Image, error) {
	if e.AvatarImage != nil {
		return e.AvatarImage, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: image.Label}
	}
	return nil, &NotLoadedError{edge: "avatarImage"}
}

// HeaderImageOrErr returns the HeaderImage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) HeaderImageOrErr() (*Image, error) {
	if e.HeaderImage != nil {
		return e.HeaderImage, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: image.Label}
	}
	return nil, &NotLoadedError{edge: "headerImage"}
}

// AuthoredNotesOrErr returns the AuthoredNotes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AuthoredNotesOrErr() ([]*Note, error) {
	if e.loadedTypes[2] {
		return e.AuthoredNotes, nil
	}
	return nil, &NotLoadedError{edge: "authoredNotes"}
}

// MentionedNotesOrErr returns the MentionedNotes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MentionedNotesOrErr() ([]*Note, error) {
	if e.loadedTypes[3] {
		return e.MentionedNotes, nil
	}
	return nil, &NotLoadedError{edge: "mentionedNotes"}
}

// ServersOrErr returns the Servers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ServersOrErr() ([]*InstanceMetadata, error) {
	if e.loadedTypes[4] {
		return e.Servers, nil
	}
	return nil, &NotLoadedError{edge: "servers"}
}

// ModeratedServersOrErr returns the ModeratedServers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ModeratedServersOrErr() ([]*InstanceMetadata, error) {
	if e.loadedTypes[5] {
		return e.ModeratedServers, nil
	}
	return nil, &NotLoadedError{edge: "moderatedServers"}
}

// AdministeredServersOrErr returns the AdministeredServers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AdministeredServersOrErr() ([]*InstanceMetadata, error) {
	if e.loadedTypes[6] {
		return e.AdministeredServers, nil
	}
	return nil, &NotLoadedError{edge: "administeredServers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldExtensions, user.FieldPasswordHash, user.FieldPublicKey, user.FieldPrivateKey, user.FieldFields:
			values[i] = new([]byte)
		case user.FieldIsRemote, user.FieldIndexable:
			values[i] = new(sql.NullBool)
		case user.FieldURI, user.FieldUsername, user.FieldDisplayName, user.FieldBiography, user.FieldPublicKeyActor, user.FieldPublicKeyAlgorithm, user.FieldPrivacyLevel, user.FieldInbox, user.FieldFeatured, user.FieldFollowers, user.FieldFollowing, user.FieldOutbox:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		case user.ForeignKeys[0]: // user_avatar_image
			values[i] = new(sql.NullInt64)
		case user.ForeignKeys[1]: // user_header_image
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldIsRemote:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isRemote", values[i])
			} else if value.Valid {
				u.IsRemote = value.Bool
			}
		case user.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				u.URI = value.String
			}
		case user.FieldExtensions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field extensions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.Extensions); err != nil {
					return fmt.Errorf("unmarshal field extensions: %w", err)
				}
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPasswordHash:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field passwordHash", values[i])
			} else if value != nil {
				u.PasswordHash = value
			}
		case user.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field displayName", values[i])
			} else if value.Valid {
				u.DisplayName = new(string)
				*u.DisplayName = value.String
			}
		case user.FieldBiography:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field biography", values[i])
			} else if value.Valid {
				u.Biography = new(string)
				*u.Biography = value.String
			}
		case user.FieldPublicKey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field publicKey", values[i])
			} else if value != nil {
				u.PublicKey = *value
			}
		case user.FieldPublicKeyActor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field publicKeyActor", values[i])
			} else if value.Valid {
				u.PublicKeyActor = value.String
			}
		case user.FieldPublicKeyAlgorithm:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field publicKeyAlgorithm", values[i])
			} else if value.Valid {
				u.PublicKeyAlgorithm = value.String
			}
		case user.FieldPrivateKey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field privateKey", values[i])
			} else if value != nil {
				u.PrivateKey = *value
			}
		case user.FieldIndexable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field indexable", values[i])
			} else if value.Valid {
				u.Indexable = value.Bool
			}
		case user.FieldPrivacyLevel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field privacyLevel", values[i])
			} else if value.Valid {
				u.PrivacyLevel = user.PrivacyLevel(value.String)
			}
		case user.FieldFields:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field fields", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.Fields); err != nil {
					return fmt.Errorf("unmarshal field fields: %w", err)
				}
			}
		case user.FieldInbox:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field inbox", values[i])
			} else if value.Valid {
				u.Inbox = value.String
			}
		case user.FieldFeatured:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field featured", values[i])
			} else if value.Valid {
				u.Featured = value.String
			}
		case user.FieldFollowers:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field followers", values[i])
			} else if value.Valid {
				u.Followers = value.String
			}
		case user.FieldFollowing:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field following", values[i])
			} else if value.Valid {
				u.Following = value.String
			}
		case user.FieldOutbox:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outbox", values[i])
			} else if value.Valid {
				u.Outbox = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_avatar_image", value)
			} else if value.Valid {
				u.user_avatar_image = new(int)
				*u.user_avatar_image = int(value.Int64)
			}
		case user.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_header_image", value)
			} else if value.Valid {
				u.user_header_image = new(int)
				*u.user_header_image = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryAvatarImage queries the "avatarImage" edge of the User entity.
func (u *User) QueryAvatarImage() *ImageQuery {
	return NewUserClient(u.config).QueryAvatarImage(u)
}

// QueryHeaderImage queries the "headerImage" edge of the User entity.
func (u *User) QueryHeaderImage() *ImageQuery {
	return NewUserClient(u.config).QueryHeaderImage(u)
}

// QueryAuthoredNotes queries the "authoredNotes" edge of the User entity.
func (u *User) QueryAuthoredNotes() *NoteQuery {
	return NewUserClient(u.config).QueryAuthoredNotes(u)
}

// QueryMentionedNotes queries the "mentionedNotes" edge of the User entity.
func (u *User) QueryMentionedNotes() *NoteQuery {
	return NewUserClient(u.config).QueryMentionedNotes(u)
}

// QueryServers queries the "servers" edge of the User entity.
func (u *User) QueryServers() *InstanceMetadataQuery {
	return NewUserClient(u.config).QueryServers(u)
}

// QueryModeratedServers queries the "moderatedServers" edge of the User entity.
func (u *User) QueryModeratedServers() *InstanceMetadataQuery {
	return NewUserClient(u.config).QueryModeratedServers(u)
}

// QueryAdministeredServers queries the "administeredServers" edge of the User entity.
func (u *User) QueryAdministeredServers() *InstanceMetadataQuery {
	return NewUserClient(u.config).QueryAdministeredServers(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("isRemote=")
	builder.WriteString(fmt.Sprintf("%v", u.IsRemote))
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(u.URI)
	builder.WriteString(", ")
	builder.WriteString("extensions=")
	builder.WriteString(fmt.Sprintf("%v", u.Extensions))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	if v := u.PasswordHash; v != nil {
		builder.WriteString("passwordHash=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.DisplayName; v != nil {
		builder.WriteString("displayName=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Biography; v != nil {
		builder.WriteString("biography=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("publicKey=")
	builder.WriteString(fmt.Sprintf("%v", u.PublicKey))
	builder.WriteString(", ")
	builder.WriteString("publicKeyActor=")
	builder.WriteString(u.PublicKeyActor)
	builder.WriteString(", ")
	builder.WriteString("publicKeyAlgorithm=")
	builder.WriteString(u.PublicKeyAlgorithm)
	builder.WriteString(", ")
	builder.WriteString("privateKey=")
	builder.WriteString(fmt.Sprintf("%v", u.PrivateKey))
	builder.WriteString(", ")
	builder.WriteString("indexable=")
	builder.WriteString(fmt.Sprintf("%v", u.Indexable))
	builder.WriteString(", ")
	builder.WriteString("privacyLevel=")
	builder.WriteString(fmt.Sprintf("%v", u.PrivacyLevel))
	builder.WriteString(", ")
	builder.WriteString("fields=")
	builder.WriteString(fmt.Sprintf("%v", u.Fields))
	builder.WriteString(", ")
	builder.WriteString("inbox=")
	builder.WriteString(u.Inbox)
	builder.WriteString(", ")
	builder.WriteString("featured=")
	builder.WriteString(u.Featured)
	builder.WriteString(", ")
	builder.WriteString("followers=")
	builder.WriteString(u.Followers)
	builder.WriteString(", ")
	builder.WriteString("following=")
	builder.WriteString(u.Following)
	builder.WriteString(", ")
	builder.WriteString("outbox=")
	builder.WriteString(u.Outbox)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
