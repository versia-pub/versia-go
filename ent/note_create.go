// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/versia-pub/versia-go/ent/attachment"
	"github.com/versia-pub/versia-go/ent/note"
	"github.com/versia-pub/versia-go/ent/user"
	"github.com/versia-pub/versia-go/pkg/versia"
)

// NoteCreate is the builder for creating a Note entity.
type NoteCreate struct {
	config
	mutation *NoteMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsRemote sets the "isRemote" field.
func (nc *NoteCreate) SetIsRemote(b bool) *NoteCreate {
	nc.mutation.SetIsRemote(b)
	return nc
}

// SetURI sets the "uri" field.
func (nc *NoteCreate) SetURI(s string) *NoteCreate {
	nc.mutation.SetURI(s)
	return nc
}

// SetExtensions sets the "extensions" field.
func (nc *NoteCreate) SetExtensions(v versia.Extensions) *NoteCreate {
	nc.mutation.SetExtensions(v)
	return nc
}

// SetCreatedAt sets the "created_at" field.
func (nc *NoteCreate) SetCreatedAt(t time.Time) *NoteCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NoteCreate) SetNillableCreatedAt(t *time.Time) *NoteCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NoteCreate) SetUpdatedAt(t time.Time) *NoteCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NoteCreate) SetNillableUpdatedAt(t *time.Time) *NoteCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetSubject sets the "subject" field.
func (nc *NoteCreate) SetSubject(s string) *NoteCreate {
	nc.mutation.SetSubject(s)
	return nc
}

// SetNillableSubject sets the "subject" field if the given value is not nil.
func (nc *NoteCreate) SetNillableSubject(s *string) *NoteCreate {
	if s != nil {
		nc.SetSubject(*s)
	}
	return nc
}

// SetContent sets the "content" field.
func (nc *NoteCreate) SetContent(s string) *NoteCreate {
	nc.mutation.SetContent(s)
	return nc
}

// SetIsSensitive sets the "isSensitive" field.
func (nc *NoteCreate) SetIsSensitive(b bool) *NoteCreate {
	nc.mutation.SetIsSensitive(b)
	return nc
}

// SetNillableIsSensitive sets the "isSensitive" field if the given value is not nil.
func (nc *NoteCreate) SetNillableIsSensitive(b *bool) *NoteCreate {
	if b != nil {
		nc.SetIsSensitive(*b)
	}
	return nc
}

// SetVisibility sets the "visibility" field.
func (nc *NoteCreate) SetVisibility(n note.Visibility) *NoteCreate {
	nc.mutation.SetVisibility(n)
	return nc
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (nc *NoteCreate) SetNillableVisibility(n *note.Visibility) *NoteCreate {
	if n != nil {
		nc.SetVisibility(*n)
	}
	return nc
}

// SetID sets the "id" field.
func (nc *NoteCreate) SetID(u uuid.UUID) *NoteCreate {
	nc.mutation.SetID(u)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NoteCreate) SetNillableID(u *uuid.UUID) *NoteCreate {
	if u != nil {
		nc.SetID(*u)
	}
	return nc
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (nc *NoteCreate) SetAuthorID(id uuid.UUID) *NoteCreate {
	nc.mutation.SetAuthorID(id)
	return nc
}

// SetAuthor sets the "author" edge to the User entity.
func (nc *NoteCreate) SetAuthor(u *User) *NoteCreate {
	return nc.SetAuthorID(u.ID)
}

// AddMentionIDs adds the "mentions" edge to the User entity by IDs.
func (nc *NoteCreate) AddMentionIDs(ids ...uuid.UUID) *NoteCreate {
	nc.mutation.AddMentionIDs(ids...)
	return nc
}

// AddMentions adds the "mentions" edges to the User entity.
func (nc *NoteCreate) AddMentions(u ...*User) *NoteCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nc.AddMentionIDs(ids...)
}

// AddAttachmentIDs adds the "attachments" edge to the Attachment entity by IDs.
func (nc *NoteCreate) AddAttachmentIDs(ids ...uuid.UUID) *NoteCreate {
	nc.mutation.AddAttachmentIDs(ids...)
	return nc
}

// AddAttachments adds the "attachments" edges to the Attachment entity.
func (nc *NoteCreate) AddAttachments(a ...*Attachment) *NoteCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return nc.AddAttachmentIDs(ids...)
}

// Mutation returns the NoteMutation object of the builder.
func (nc *NoteCreate) Mutation() *NoteMutation {
	return nc.mutation
}

