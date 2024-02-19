// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"realm.pub/tavern/internal/ent/repository"
	"realm.pub/tavern/internal/ent/tome"
	"realm.pub/tavern/internal/ent/user"
)

// RepositoryCreate is the builder for creating a Repository entity.
type RepositoryCreate struct {
	config
	mutation *RepositoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (rc *RepositoryCreate) SetCreatedAt(t time.Time) *RepositoryCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RepositoryCreate) SetNillableCreatedAt(t *time.Time) *RepositoryCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (rc *RepositoryCreate) SetLastModifiedAt(t time.Time) *RepositoryCreate {
	rc.mutation.SetLastModifiedAt(t)
	return rc
}

// SetNillableLastModifiedAt sets the "last_modified_at" field if the given value is not nil.
func (rc *RepositoryCreate) SetNillableLastModifiedAt(t *time.Time) *RepositoryCreate {
	if t != nil {
		rc.SetLastModifiedAt(*t)
	}
	return rc
}

// SetURL sets the "url" field.
func (rc *RepositoryCreate) SetURL(s string) *RepositoryCreate {
	rc.mutation.SetURL(s)
	return rc
}

// SetPublicKey sets the "public_key" field.
func (rc *RepositoryCreate) SetPublicKey(s string) *RepositoryCreate {
	rc.mutation.SetPublicKey(s)
	return rc
}

// SetPrivateKey sets the "private_key" field.
func (rc *RepositoryCreate) SetPrivateKey(s string) *RepositoryCreate {
	rc.mutation.SetPrivateKey(s)
	return rc
}

// AddTomeIDs adds the "tomes" edge to the Tome entity by IDs.
func (rc *RepositoryCreate) AddTomeIDs(ids ...int) *RepositoryCreate {
	rc.mutation.AddTomeIDs(ids...)
	return rc
}

// AddTomes adds the "tomes" edges to the Tome entity.
func (rc *RepositoryCreate) AddTomes(t ...*Tome) *RepositoryCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rc.AddTomeIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (rc *RepositoryCreate) SetOwnerID(id int) *RepositoryCreate {
	rc.mutation.SetOwnerID(id)
	return rc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (rc *RepositoryCreate) SetNillableOwnerID(id *int) *RepositoryCreate {
	if id != nil {
		rc = rc.SetOwnerID(*id)
	}
	return rc
}

// SetOwner sets the "owner" edge to the User entity.
func (rc *RepositoryCreate) SetOwner(u *User) *RepositoryCreate {
	return rc.SetOwnerID(u.ID)
}

// Mutation returns the RepositoryMutation object of the builder.
func (rc *RepositoryCreate) Mutation() *RepositoryMutation {
	return rc.mutation
}

// Save creates the Repository in the database.
func (rc *RepositoryCreate) Save(ctx context.Context) (*Repository, error) {
	if err := rc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RepositoryCreate) SaveX(ctx context.Context) *Repository {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RepositoryCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RepositoryCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RepositoryCreate) defaults() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		if repository.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized repository.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := repository.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.LastModifiedAt(); !ok {
		if repository.DefaultLastModifiedAt == nil {
			return fmt.Errorf("ent: uninitialized repository.DefaultLastModifiedAt (forgotten import ent/runtime?)")
		}
		v := repository.DefaultLastModifiedAt()
		rc.mutation.SetLastModifiedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rc *RepositoryCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Repository.created_at"`)}
	}
	if _, ok := rc.mutation.LastModifiedAt(); !ok {
		return &ValidationError{Name: "last_modified_at", err: errors.New(`ent: missing required field "Repository.last_modified_at"`)}
	}
	if _, ok := rc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Repository.url"`)}
	}
	if v, ok := rc.mutation.URL(); ok {
		if err := repository.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Repository.url": %w`, err)}
		}
	}
	if _, ok := rc.mutation.PublicKey(); !ok {
		return &ValidationError{Name: "public_key", err: errors.New(`ent: missing required field "Repository.public_key"`)}
	}
	if v, ok := rc.mutation.PublicKey(); ok {
		if err := repository.PublicKeyValidator(v); err != nil {
			return &ValidationError{Name: "public_key", err: fmt.Errorf(`ent: validator failed for field "Repository.public_key": %w`, err)}
		}
	}
	if _, ok := rc.mutation.PrivateKey(); !ok {
		return &ValidationError{Name: "private_key", err: errors.New(`ent: missing required field "Repository.private_key"`)}
	}
	if v, ok := rc.mutation.PrivateKey(); ok {
		if err := repository.PrivateKeyValidator(v); err != nil {
			return &ValidationError{Name: "private_key", err: fmt.Errorf(`ent: validator failed for field "Repository.private_key": %w`, err)}
		}
	}
	return nil
}

