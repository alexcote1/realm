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
	"realm.pub/tavern/internal/ent/beacon"
	"realm.pub/tavern/internal/ent/quest"
	"realm.pub/tavern/internal/ent/task"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (tc *TaskCreate) SetLastModifiedAt(t time.Time) *TaskCreate {
	tc.mutation.SetLastModifiedAt(t)
	return tc
}

// SetNillableLastModifiedAt sets the "last_modified_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableLastModifiedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetLastModifiedAt(*t)
	}
	return tc
}

// SetClaimedAt sets the "claimed_at" field.
func (tc *TaskCreate) SetClaimedAt(t time.Time) *TaskCreate {
	tc.mutation.SetClaimedAt(t)
	return tc
}

// SetNillableClaimedAt sets the "claimed_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableClaimedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetClaimedAt(*t)
	}
	return tc
}

// SetExecStartedAt sets the "exec_started_at" field.
func (tc *TaskCreate) SetExecStartedAt(t time.Time) *TaskCreate {
	tc.mutation.SetExecStartedAt(t)
	return tc
}

// SetNillableExecStartedAt sets the "exec_started_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableExecStartedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetExecStartedAt(*t)
	}
	return tc
}

// SetExecFinishedAt sets the "exec_finished_at" field.
func (tc *TaskCreate) SetExecFinishedAt(t time.Time) *TaskCreate {
	tc.mutation.SetExecFinishedAt(t)
	return tc
}

// SetNillableExecFinishedAt sets the "exec_finished_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableExecFinishedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetExecFinishedAt(*t)
	}
	return tc
}

