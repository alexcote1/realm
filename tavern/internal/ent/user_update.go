// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"realm.pub/tavern/internal/ent/predicate"
	"realm.pub/tavern/internal/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetPhotoURL sets the "photo_url" field.
func (uu *UserUpdate) SetPhotoURL(s string) *UserUpdate {
	uu.mutation.SetPhotoURL(s)
	return uu
}

// SetSessionToken sets the "session_token" field.
func (uu *UserUpdate) SetSessionToken(s string) *UserUpdate {
	uu.mutation.SetSessionToken(s)
	return uu
}

// SetNillableSessionToken sets the "session_token" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSessionToken(s *string) *UserUpdate {
	if s != nil {
		uu.SetSessionToken(*s)
	}
	return uu
}

// SetIsActivated sets the "is_activated" field.
func (uu *UserUpdate) SetIsActivated(b bool) *UserUpdate {
	uu.mutation.SetIsActivated(b)
	return uu
}

// SetNillableIsActivated sets the "is_activated" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsActivated(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsActivated(*b)
	}
	return uu
}

// SetIsAdmin sets the "is_admin" field.
func (uu *UserUpdate) SetIsAdmin(b bool) *UserUpdate {
	uu.mutation.SetIsAdmin(b)
	return uu
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsAdmin(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsAdmin(*b)
	}
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.SessionToken(); ok {
		if err := user.SessionTokenValidator(v); err != nil {
			return &ValidationError{Name: "session_token", err: fmt.Errorf(`ent: validator failed for field "User.session_token": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.PhotoURL(); ok {
		_spec.SetField(user.FieldPhotoURL, field.TypeString, value)
	}
	if value, ok := uu.mutation.SessionToken(); ok {
		_spec.SetField(user.FieldSessionToken, field.TypeString, value)
	}
	if value, ok := uu.mutation.IsActivated(); ok {
		_spec.SetField(user.FieldIsActivated, field.TypeBool, value)
	}
	if value, ok := uu.mutation.IsAdmin(); ok {
		_spec.SetField(user.FieldIsAdmin, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetPhotoURL sets the "photo_url" field.
func (uuo *UserUpdateOne) SetPhotoURL(s string) *UserUpdateOne {
	uuo.mutation.SetPhotoURL(s)
	return uuo
}

// SetSessionToken sets the "session_token" field.
func (uuo *UserUpdateOne) SetSessionToken(s string) *UserUpdateOne {
	uuo.mutation.SetSessionToken(s)
	return uuo
}

// SetNillableSessionToken sets the "session_token" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSessionToken(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSessionToken(*s)
	}
	return uuo
}

// SetIsActivated sets the "is_activated" field.
func (uuo *UserUpdateOne) SetIsActivated(b bool) *UserUpdateOne {
	uuo.mutation.SetIsActivated(b)
	return uuo
}

// SetNillableIsActivated sets the "is_activated" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsActivated(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsActivated(*b)
	}
	return uuo
}

// SetIsAdmin sets the "is_admin" field.
func (uuo *UserUpdateOne) SetIsAdmin(b bool) *UserUpdateOne {
	uuo.mutation.SetIsAdmin(b)
	return uuo
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsAdmin(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsAdmin(*b)
	}
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.SessionToken(); ok {
		if err := user.SessionTokenValidator(v); err != nil {
			return &ValidationError{Name: "session_token", err: fmt.Errorf(`ent: validator failed for field "User.session_token": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.PhotoURL(); ok {
		_spec.SetField(user.FieldPhotoURL, field.TypeString, value)
	}
	if value, ok := uuo.mutation.SessionToken(); ok {
		_spec.SetField(user.FieldSessionToken, field.TypeString, value)
	}
	if value, ok := uuo.mutation.IsActivated(); ok {
		_spec.SetField(user.FieldIsActivated, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.IsAdmin(); ok {
		_spec.SetField(user.FieldIsAdmin, field.TypeBool, value)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
