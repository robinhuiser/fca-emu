// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/finite-mock-server/ent/entitypreference"
	"github.com/robinhuiser/finite-mock-server/ent/predicate"
)

// EntityPreferenceUpdate is the builder for updating EntityPreference entities.
type EntityPreferenceUpdate struct {
	config
	hooks    []Hook
	mutation *EntityPreferenceMutation
}

// Where adds a new predicate for the EntityPreferenceUpdate builder.
func (epu *EntityPreferenceUpdate) Where(ps ...predicate.EntityPreference) *EntityPreferenceUpdate {
	epu.mutation.predicates = append(epu.mutation.predicates, ps...)
	return epu
}

// SetName sets the "name" field.
func (epu *EntityPreferenceUpdate) SetName(s string) *EntityPreferenceUpdate {
	epu.mutation.SetName(s)
	return epu
}

// SetValue sets the "value" field.
func (epu *EntityPreferenceUpdate) SetValue(s string) *EntityPreferenceUpdate {
	epu.mutation.SetValue(s)
	return epu
}

// Mutation returns the EntityPreferenceMutation object of the builder.
func (epu *EntityPreferenceUpdate) Mutation() *EntityPreferenceMutation {
	return epu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (epu *EntityPreferenceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(epu.hooks) == 0 {
		affected, err = epu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityPreferenceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			epu.mutation = mutation
			affected, err = epu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(epu.hooks) - 1; i >= 0; i-- {
			mut = epu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, epu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (epu *EntityPreferenceUpdate) SaveX(ctx context.Context) int {
	affected, err := epu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (epu *EntityPreferenceUpdate) Exec(ctx context.Context) error {
	_, err := epu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epu *EntityPreferenceUpdate) ExecX(ctx context.Context) {
	if err := epu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (epu *EntityPreferenceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entitypreference.Table,
			Columns: entitypreference.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entitypreference.FieldID,
			},
		},
	}
	if ps := epu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := epu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entitypreference.FieldName,
		})
	}
	if value, ok := epu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entitypreference.FieldValue,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, epu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitypreference.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EntityPreferenceUpdateOne is the builder for updating a single EntityPreference entity.
type EntityPreferenceUpdateOne struct {
	config
	hooks    []Hook
	mutation *EntityPreferenceMutation
}

// SetName sets the "name" field.
func (epuo *EntityPreferenceUpdateOne) SetName(s string) *EntityPreferenceUpdateOne {
	epuo.mutation.SetName(s)
	return epuo
}

// SetValue sets the "value" field.
func (epuo *EntityPreferenceUpdateOne) SetValue(s string) *EntityPreferenceUpdateOne {
	epuo.mutation.SetValue(s)
	return epuo
}

// Mutation returns the EntityPreferenceMutation object of the builder.
func (epuo *EntityPreferenceUpdateOne) Mutation() *EntityPreferenceMutation {
	return epuo.mutation
}

// Save executes the query and returns the updated EntityPreference entity.
func (epuo *EntityPreferenceUpdateOne) Save(ctx context.Context) (*EntityPreference, error) {
	var (
		err  error
		node *EntityPreference
	)
	if len(epuo.hooks) == 0 {
		node, err = epuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityPreferenceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			epuo.mutation = mutation
			node, err = epuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(epuo.hooks) - 1; i >= 0; i-- {
			mut = epuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, epuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (epuo *EntityPreferenceUpdateOne) SaveX(ctx context.Context) *EntityPreference {
	node, err := epuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (epuo *EntityPreferenceUpdateOne) Exec(ctx context.Context) error {
	_, err := epuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epuo *EntityPreferenceUpdateOne) ExecX(ctx context.Context) {
	if err := epuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (epuo *EntityPreferenceUpdateOne) sqlSave(ctx context.Context) (_node *EntityPreference, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entitypreference.Table,
			Columns: entitypreference.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entitypreference.FieldID,
			},
		},
	}
	id, ok := epuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing EntityPreference.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := epuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := epuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entitypreference.FieldName,
		})
	}
	if value, ok := epuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entitypreference.FieldValue,
		})
	}
	_node = &EntityPreference{config: epuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, epuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitypreference.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