// SetOutput sets the "output" field.
func (tc *TaskCreate) SetOutput(s string) *TaskCreate {
	tc.mutation.SetOutput(s)
	return tc
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (tc *TaskCreate) SetNillableOutput(s *string) *TaskCreate {
	if s != nil {
		tc.SetOutput(*s)
	}
	return tc
}

// SetError sets the "error" field.
func (tc *TaskCreate) SetError(s string) *TaskCreate {
	tc.mutation.SetError(s)
	return tc
}

// SetNillableError sets the "error" field if the given value is not nil.
func (tc *TaskCreate) SetNillableError(s *string) *TaskCreate {
	if s != nil {
		tc.SetError(*s)
	}
	return tc
}

// SetQuestID sets the "quest" edge to the Quest entity by ID.
func (tc *TaskCreate) SetQuestID(id int) *TaskCreate {
	tc.mutation.SetQuestID(id)
	return tc
}

// SetQuest sets the "quest" edge to the Quest entity.
func (tc *TaskCreate) SetQuest(q *Quest) *TaskCreate {
	return tc.SetQuestID(q.ID)
}

// SetBeaconID sets the "beacon" edge to the Beacon entity by ID.
func (tc *TaskCreate) SetBeaconID(id int) *TaskCreate {
	tc.mutation.SetBeaconID(id)
	return tc
}

// SetBeacon sets the "beacon" edge to the Beacon entity.
func (tc *TaskCreate) SetBeacon(b *Beacon) *TaskCreate {
	return tc.SetBeaconID(b.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := task.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.LastModifiedAt(); !ok {
		v := task.DefaultLastModifiedAt()
		tc.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Task.created_at"`)}
	}
	if _, ok := tc.mutation.LastModifiedAt(); !ok {
		return &ValidationError{Name: "last_modified_at", err: errors.New(`ent: missing required field "Task.last_modified_at"`)}
	}
	if _, ok := tc.mutation.QuestID(); !ok {
		return &ValidationError{Name: "quest", err: errors.New(`ent: missing required edge "Task.quest"`)}
	}
	if _, ok := tc.mutation.BeaconID(); !ok {
		return &ValidationError{Name: "beacon", err: errors.New(`ent: missing required edge "Task.beacon"`)}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(task.Table, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.LastModifiedAt(); ok {
		_spec.SetField(task.FieldLastModifiedAt, field.TypeTime, value)
		_node.LastModifiedAt = value
	}
	if value, ok := tc.mutation.ClaimedAt(); ok {
		_spec.SetField(task.FieldClaimedAt, field.TypeTime, value)
		_node.ClaimedAt = value
	}
	if value, ok := tc.mutation.ExecStartedAt(); ok {
		_spec.SetField(task.FieldExecStartedAt, field.TypeTime, value)
		_node.ExecStartedAt = value
	}
	if value, ok := tc.mutation.ExecFinishedAt(); ok {
		_spec.SetField(task.FieldExecFinishedAt, field.TypeTime, value)
		_node.ExecFinishedAt = value
	}
	if value, ok := tc.mutation.Output(); ok {
		_spec.SetField(task.FieldOutput, field.TypeString, value)
		_node.Output = value
	}
	if value, ok := tc.mutation.Error(); ok {
		_spec.SetField(task.FieldError, field.TypeString, value)
		_node.Error = value
	}
	if nodes := tc.mutation.QuestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.QuestTable,
			Columns: []string{task.QuestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(quest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.quest_tasks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.BeaconIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   task.BeaconTable,
			Columns: []string{task.BeaconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(beacon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_beacon = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Task.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tc *TaskCreate) OnConflict(opts ...sql.ConflictOption) *TaskUpsertOne {
	tc.conflict = opts
	return &TaskUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TaskCreate) OnConflictColumns(columns ...string) *TaskUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TaskUpsertOne{
		create: tc,
	}
}

type (
	// TaskUpsertOne is the builder for "upsert"-ing
	//  one Task node.
	TaskUpsertOne struct {
		create *TaskCreate
	}

	// TaskUpsert is the "OnConflict" setter.
	TaskUpsert struct {
		*sql.UpdateSet
	}
)

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *TaskUpsert) SetLastModifiedAt(v time.Time) *TaskUpsert {
	u.Set(task.FieldLastModifiedAt, v)
	return u
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *TaskUpsert) UpdateLastModifiedAt() *TaskUpsert {
	u.SetExcluded(task.FieldLastModifiedAt)
	return u
}

// SetClaimedAt sets the "claimed_at" field.
func (u *TaskUpsert) SetClaimedAt(v time.Time) *TaskUpsert {
	u.Set(task.FieldClaimedAt, v)
	return u
}

// UpdateClaimedAt sets the "claimed_at" field to the value that was provided on create.
func (u *TaskUpsert) UpdateClaimedAt() *TaskUpsert {
	u.SetExcluded(task.FieldClaimedAt)
	return u
}

// ClearClaimedAt clears the value of the "claimed_at" field.
func (u *TaskUpsert) ClearClaimedAt() *TaskUpsert {
	u.SetNull(task.FieldClaimedAt)
	return u
}

// SetExecStartedAt sets the "exec_started_at" field.
func (u *TaskUpsert) SetExecStartedAt(v time.Time) *TaskUpsert {
	u.Set(task.FieldExecStartedAt, v)
	return u
}

// UpdateExecStartedAt sets the "exec_started_at" field to the value that was provided on create.
func (u *TaskUpsert) UpdateExecStartedAt() *TaskUpsert {
	u.SetExcluded(task.FieldExecStartedAt)
	return u
}

// ClearExecStartedAt clears the value of the "exec_started_at" field.
func (u *TaskUpsert) ClearExecStartedAt() *TaskUpsert {
	u.SetNull(task.FieldExecStartedAt)
	return u
}

// SetExecFinishedAt sets the "exec_finished_at" field.
func (u *TaskUpsert) SetExecFinishedAt(v time.Time) *TaskUpsert {
	u.Set(task.FieldExecFinishedAt, v)
	return u
}

// UpdateExecFinishedAt sets the "exec_finished_at" field to the value that was provided on create.
func (u *TaskUpsert) UpdateExecFinishedAt() *TaskUpsert {
	u.SetExcluded(task.FieldExecFinishedAt)
	return u
}

// ClearExecFinishedAt clears the value of the "exec_finished_at" field.
func (u *TaskUpsert) ClearExecFinishedAt() *TaskUpsert {
	u.SetNull(task.FieldExecFinishedAt)
	return u
}

// SetOutput sets the "output" field.
func (u *TaskUpsert) SetOutput(v string) *TaskUpsert {
	u.Set(task.FieldOutput, v)
	return u
}

// UpdateOutput sets the "output" field to the value that was provided on create.
func (u *TaskUpsert) UpdateOutput() *TaskUpsert {
	u.SetExcluded(task.FieldOutput)
	return u
}

// ClearOutput clears the value of the "output" field.
func (u *TaskUpsert) ClearOutput() *TaskUpsert {
	u.SetNull(task.FieldOutput)
	return u
}

// SetError sets the "error" field.
func (u *TaskUpsert) SetError(v string) *TaskUpsert {
	u.Set(task.FieldError, v)
	return u
}

// UpdateError sets the "error" field to the value that was provided on create.
func (u *TaskUpsert) UpdateError() *TaskUpsert {
	u.SetExcluded(task.FieldError)
	return u
}

// ClearError clears the value of the "error" field.
func (u *TaskUpsert) ClearError() *TaskUpsert {
	u.SetNull(task.FieldError)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TaskUpsertOne) UpdateNewValues() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(task.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TaskUpsertOne) Ignore() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskUpsertOne) DoNothing() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCreate.OnConflict
// documentation for more info.
func (u *TaskUpsertOne) Update(set func(*TaskUpsert)) *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *TaskUpsertOne) SetLastModifiedAt(v time.Time) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetLastModifiedAt(v)
	})
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateLastModifiedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateLastModifiedAt()
	})
}

// SetClaimedAt sets the "claimed_at" field.
func (u *TaskUpsertOne) SetClaimedAt(v time.Time) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetClaimedAt(v)
	})
}

// UpdateClaimedAt sets the "claimed_at" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateClaimedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateClaimedAt()
	})
}

