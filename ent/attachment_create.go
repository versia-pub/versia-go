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
	"github.com/versia-pub/versia-go/ent/user"
	"github.com/versia-pub/versia-go/pkg/versia"
)

// AttachmentCreate is the builder for creating a Attachment entity.
type AttachmentCreate struct {
	config
	mutation *AttachmentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsRemote sets the "isRemote" field.
func (ac *AttachmentCreate) SetIsRemote(b bool) *AttachmentCreate {
	ac.mutation.SetIsRemote(b)
	return ac
}

// SetURI sets the "uri" field.
func (ac *AttachmentCreate) SetURI(s string) *AttachmentCreate {
	ac.mutation.SetURI(s)
	return ac
}

// SetExtensions sets the "extensions" field.
func (ac *AttachmentCreate) SetExtensions(v versia.Extensions) *AttachmentCreate {
	ac.mutation.SetExtensions(v)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AttachmentCreate) SetCreatedAt(t time.Time) *AttachmentCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCreatedAt(t *time.Time) *AttachmentCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AttachmentCreate) SetUpdatedAt(t time.Time) *AttachmentCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableUpdatedAt(t *time.Time) *AttachmentCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDescription sets the "description" field.
func (ac *AttachmentCreate) SetDescription(s string) *AttachmentCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetSha256 sets the "sha256" field.
func (ac *AttachmentCreate) SetSha256(b []byte) *AttachmentCreate {
	ac.mutation.SetSha256(b)
	return ac
}

// SetSize sets the "size" field.
func (ac *AttachmentCreate) SetSize(i int) *AttachmentCreate {
	ac.mutation.SetSize(i)
	return ac
}

// SetBlurhash sets the "blurhash" field.
func (ac *AttachmentCreate) SetBlurhash(s string) *AttachmentCreate {
	ac.mutation.SetBlurhash(s)
	return ac
}

// SetNillableBlurhash sets the "blurhash" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableBlurhash(s *string) *AttachmentCreate {
	if s != nil {
		ac.SetBlurhash(*s)
	}
	return ac
}

// SetHeight sets the "height" field.
func (ac *AttachmentCreate) SetHeight(i int) *AttachmentCreate {
	ac.mutation.SetHeight(i)
	return ac
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableHeight(i *int) *AttachmentCreate {
	if i != nil {
		ac.SetHeight(*i)
	}
	return ac
}

// SetWidth sets the "width" field.
func (ac *AttachmentCreate) SetWidth(i int) *AttachmentCreate {
	ac.mutation.SetWidth(i)
	return ac
}

// SetNillableWidth sets the "width" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableWidth(i *int) *AttachmentCreate {
	if i != nil {
		ac.SetWidth(*i)
	}
	return ac
}

// SetFps sets the "fps" field.
func (ac *AttachmentCreate) SetFps(i int) *AttachmentCreate {
	ac.mutation.SetFps(i)
	return ac
}

// SetNillableFps sets the "fps" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableFps(i *int) *AttachmentCreate {
	if i != nil {
		ac.SetFps(*i)
	}
	return ac
}

// SetMimeType sets the "mimeType" field.
func (ac *AttachmentCreate) SetMimeType(s string) *AttachmentCreate {
	ac.mutation.SetMimeType(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AttachmentCreate) SetID(u uuid.UUID) *AttachmentCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableID(u *uuid.UUID) *AttachmentCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (ac *AttachmentCreate) SetAuthorID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetAuthorID(id)
	return ac
}

// SetAuthor sets the "author" edge to the User entity.
func (ac *AttachmentCreate) SetAuthor(u *User) *AttachmentCreate {
	return ac.SetAuthorID(u.ID)
}

// Mutation returns the AttachmentMutation object of the builder.
func (ac *AttachmentCreate) Mutation() *AttachmentMutation {
	return ac.mutation
}