func (rc *RepositoryCreate) sqlSave(ctx context.Context) (*Repository, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RepositoryCreate) createSpec() (*Repository, *sqlgraph.CreateSpec) {
	var (
		_node = &Repository{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(repository.Table, sqlgraph.NewFieldSpec(repository.FieldID, field.TypeInt))
	)
	_spec.OnConflict = rc.conflict
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(repository.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.LastModifiedAt(); ok {
		_spec.SetField(repository.FieldLastModifiedAt, field.TypeTime, value)
		_node.LastModifiedAt = value
	}
	if value, ok := rc.mutation.URL(); ok {
		_spec.SetField(repository.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := rc.mutation.PublicKey(); ok {
		_spec.SetField(repository.FieldPublicKey, field.TypeString, value)
		_node.PublicKey = value
	}
	if value, ok := rc.mutation.PrivateKey(); ok {
		_spec.SetField(repository.FieldPrivateKey, field.TypeString, value)
		_node.PrivateKey = value
	}
	if nodes := rc.mutation.TomesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   repository.TomesTable,
			Columns: []string{repository.TomesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tome.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   repository.OwnerTable,
			Columns: []string{repository.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.repository_owner = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Repository.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RepositoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rc *RepositoryCreate) OnConflict(opts ...sql.ConflictOption) *RepositoryUpsertOne {
	rc.conflict = opts
	return &RepositoryUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Repository.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RepositoryCreate) OnConflictColumns(columns ...string) *RepositoryUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RepositoryUpsertOne{
		create: rc,
	}
}

type (
	// RepositoryUpsertOne is the builder for "upsert"-ing
	//  one Repository node.
	RepositoryUpsertOne struct {
		create *RepositoryCreate
	}

	// RepositoryUpsert is the "OnConflict" setter.
	RepositoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *RepositoryUpsert) SetLastModifiedAt(v time.Time) *RepositoryUpsert {
	u.Set(repository.FieldLastModifiedAt, v)
	return u
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *RepositoryUpsert) UpdateLastModifiedAt() *RepositoryUpsert {
	u.SetExcluded(repository.FieldLastModifiedAt)
	return u
}

// SetURL sets the "url" field.
func (u *RepositoryUpsert) SetURL(v string) *RepositoryUpsert {
	u.Set(repository.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *RepositoryUpsert) UpdateURL() *RepositoryUpsert {
	u.SetExcluded(repository.FieldURL)
	return u
}

// SetPublicKey sets the "public_key" field.
func (u *RepositoryUpsert) SetPublicKey(v string) *RepositoryUpsert {
	u.Set(repository.FieldPublicKey, v)
	return u
}

// UpdatePublicKey sets the "public_key" field to the value that was provided on create.
func (u *RepositoryUpsert) UpdatePublicKey() *RepositoryUpsert {
	u.SetExcluded(repository.FieldPublicKey)
	return u
}

// SetPrivateKey sets the "private_key" field.
func (u *RepositoryUpsert) SetPrivateKey(v string) *RepositoryUpsert {
	u.Set(repository.FieldPrivateKey, v)
	return u
}

// UpdatePrivateKey sets the "private_key" field to the value that was provided on create.
func (u *RepositoryUpsert) UpdatePrivateKey() *RepositoryUpsert {
	u.SetExcluded(repository.FieldPrivateKey)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Repository.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RepositoryUpsertOne) UpdateNewValues() *RepositoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(repository.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Repository.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RepositoryUpsertOne) Ignore() *RepositoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RepositoryUpsertOne) DoNothing() *RepositoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RepositoryCreate.OnConflict
// documentation for more info.
func (u *RepositoryUpsertOne) Update(set func(*RepositoryUpsert)) *RepositoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RepositoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *RepositoryUpsertOne) SetLastModifiedAt(v time.Time) *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetLastModifiedAt(v)
	})
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *RepositoryUpsertOne) UpdateLastModifiedAt() *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateLastModifiedAt()
	})
}