// ClearClaimedAt clears the value of the "claimed_at" field.
func (u *TaskUpsertOne) ClearClaimedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearClaimedAt()
	})
}

// SetExecStartedAt sets the "exec_started_at" field.
func (u *TaskUpsertOne) SetExecStartedAt(v time.Time) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetExecStartedAt(v)
	})
}

// UpdateExecStartedAt sets the "exec_started_at" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateExecStartedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateExecStartedAt()
	})
}

// ClearExecStartedAt clears the value of the "exec_started_at" field.
func (u *TaskUpsertOne) ClearExecStartedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearExecStartedAt()
	})
}

// SetExecFinishedAt sets the "exec_finished_at" field.
func (u *TaskUpsertOne) SetExecFinishedAt(v time.Time) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetExecFinishedAt(v)
	})
}

// UpdateExecFinishedAt sets the "exec_finished_at" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateExecFinishedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateExecFinishedAt()
	})
}

// ClearExecFinishedAt clears the value of the "exec_finished_at" field.
func (u *TaskUpsertOne) ClearExecFinishedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearExecFinishedAt()
	})
}

// SetOutput sets the "output" field.
func (u *TaskUpsertOne) SetOutput(v string) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetOutput(v)
	})
}

// UpdateOutput sets the "output" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateOutput() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateOutput()
	})
}

// ClearOutput clears the value of the "output" field.
func (u *TaskUpsertOne) ClearOutput() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearOutput()
	})
}

// SetError sets the "error" field.
func (u *TaskUpsertOne) SetError(v string) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetError(v)
	})
}

// UpdateError sets the "error" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateError() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateError()
	})
}

// ClearError clears the value of the "error" field.
func (u *TaskUpsertOne) ClearError() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearError()
	})
}

// Exec executes the query.
func (u *TaskUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TaskUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TaskUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	err      error
	builders []*TaskCreate
	conflict []sql.ConflictOption
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Task.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tcb *TaskCreateBulk) OnConflict(opts ...sql.ConflictOption) *TaskUpsertBulk {
	tcb.conflict = opts
	return &TaskUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TaskCreateBulk) OnConflictColumns(columns ...string) *TaskUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TaskUpsertBulk{
		create: tcb,
	}
}

// TaskUpsertBulk is the builder for "upsert"-ing
// a bulk of Task nodes.
type TaskUpsertBulk struct {
	create *TaskCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TaskUpsertBulk) UpdateNewValues() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(task.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TaskUpsertBulk) Ignore() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskUpsertBulk) DoNothing() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCreateBulk.OnConflict
// documentation for more info.
func (u *TaskUpsertBulk) Update(set func(*TaskUpsert)) *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *TaskUpsertBulk) SetLastModifiedAt(v time.Time) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetLastModifiedAt(v)
	})
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateLastModifiedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateLastModifiedAt()
	})
}

// SetClaimedAt sets the "claimed_at" field.
func (u *TaskUpsertBulk) SetClaimedAt(v time.Time) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetClaimedAt(v)
	})
}

// UpdateClaimedAt sets the "claimed_at" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateClaimedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateClaimedAt()
	})
}

// ClearClaimedAt clears the value of the "claimed_at" field.
func (u *TaskUpsertBulk) ClearClaimedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearClaimedAt()
	})
}

// SetExecStartedAt sets the "exec_started_at" field.
func (u *TaskUpsertBulk) SetExecStartedAt(v time.Time) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetExecStartedAt(v)
	})
}

// UpdateExecStartedAt sets the "exec_started_at" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateExecStartedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateExecStartedAt()
	})
}

// ClearExecStartedAt clears the value of the "exec_started_at" field.
func (u *TaskUpsertBulk) ClearExecStartedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearExecStartedAt()
	})
}

// SetExecFinishedAt sets the "exec_finished_at" field.
func (u *TaskUpsertBulk) SetExecFinishedAt(v time.Time) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetExecFinishedAt(v)
	})
}

// UpdateExecFinishedAt sets the "exec_finished_at" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateExecFinishedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateExecFinishedAt()
	})
}

// ClearExecFinishedAt clears the value of the "exec_finished_at" field.
func (u *TaskUpsertBulk) ClearExecFinishedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearExecFinishedAt()
	})
}

// SetOutput sets the "output" field.
func (u *TaskUpsertBulk) SetOutput(v string) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetOutput(v)
	})
}

// UpdateOutput sets the "output" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateOutput() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateOutput()
	})
}

// ClearOutput clears the value of the "output" field.
func (u *TaskUpsertBulk) ClearOutput() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearOutput()
	})
}

// SetError sets the "error" field.
func (u *TaskUpsertBulk) SetError(v string) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetError(v)
	})
}

// UpdateError sets the "error" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateError() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateError()
	})
}

// ClearError clears the value of the "error" field.
func (u *TaskUpsertBulk) ClearError() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearError()
	})
}

// Exec executes the query.
func (u *TaskUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TaskCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