// Save creates the Note in the database.
func (nc *NoteCreate) Save(ctx context.Context) (*Note, error) {
	nc.defaults()
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NoteCreate) SaveX(ctx context.Context) *Note {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NoteCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NoteCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NoteCreate) defaults() {
	if _, ok := nc.mutation.Extensions(); !ok {
		v := note.DefaultExtensions
		nc.mutation.SetExtensions(v)
	}
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := note.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		v := note.DefaultUpdatedAt()
		nc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nc.mutation.IsSensitive(); !ok {
		v := note.DefaultIsSensitive
		nc.mutation.SetIsSensitive(v)
	}
	if _, ok := nc.mutation.Visibility(); !ok {
		v := note.DefaultVisibility
		nc.mutation.SetVisibility(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := note.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NoteCreate) check() error {
	if _, ok := nc.mutation.IsRemote(); !ok {
		return &ValidationError{Name: "isRemote", err: errors.New(`ent: missing required field "Note.isRemote"`)}
	}
	if _, ok := nc.mutation.URI(); !ok {
		return &ValidationError{Name: "uri", err: errors.New(`ent: missing required field "Note.uri"`)}
	}
	if v, ok := nc.mutation.URI(); ok {
		if err := note.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Note.uri": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Extensions(); !ok {
		return &ValidationError{Name: "extensions", err: errors.New(`ent: missing required field "Note.extensions"`)}
	}
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Note.created_at"`)}
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Note.updated_at"`)}
	}
	if v, ok := nc.mutation.Subject(); ok {
		if err := note.SubjectValidator(v); err != nil {
			return &ValidationError{Name: "subject", err: fmt.Errorf(`ent: validator failed for field "Note.subject": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Note.content"`)}
	}
	if _, ok := nc.mutation.IsSensitive(); !ok {
		return &ValidationError{Name: "isSensitive", err: errors.New(`ent: missing required field "Note.isSensitive"`)}
	}
	if _, ok := nc.mutation.Visibility(); !ok {
		return &ValidationError{Name: "visibility", err: errors.New(`ent: missing required field "Note.visibility"`)}
	}
	if v, ok := nc.mutation.Visibility(); ok {
		if err := note.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`ent: validator failed for field "Note.visibility": %w`, err)}
		}
	}
	if _, ok := nc.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required edge "Note.author"`)}
	}
	return nil
}

func (nc *NoteCreate) sqlSave(ctx context.Context) (*Note, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NoteCreate) createSpec() (*Note, *sqlgraph.CreateSpec) {
	var (
		_node = &Note{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(note.Table, sqlgraph.NewFieldSpec(note.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = nc.conflict
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.IsRemote(); ok {
		_spec.SetField(note.FieldIsRemote, field.TypeBool, value)
		_node.IsRemote = value
	}
	if value, ok := nc.mutation.URI(); ok {
		_spec.SetField(note.FieldURI, field.TypeString, value)
		_node.URI = value
	}
	if value, ok := nc.mutation.Extensions(); ok {
		_spec.SetField(note.FieldExtensions, field.TypeJSON, value)
		_node.Extensions = value
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(note.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.SetField(note.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.Subject(); ok {
		_spec.SetField(note.FieldSubject, field.TypeString, value)
		_node.Subject = &value
	}
	if value, ok := nc.mutation.Content(); ok {
		_spec.SetField(note.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := nc.mutation.IsSensitive(); ok {
		_spec.SetField(note.FieldIsSensitive, field.TypeBool, value)
		_node.IsSensitive = value
	}
	if value, ok := nc.mutation.Visibility(); ok {
		_spec.SetField(note.FieldVisibility, field.TypeEnum, value)
		_node.Visibility = value
	}
	if nodes := nc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   note.AuthorTable,
			Columns: []string{note.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.note_author = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.MentionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   note.MentionsTable,
			Columns: note.MentionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   note.AttachmentsTable,
			Columns: []string{note.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Note.Create().
//		SetIsRemote(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NoteUpsert) {
//			SetIsRemote(v+v).
//		}).
//		Exec(ctx)
func (nc *NoteCreate) OnConflict(opts ...sql.ConflictOption) *NoteUpsertOne {
	nc.conflict = opts
	return &NoteUpsertOne{
		create: nc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Note.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nc *NoteCreate) OnConflictColumns(columns ...string) *NoteUpsertOne {
	nc.conflict = append(nc.conflict, sql.ConflictColumns(columns...))
	return &NoteUpsertOne{
		create: nc,
	}
}

type (
	// NoteUpsertOne is the builder for "upsert"-ing
	//  one Note node.
	NoteUpsertOne struct {
		create *NoteCreate
	}

	// NoteUpsert is the "OnConflict" setter.
	NoteUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsRemote sets the "isRemote" field.
func (u *NoteUpsert) SetIsRemote(v bool) *NoteUpsert {
	u.Set(note.FieldIsRemote, v)
	return u
}

// UpdateIsRemote sets the "isRemote" field to the value that was provided on create.
func (u *NoteUpsert) UpdateIsRemote() *NoteUpsert {
	u.SetExcluded(note.FieldIsRemote)
	return u
}

// SetURI sets the "uri" field.
func (u *NoteUpsert) SetURI(v string) *NoteUpsert {
	u.Set(note.FieldURI, v)
	return u
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *NoteUpsert) UpdateURI() *NoteUpsert {
	u.SetExcluded(note.FieldURI)
	return u
}

// SetExtensions sets the "extensions" field.
func (u *NoteUpsert) SetExtensions(v versia.Extensions) *NoteUpsert {
	u.Set(note.FieldExtensions, v)
	return u
}

// UpdateExtensions sets the "extensions" field to the value that was provided on create.
func (u *NoteUpsert) UpdateExtensions() *NoteUpsert {
	u.SetExcluded(note.FieldExtensions)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NoteUpsert) SetUpdatedAt(v time.Time) *NoteUpsert {
	u.Set(note.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NoteUpsert) UpdateUpdatedAt() *NoteUpsert {
	u.SetExcluded(note.FieldUpdatedAt)
	return u
}

// SetSubject sets the "subject" field.
func (u *NoteUpsert) SetSubject(v string) *NoteUpsert {
	u.Set(note.FieldSubject, v)
	return u
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *NoteUpsert) UpdateSubject() *NoteUpsert {
	u.SetExcluded(note.FieldSubject)
	return u
}

// ClearSubject clears the value of the "subject" field.
func (u *NoteUpsert) ClearSubject() *NoteUpsert {
	u.SetNull(note.FieldSubject)
	return u
}

// SetContent sets the "content" field.
func (u *NoteUpsert) SetContent(v string) *NoteUpsert {
	u.Set(note.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *NoteUpsert) UpdateContent() *NoteUpsert {
	u.SetExcluded(note.FieldContent)
	return u
}

// SetIsSensitive sets the "isSensitive" field.
func (u *NoteUpsert) SetIsSensitive(v bool) *NoteUpsert {
	u.Set(note.FieldIsSensitive, v)
	return u
}

// UpdateIsSensitive sets the "isSensitive" field to the value that was provided on create.
func (u *NoteUpsert) UpdateIsSensitive() *NoteUpsert {
	u.SetExcluded(note.FieldIsSensitive)
	return u
}

// SetVisibility sets the "visibility" field.
func (u *NoteUpsert) SetVisibility(v note.Visibility) *NoteUpsert {
	u.Set(note.FieldVisibility, v)
	return u
}

// UpdateVisibility sets the "visibility" field to the value that was provided on create.
func (u *NoteUpsert) UpdateVisibility() *NoteUpsert {
	u.SetExcluded(note.FieldVisibility)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Note.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(note.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NoteUpsertOne) UpdateNewValues() *NoteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(note.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(note.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Note.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NoteUpsertOne) Ignore() *NoteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NoteUpsertOne) DoNothing() *NoteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NoteCreate.OnConflict
// documentation for more info.
func (u *NoteUpsertOne) Update(set func(*NoteUpsert)) *NoteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NoteUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsRemote sets the "isRemote" field.
func (u *NoteUpsertOne) SetIsRemote(v bool) *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.SetIsRemote(v)
	})
}

// UpdateIsRemote sets the "isRemote" field to the value that was provided on create.
func (u *NoteUpsertOne) UpdateIsRemote() *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateIsRemote()
	})
}

// SetURI sets the "uri" field.
func (u *NoteUpsertOne) SetURI(v string) *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *NoteUpsertOne) UpdateURI() *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateURI()
	})
}

// SetExtensions sets the "extensions" field.
func (u *NoteUpsertOne) SetExtensions(v versia.Extensions) *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.SetExtensions(v)
	})
}

// UpdateExtensions sets the "extensions" field to the value that was provided on create.
func (u *NoteUpsertOne) UpdateExtensions() *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateExtensions()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NoteUpsertOne) SetUpdatedAt(v time.Time) *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NoteUpsertOne) UpdateUpdatedAt() *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetSubject sets the "subject" field.
func (u *NoteUpsertOne) SetSubject(v string) *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.SetSubject(v)
	})
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *NoteUpsertOne) UpdateSubject() *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateSubject()
	})
}

// ClearSubject clears the value of the "subject" field.
func (u *NoteUpsertOne) ClearSubject() *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.ClearSubject()
	})
}

// SetContent sets the "content" field.
func (u *NoteUpsertOne) SetContent(v string) *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *NoteUpsertOne) UpdateContent() *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateContent()
	})
}

// SetIsSensitive sets the "isSensitive" field.
func (u *NoteUpsertOne) SetIsSensitive(v bool) *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.SetIsSensitive(v)
	})
}

// UpdateIsSensitive sets the "isSensitive" field to the value that was provided on create.
func (u *NoteUpsertOne) UpdateIsSensitive() *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateIsSensitive()
	})
}

// SetVisibility sets the "visibility" field.
func (u *NoteUpsertOne) SetVisibility(v note.Visibility) *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.SetVisibility(v)
	})
}

// UpdateVisibility sets the "visibility" field to the value that was provided on create.
func (u *NoteUpsertOne) UpdateVisibility() *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateVisibility()
	})
}

// Exec executes the query.
func (u *NoteUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NoteCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NoteUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NoteUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: NoteUpsertOne.ID is not supported by MySQL driver. Use NoteUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NoteUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NoteCreateBulk is the builder for creating many Note entities in bulk.
type NoteCreateBulk struct {
	config
	err      error
	builders []*NoteCreate
	conflict []sql.ConflictOption
}

// Save creates the Note entities in the database.
func (ncb *NoteCreateBulk) Save(ctx context.Context) ([]*Note, error) {
	if ncb.err != nil {
		return nil, ncb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Note, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NoteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ncb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NoteCreateBulk) SaveX(ctx context.Context) []*Note {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NoteCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NoteCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Note.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NoteUpsert) {
//			SetIsRemote(v+v).
//		}).
//		Exec(ctx)
func (ncb *NoteCreateBulk) OnConflict(opts ...sql.ConflictOption) *NoteUpsertBulk {
	ncb.conflict = opts
	return &NoteUpsertBulk{
		create: ncb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Note.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ncb *NoteCreateBulk) OnConflictColumns(columns ...string) *NoteUpsertBulk {
	ncb.conflict = append(ncb.conflict, sql.ConflictColumns(columns...))
	return &NoteUpsertBulk{
		create: ncb,
	}
}

// NoteUpsertBulk is the builder for "upsert"-ing
// a bulk of Note nodes.
type NoteUpsertBulk struct {
	create *NoteCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Note.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(note.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NoteUpsertBulk) UpdateNewValues() *NoteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(note.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(note.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Note.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NoteUpsertBulk) Ignore() *NoteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NoteUpsertBulk) DoNothing() *NoteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NoteCreateBulk.OnConflict
// documentation for more info.
func (u *NoteUpsertBulk) Update(set func(*NoteUpsert)) *NoteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NoteUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsRemote sets the "isRemote" field.
func (u *NoteUpsertBulk) SetIsRemote(v bool) *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.SetIsRemote(v)
	})
}

// UpdateIsRemote sets the "isRemote" field to the value that was provided on create.
func (u *NoteUpsertBulk) UpdateIsRemote() *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateIsRemote()
	})
}

// SetURI sets the "uri" field.
func (u *NoteUpsertBulk) SetURI(v string) *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *NoteUpsertBulk) UpdateURI() *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateURI()
	})
}

// SetExtensions sets the "extensions" field.
func (u *NoteUpsertBulk) SetExtensions(v versia.Extensions) *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.SetExtensions(v)
	})
}

// UpdateExtensions sets the "extensions" field to the value that was provided on create.
func (u *NoteUpsertBulk) UpdateExtensions() *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateExtensions()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NoteUpsertBulk) SetUpdatedAt(v time.Time) *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NoteUpsertBulk) UpdateUpdatedAt() *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetSubject sets the "subject" field.
func (u *NoteUpsertBulk) SetSubject(v string) *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.SetSubject(v)
	})
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *NoteUpsertBulk) UpdateSubject() *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateSubject()
	})
}

// ClearSubject clears the value of the "subject" field.
func (u *NoteUpsertBulk) ClearSubject() *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.ClearSubject()
	})
}

// SetContent sets the "content" field.
func (u *NoteUpsertBulk) SetContent(v string) *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *NoteUpsertBulk) UpdateContent() *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateContent()
	})
}

// SetIsSensitive sets the "isSensitive" field.
func (u *NoteUpsertBulk) SetIsSensitive(v bool) *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.SetIsSensitive(v)
	})
}

// UpdateIsSensitive sets the "isSensitive" field to the value that was provided on create.
func (u *NoteUpsertBulk) UpdateIsSensitive() *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateIsSensitive()
	})
}

// SetVisibility sets the "visibility" field.
func (u *NoteUpsertBulk) SetVisibility(v note.Visibility) *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.SetVisibility(v)
	})
}

// UpdateVisibility sets the "visibility" field to the value that was provided on create.
func (u *NoteUpsertBulk) UpdateVisibility() *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateVisibility()
	})
}

// Exec executes the query.
func (u *NoteUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NoteCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NoteCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NoteUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