// SetURL sets the "url" field.
func (u *RepositoryUpsertOne) SetURL(v string) *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *RepositoryUpsertOne) UpdateURL() *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateURL()
	})
}

// SetPublicKey sets the "public_key" field.
func (u *RepositoryUpsertOne) SetPublicKey(v string) *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetPublicKey(v)
	})
}

// UpdatePublicKey sets the "public_key" field to the value that was provided on create.
func (u *RepositoryUpsertOne) UpdatePublicKey() *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdatePublicKey()
	})
}

// SetPrivateKey sets the "private_key" field.
func (u *RepositoryUpsertOne) SetPrivateKey(v string) *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetPrivateKey(v)
	})
}

// UpdatePrivateKey sets the "private_key" field to the value that was provided on create.
func (u *RepositoryUpsertOne) UpdatePrivateKey() *RepositoryUpsertOne {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdatePrivateKey()
	})
}

// Exec executes the query.
func (u *RepositoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RepositoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RepositoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RepositoryUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RepositoryUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RepositoryCreateBulk is the builder for creating many Repository entities in bulk.
type RepositoryCreateBulk struct {
	config
	err      error
	builders []*RepositoryCreate
	conflict []sql.ConflictOption
}

// Save creates the Repository entities in the database.
func (rcb *RepositoryCreateBulk) Save(ctx context.Context) ([]*Repository, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Repository, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RepositoryMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RepositoryCreateBulk) SaveX(ctx context.Context) []*Repository {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RepositoryCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RepositoryCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Repository.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RepositoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rcb *RepositoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *RepositoryUpsertBulk {
	rcb.conflict = opts
	return &RepositoryUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Repository.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RepositoryCreateBulk) OnConflictColumns(columns ...string) *RepositoryUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RepositoryUpsertBulk{
		create: rcb,
	}
}

// RepositoryUpsertBulk is the builder for "upsert"-ing
// a bulk of Repository nodes.
type RepositoryUpsertBulk struct {
	create *RepositoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Repository.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RepositoryUpsertBulk) UpdateNewValues() *RepositoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(repository.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Repository.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RepositoryUpsertBulk) Ignore() *RepositoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RepositoryUpsertBulk) DoNothing() *RepositoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RepositoryCreateBulk.OnConflict
// documentation for more info.
func (u *RepositoryUpsertBulk) Update(set func(*RepositoryUpsert)) *RepositoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RepositoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *RepositoryUpsertBulk) SetLastModifiedAt(v time.Time) *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetLastModifiedAt(v)
	})
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *RepositoryUpsertBulk) UpdateLastModifiedAt() *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateLastModifiedAt()
	})
}

// SetURL sets the "url" field.
func (u *RepositoryUpsertBulk) SetURL(v string) *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *RepositoryUpsertBulk) UpdateURL() *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdateURL()
	})
}

// SetPublicKey sets the "public_key" field.
func (u *RepositoryUpsertBulk) SetPublicKey(v string) *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetPublicKey(v)
	})
}

// UpdatePublicKey sets the "public_key" field to the value that was provided on create.
func (u *RepositoryUpsertBulk) UpdatePublicKey() *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdatePublicKey()
	})
}

// SetPrivateKey sets the "private_key" field.
func (u *RepositoryUpsertBulk) SetPrivateKey(v string) *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.SetPrivateKey(v)
	})
}

// UpdatePrivateKey sets the "private_key" field to the value that was provided on create.
func (u *RepositoryUpsertBulk) UpdatePrivateKey() *RepositoryUpsertBulk {
	return u.Update(func(s *RepositoryUpsert) {
		s.UpdatePrivateKey()
	})
}

// Exec executes the query.
func (u *RepositoryUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RepositoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RepositoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RepositoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}