// Save creates the Attachment in the database.
func (ac *AttachmentCreate) Save(ctx context.Context) (*Attachment, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AttachmentCreate) SaveX(ctx context.Context) *Attachment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AttachmentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AttachmentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AttachmentCreate) defaults() {
	if _, ok := ac.mutation.Extensions(); !ok {
		v := attachment.DefaultExtensions
		ac.mutation.SetExtensions(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := attachment.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := attachment.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := attachment.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AttachmentCreate) check() error {
	if _, ok := ac.mutation.IsRemote(); !ok {
		return &ValidationError{Name: "isRemote", err: errors.New(`ent: missing required field "Attachment.isRemote"`)}
	}
	if _, ok := ac.mutation.URI(); !ok {
		return &ValidationError{Name: "uri", err: errors.New(`ent: missing required field "Attachment.uri"`)}
	}
	if v, ok := ac.mutation.URI(); ok {
		if err := attachment.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Attachment.uri": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Extensions(); !ok {
		return &ValidationError{Name: "extensions", err: errors.New(`ent: missing required field "Attachment.extensions"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Attachment.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Attachment.updated_at"`)}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Attachment.description"`)}
	}
	if v, ok := ac.mutation.Description(); ok {
		if err := attachment.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Attachment.description": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Sha256(); !ok {
		return &ValidationError{Name: "sha256", err: errors.New(`ent: missing required field "Attachment.sha256"`)}
	}
	if _, ok := ac.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Attachment.size"`)}
	}
	if _, ok := ac.mutation.MimeType(); !ok {
		return &ValidationError{Name: "mimeType", err: errors.New(`ent: missing required field "Attachment.mimeType"`)}
	}
	if _, ok := ac.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required edge "Attachment.author"`)}
	}
	return nil
}

func (ac *AttachmentCreate) sqlSave(ctx context.Context) (*Attachment, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AttachmentCreate) createSpec() (*Attachment, *sqlgraph.CreateSpec) {
	var (
		_node = &Attachment{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(attachment.Table, sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.IsRemote(); ok {
		_spec.SetField(attachment.FieldIsRemote, field.TypeBool, value)
		_node.IsRemote = value
	}
	if value, ok := ac.mutation.URI(); ok {
		_spec.SetField(attachment.FieldURI, field.TypeString, value)
		_node.URI = value
	}
	if value, ok := ac.mutation.Extensions(); ok {
		_spec.SetField(attachment.FieldExtensions, field.TypeJSON, value)
		_node.Extensions = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(attachment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(attachment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(attachment.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.Sha256(); ok {
		_spec.SetField(attachment.FieldSha256, field.TypeBytes, value)
		_node.Sha256 = value
	}
	if value, ok := ac.mutation.Size(); ok {
		_spec.SetField(attachment.FieldSize, field.TypeInt, value)
		_node.Size = value
	}
	if value, ok := ac.mutation.Blurhash(); ok {
		_spec.SetField(attachment.FieldBlurhash, field.TypeString, value)
		_node.Blurhash = &value
	}
	if value, ok := ac.mutation.Height(); ok {
		_spec.SetField(attachment.FieldHeight, field.TypeInt, value)
		_node.Height = &value
	}
	if value, ok := ac.mutation.Width(); ok {
		_spec.SetField(attachment.FieldWidth, field.TypeInt, value)
		_node.Width = &value
	}
	if value, ok := ac.mutation.Fps(); ok {
		_spec.SetField(attachment.FieldFps, field.TypeInt, value)
		_node.Fps = &value
	}
	if value, ok := ac.mutation.MimeType(); ok {
		_spec.SetField(attachment.FieldMimeType, field.TypeString, value)
		_node.MimeType = value
	}
	if nodes := ac.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attachment.AuthorTable,
			Columns: []string{attachment.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.attachment_author = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Attachment.Create().
//		SetIsRemote(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttachmentUpsert) {
//			SetIsRemote(v+v).
//		}).
//		Exec(ctx)
func (ac *AttachmentCreate) OnConflict(opts ...sql.ConflictOption) *AttachmentUpsertOne {
	ac.conflict = opts
	return &AttachmentUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Attachment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AttachmentCreate) OnConflictColumns(columns ...string) *AttachmentUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AttachmentUpsertOne{
		create: ac,
	}
}

type (
	// AttachmentUpsertOne is the builder for "upsert"-ing
	//  one Attachment node.
	AttachmentUpsertOne struct {
		create *AttachmentCreate
	}

	// AttachmentUpsert is the "OnConflict" setter.
	AttachmentUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsRemote sets the "isRemote" field.
func (u *AttachmentUpsert) SetIsRemote(v bool) *AttachmentUpsert {
	u.Set(attachment.FieldIsRemote, v)
	return u
}

// UpdateIsRemote sets the "isRemote" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateIsRemote() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldIsRemote)
	return u
}

// SetURI sets the "uri" field.
func (u *AttachmentUpsert) SetURI(v string) *AttachmentUpsert {
	u.Set(attachment.FieldURI, v)
	return u
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateURI() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldURI)
	return u
}

// SetExtensions sets the "extensions" field.
func (u *AttachmentUpsert) SetExtensions(v versia.Extensions) *AttachmentUpsert {
	u.Set(attachment.FieldExtensions, v)
	return u
}

// UpdateExtensions sets the "extensions" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateExtensions() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldExtensions)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AttachmentUpsert) SetUpdatedAt(v time.Time) *AttachmentUpsert {
	u.Set(attachment.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateUpdatedAt() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldUpdatedAt)
	return u
}

// SetDescription sets the "description" field.
func (u *AttachmentUpsert) SetDescription(v string) *AttachmentUpsert {
	u.Set(attachment.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateDescription() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldDescription)
	return u
}

// SetSha256 sets the "sha256" field.
func (u *AttachmentUpsert) SetSha256(v []byte) *AttachmentUpsert {
	u.Set(attachment.FieldSha256, v)
	return u
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateSha256() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldSha256)
	return u
}

// SetSize sets the "size" field.
func (u *AttachmentUpsert) SetSize(v int) *AttachmentUpsert {
	u.Set(attachment.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateSize() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *AttachmentUpsert) AddSize(v int) *AttachmentUpsert {
	u.Add(attachment.FieldSize, v)
	return u
}

// SetBlurhash sets the "blurhash" field.
func (u *AttachmentUpsert) SetBlurhash(v string) *AttachmentUpsert {
	u.Set(attachment.FieldBlurhash, v)
	return u
}

// UpdateBlurhash sets the "blurhash" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateBlurhash() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldBlurhash)
	return u
}

// ClearBlurhash clears the value of the "blurhash" field.
func (u *AttachmentUpsert) ClearBlurhash() *AttachmentUpsert {
	u.SetNull(attachment.FieldBlurhash)
	return u
}

// SetHeight sets the "height" field.
func (u *AttachmentUpsert) SetHeight(v int) *AttachmentUpsert {
	u.Set(attachment.FieldHeight, v)
	return u
}

// UpdateHeight sets the "height" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateHeight() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldHeight)
	return u
}

// AddHeight adds v to the "height" field.
func (u *AttachmentUpsert) AddHeight(v int) *AttachmentUpsert {
	u.Add(attachment.FieldHeight, v)
	return u
}

// ClearHeight clears the value of the "height" field.
func (u *AttachmentUpsert) ClearHeight() *AttachmentUpsert {
	u.SetNull(attachment.FieldHeight)
	return u
}

// SetWidth sets the "width" field.
func (u *AttachmentUpsert) SetWidth(v int) *AttachmentUpsert {
	u.Set(attachment.FieldWidth, v)
	return u
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateWidth() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldWidth)
	return u
}

// AddWidth adds v to the "width" field.
func (u *AttachmentUpsert) AddWidth(v int) *AttachmentUpsert {
	u.Add(attachment.FieldWidth, v)
	return u
}

// ClearWidth clears the value of the "width" field.
func (u *AttachmentUpsert) ClearWidth() *AttachmentUpsert {
	u.SetNull(attachment.FieldWidth)
	return u
}

// SetFps sets the "fps" field.
func (u *AttachmentUpsert) SetFps(v int) *AttachmentUpsert {
	u.Set(attachment.FieldFps, v)
	return u
}

// UpdateFps sets the "fps" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateFps() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldFps)
	return u
}

// AddFps adds v to the "fps" field.
func (u *AttachmentUpsert) AddFps(v int) *AttachmentUpsert {
	u.Add(attachment.FieldFps, v)
	return u
}

// ClearFps clears the value of the "fps" field.
func (u *AttachmentUpsert) ClearFps() *AttachmentUpsert {
	u.SetNull(attachment.FieldFps)
	return u
}

// SetMimeType sets the "mimeType" field.
func (u *AttachmentUpsert) SetMimeType(v string) *AttachmentUpsert {
	u.Set(attachment.FieldMimeType, v)
	return u
}

// UpdateMimeType sets the "mimeType" field to the value that was provided on create.
func (u *AttachmentUpsert) UpdateMimeType() *AttachmentUpsert {
	u.SetExcluded(attachment.FieldMimeType)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Attachment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(attachment.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AttachmentUpsertOne) UpdateNewValues() *AttachmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(attachment.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(attachment.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Attachment.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AttachmentUpsertOne) Ignore() *AttachmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttachmentUpsertOne) DoNothing() *AttachmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttachmentCreate.OnConflict
// documentation for more info.
func (u *AttachmentUpsertOne) Update(set func(*AttachmentUpsert)) *AttachmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttachmentUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsRemote sets the "isRemote" field.
func (u *AttachmentUpsertOne) SetIsRemote(v bool) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetIsRemote(v)
	})
}

// UpdateIsRemote sets the "isRemote" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateIsRemote() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateIsRemote()
	})
}

// SetURI sets the "uri" field.
func (u *AttachmentUpsertOne) SetURI(v string) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateURI() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateURI()
	})
}

// SetExtensions sets the "extensions" field.
func (u *AttachmentUpsertOne) SetExtensions(v versia.Extensions) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetExtensions(v)
	})
}

// UpdateExtensions sets the "extensions" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateExtensions() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateExtensions()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AttachmentUpsertOne) SetUpdatedAt(v time.Time) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateUpdatedAt() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDescription sets the "description" field.
func (u *AttachmentUpsertOne) SetDescription(v string) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateDescription() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateDescription()
	})
}

// SetSha256 sets the "sha256" field.
func (u *AttachmentUpsertOne) SetSha256(v []byte) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetSha256(v)
	})
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateSha256() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateSha256()
	})
}

// SetSize sets the "size" field.
func (u *AttachmentUpsertOne) SetSize(v int) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *AttachmentUpsertOne) AddSize(v int) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateSize() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateSize()
	})
}

// SetBlurhash sets the "blurhash" field.
func (u *AttachmentUpsertOne) SetBlurhash(v string) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetBlurhash(v)
	})
}

// UpdateBlurhash sets the "blurhash" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateBlurhash() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateBlurhash()
	})
}

// ClearBlurhash clears the value of the "blurhash" field.
func (u *AttachmentUpsertOne) ClearBlurhash() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.ClearBlurhash()
	})
}

// SetHeight sets the "height" field.
func (u *AttachmentUpsertOne) SetHeight(v int) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetHeight(v)
	})
}

// AddHeight adds v to the "height" field.
func (u *AttachmentUpsertOne) AddHeight(v int) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.AddHeight(v)
	})
}

// UpdateHeight sets the "height" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateHeight() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateHeight()
	})
}

// ClearHeight clears the value of the "height" field.
func (u *AttachmentUpsertOne) ClearHeight() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.ClearHeight()
	})
}

// SetWidth sets the "width" field.
func (u *AttachmentUpsertOne) SetWidth(v int) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetWidth(v)
	})
}

// AddWidth adds v to the "width" field.
func (u *AttachmentUpsertOne) AddWidth(v int) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.AddWidth(v)
	})
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateWidth() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateWidth()
	})
}

// ClearWidth clears the value of the "width" field.
func (u *AttachmentUpsertOne) ClearWidth() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.ClearWidth()
	})
}

// SetFps sets the "fps" field.
func (u *AttachmentUpsertOne) SetFps(v int) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetFps(v)
	})
}

// AddFps adds v to the "fps" field.
func (u *AttachmentUpsertOne) AddFps(v int) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.AddFps(v)
	})
}

// UpdateFps sets the "fps" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateFps() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateFps()
	})
}

// ClearFps clears the value of the "fps" field.
func (u *AttachmentUpsertOne) ClearFps() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.ClearFps()
	})
}

// SetMimeType sets the "mimeType" field.
func (u *AttachmentUpsertOne) SetMimeType(v string) *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetMimeType(v)
	})
}

// UpdateMimeType sets the "mimeType" field to the value that was provided on create.
func (u *AttachmentUpsertOne) UpdateMimeType() *AttachmentUpsertOne {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateMimeType()
	})
}

// Exec executes the query.
func (u *AttachmentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttachmentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttachmentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AttachmentUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AttachmentUpsertOne.ID is not supported by MySQL driver. Use AttachmentUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AttachmentUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AttachmentCreateBulk is the builder for creating many Attachment entities in bulk.
type AttachmentCreateBulk struct {
	config
	err      error
	builders []*AttachmentCreate
	conflict []sql.ConflictOption
}

// Save creates the Attachment entities in the database.
func (acb *AttachmentCreateBulk) Save(ctx context.Context) ([]*Attachment, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Attachment, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttachmentMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AttachmentCreateBulk) SaveX(ctx context.Context) []*Attachment {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AttachmentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AttachmentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Attachment.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttachmentUpsert) {
//			SetIsRemote(v+v).
//		}).
//		Exec(ctx)
func (acb *AttachmentCreateBulk) OnConflict(opts ...sql.ConflictOption) *AttachmentUpsertBulk {
	acb.conflict = opts
	return &AttachmentUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Attachment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AttachmentCreateBulk) OnConflictColumns(columns ...string) *AttachmentUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AttachmentUpsertBulk{
		create: acb,
	}
}

// AttachmentUpsertBulk is the builder for "upsert"-ing
// a bulk of Attachment nodes.
type AttachmentUpsertBulk struct {
	create *AttachmentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Attachment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(attachment.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AttachmentUpsertBulk) UpdateNewValues() *AttachmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(attachment.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(attachment.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Attachment.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AttachmentUpsertBulk) Ignore() *AttachmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttachmentUpsertBulk) DoNothing() *AttachmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttachmentCreateBulk.OnConflict
// documentation for more info.
func (u *AttachmentUpsertBulk) Update(set func(*AttachmentUpsert)) *AttachmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttachmentUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsRemote sets the "isRemote" field.
func (u *AttachmentUpsertBulk) SetIsRemote(v bool) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetIsRemote(v)
	})
}

// UpdateIsRemote sets the "isRemote" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateIsRemote() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateIsRemote()
	})
}

// SetURI sets the "uri" field.
func (u *AttachmentUpsertBulk) SetURI(v string) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateURI() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateURI()
	})
}

// SetExtensions sets the "extensions" field.
func (u *AttachmentUpsertBulk) SetExtensions(v versia.Extensions) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetExtensions(v)
	})
}

// UpdateExtensions sets the "extensions" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateExtensions() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateExtensions()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AttachmentUpsertBulk) SetUpdatedAt(v time.Time) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateUpdatedAt() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDescription sets the "description" field.
func (u *AttachmentUpsertBulk) SetDescription(v string) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateDescription() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateDescription()
	})
}

// SetSha256 sets the "sha256" field.
func (u *AttachmentUpsertBulk) SetSha256(v []byte) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetSha256(v)
	})
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateSha256() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateSha256()
	})
}

// SetSize sets the "size" field.
func (u *AttachmentUpsertBulk) SetSize(v int) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *AttachmentUpsertBulk) AddSize(v int) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateSize() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateSize()
	})
}

// SetBlurhash sets the "blurhash" field.
func (u *AttachmentUpsertBulk) SetBlurhash(v string) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetBlurhash(v)
	})
}

// UpdateBlurhash sets the "blurhash" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateBlurhash() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateBlurhash()
	})
}

// ClearBlurhash clears the value of the "blurhash" field.
func (u *AttachmentUpsertBulk) ClearBlurhash() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.ClearBlurhash()
	})
}

// SetHeight sets the "height" field.
func (u *AttachmentUpsertBulk) SetHeight(v int) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetHeight(v)
	})
}

// AddHeight adds v to the "height" field.
func (u *AttachmentUpsertBulk) AddHeight(v int) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.AddHeight(v)
	})
}

// UpdateHeight sets the "height" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateHeight() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateHeight()
	})
}

// ClearHeight clears the value of the "height" field.
func (u *AttachmentUpsertBulk) ClearHeight() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.ClearHeight()
	})
}

// SetWidth sets the "width" field.
func (u *AttachmentUpsertBulk) SetWidth(v int) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetWidth(v)
	})
}

// AddWidth adds v to the "width" field.
func (u *AttachmentUpsertBulk) AddWidth(v int) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.AddWidth(v)
	})
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateWidth() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateWidth()
	})
}

// ClearWidth clears the value of the "width" field.
func (u *AttachmentUpsertBulk) ClearWidth() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.ClearWidth()
	})
}

// SetFps sets the "fps" field.
func (u *AttachmentUpsertBulk) SetFps(v int) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetFps(v)
	})
}

// AddFps adds v to the "fps" field.
func (u *AttachmentUpsertBulk) AddFps(v int) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.AddFps(v)
	})
}

// UpdateFps sets the "fps" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateFps() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateFps()
	})
}

// ClearFps clears the value of the "fps" field.
func (u *AttachmentUpsertBulk) ClearFps() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.ClearFps()
	})
}

// SetMimeType sets the "mimeType" field.
func (u *AttachmentUpsertBulk) SetMimeType(v string) *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.SetMimeType(v)
	})
}

// UpdateMimeType sets the "mimeType" field to the value that was provided on create.
func (u *AttachmentUpsertBulk) UpdateMimeType() *AttachmentUpsertBulk {
	return u.Update(func(s *AttachmentUpsert) {
		s.UpdateMimeType()
	})
}

// Exec executes the query.
func (u *AttachmentUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AttachmentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttachmentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttachmentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
