// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/versia-pub/versia-go/ent/image"
	"github.com/versia-pub/versia-go/ent/predicate"
)

// ImageUpdate is the builder for updating Image entities.
type ImageUpdate struct {
	config
	hooks    []Hook
	mutation *ImageMutation
}

// Where appends a list predicates to the ImageUpdate builder.
func (iu *ImageUpdate) Where(ps ...predicate.Image) *ImageUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetURL sets the "url" field.
func (iu *ImageUpdate) SetURL(s string) *ImageUpdate {
	iu.mutation.SetURL(s)
	return iu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableURL(s *string) *ImageUpdate {
	if s != nil {
		iu.SetURL(*s)
	}
	return iu
}

// SetMimeType sets the "mimeType" field.
func (iu *ImageUpdate) SetMimeType(s string) *ImageUpdate {
	iu.mutation.SetMimeType(s)
	return iu
}

// SetNillableMimeType sets the "mimeType" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableMimeType(s *string) *ImageUpdate {
	if s != nil {
		iu.SetMimeType(*s)
	}
	return iu
}

// Mutation returns the ImageMutation object of the builder.
func (iu *ImageUpdate) Mutation() *ImageMutation {
	return iu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImageUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImageUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImageUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ImageUpdate) check() error {
	if v, ok := iu.mutation.URL(); ok {
		if err := image.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Image.url": %w`, err)}
		}
	}
	return nil
}

func (iu *ImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(image.Table, image.Columns, sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.URL(); ok {
		_spec.SetField(image.FieldURL, field.TypeString, value)
	}
	if value, ok := iu.mutation.MimeType(); ok {
		_spec.SetField(image.FieldMimeType, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ImageUpdateOne is the builder for updating a single Image entity.
type ImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImageMutation
}

// SetURL sets the "url" field.
func (iuo *ImageUpdateOne) SetURL(s string) *ImageUpdateOne {
	iuo.mutation.SetURL(s)
	return iuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableURL(s *string) *ImageUpdateOne {
	if s != nil {
		iuo.SetURL(*s)
	}
	return iuo
}

// SetMimeType sets the "mimeType" field.
func (iuo *ImageUpdateOne) SetMimeType(s string) *ImageUpdateOne {
	iuo.mutation.SetMimeType(s)
	return iuo
}

// SetNillableMimeType sets the "mimeType" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableMimeType(s *string) *ImageUpdateOne {
	if s != nil {
		iuo.SetMimeType(*s)
	}
	return iuo
}

// Mutation returns the ImageMutation object of the builder.
func (iuo *ImageUpdateOne) Mutation() *ImageMutation {
	return iuo.mutation
}

// Where appends a list predicates to the ImageUpdate builder.
func (iuo *ImageUpdateOne) Where(ps ...predicate.Image) *ImageUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImageUpdateOne) Select(field string, fields ...string) *ImageUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Image entity.
func (iuo *ImageUpdateOne) Save(ctx context.Context) (*Image, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImageUpdateOne) SaveX(ctx context.Context) *Image {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImageUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImageUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ImageUpdateOne) check() error {
	if v, ok := iuo.mutation.URL(); ok {
		if err := image.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Image.url": %w`, err)}
		}
	}
	return nil
}

func (iuo *ImageUpdateOne) sqlSave(ctx context.Context) (_node *Image, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(image.Table, image.Columns, sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Image.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, image.FieldID)
		for _, f := range fields {
			if !image.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != image.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.URL(); ok {
		_spec.SetField(image.FieldURL, field.TypeString, value)
	}
	if value, ok := iuo.mutation.MimeType(); ok {
		_spec.SetField(image.FieldMimeType, field.TypeString, value)
	}
	_node = &Image{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
