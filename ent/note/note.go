// Code generated by ent, DO NOT EDIT.

package note

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/versia-pub/versia-go/pkg/versia"
)

const (
	// Label holds the string label denoting the note type in the database.
	Label = "note"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsRemote holds the string denoting the isremote field in the database.
	FieldIsRemote = "is_remote"
	// FieldURI holds the string denoting the uri field in the database.
	FieldURI = "uri"
	// FieldExtensions holds the string denoting the extensions field in the database.
	FieldExtensions = "extensions"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSubject holds the string denoting the subject field in the database.
	FieldSubject = "subject"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldIsSensitive holds the string denoting the issensitive field in the database.
	FieldIsSensitive = "is_sensitive"
	// FieldVisibility holds the string denoting the visibility field in the database.
	FieldVisibility = "visibility"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeMentions holds the string denoting the mentions edge name in mutations.
	EdgeMentions = "mentions"
	// EdgeAttachments holds the string denoting the attachments edge name in mutations.
	EdgeAttachments = "attachments"
	// Table holds the table name of the note in the database.
	Table = "notes"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "notes"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "note_author"
	// MentionsTable is the table that holds the mentions relation/edge. The primary key declared below.
	MentionsTable = "note_mentions"
	// MentionsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	MentionsInverseTable = "users"
	// AttachmentsTable is the table that holds the attachments relation/edge.
	AttachmentsTable = "attachments"
	// AttachmentsInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentsInverseTable = "attachments"
	// AttachmentsColumn is the table column denoting the attachments relation/edge.
	AttachmentsColumn = "note_attachments"
)

// Columns holds all SQL columns for note fields.
var Columns = []string{
	FieldID,
	FieldIsRemote,
	FieldURI,
	FieldExtensions,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSubject,
	FieldContent,
	FieldIsSensitive,
	FieldVisibility,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "notes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"note_author",
}

var (
	// MentionsPrimaryKey and MentionsColumn2 are the table columns denoting the
	// primary key for the mentions relation (M2M).
	MentionsPrimaryKey = []string{"note_id", "user_id"}
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
	// URIValidator is a validator for the "uri" field. It is called by the builders before save.
	URIValidator func(string) error
	// DefaultExtensions holds the default value on creation for the "extensions" field.
	DefaultExtensions versia.Extensions
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// SubjectValidator is a validator for the "subject" field. It is called by the builders before save.
	SubjectValidator func(string) error
	// DefaultIsSensitive holds the default value on creation for the "isSensitive" field.
	DefaultIsSensitive bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Visibility defines the type for the "visibility" enum field.
type Visibility string

// VisibilityPublic is the default value of the Visibility enum.
const DefaultVisibility = VisibilityPublic

// Visibility values.
const (
	VisibilityPublic    Visibility = "public"
	VisibilityUnlisted  Visibility = "unlisted"
	VisibilityFollowers Visibility = "followers"
	VisibilityDirect    Visibility = "direct"
)

func (v Visibility) String() string {
	return string(v)
}

// VisibilityValidator is a validator for the "visibility" field enum values. It is called by the builders before save.
func VisibilityValidator(v Visibility) error {
	switch v {
	case VisibilityPublic, VisibilityUnlisted, VisibilityFollowers, VisibilityDirect:
		return nil
	default:
		return fmt.Errorf("note: invalid enum value for visibility field: %q", v)
	}
}

// OrderOption defines the ordering options for the Note queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIsRemote orders the results by the isRemote field.
func ByIsRemote(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsRemote, opts...).ToFunc()
}

// ByURI orders the results by the uri field.
func ByURI(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURI, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// BySubject orders the results by the subject field.
func BySubject(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubject, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByIsSensitive orders the results by the isSensitive field.
func ByIsSensitive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsSensitive, opts...).ToFunc()
}

// ByVisibility orders the results by the visibility field.
func ByVisibility(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisibility, opts...).ToFunc()
}

// ByAuthorField orders the results by author field.
func ByAuthorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), sql.OrderByField(field, opts...))
	}
}

// ByMentionsCount orders the results by mentions count.
func ByMentionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMentionsStep(), opts...)
	}
}

// ByMentions orders the results by mentions terms.
func ByMentions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMentionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAttachmentsCount orders the results by attachments count.
func ByAttachmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAttachmentsStep(), opts...)
	}
}

// ByAttachments orders the results by attachments terms.
func ByAttachments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttachmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AuthorTable, AuthorColumn),
	)
}
func newMentionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MentionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, MentionsTable, MentionsPrimaryKey...),
	)
}
func newAttachmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttachmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AttachmentsTable, AttachmentsColumn),
	)
}
