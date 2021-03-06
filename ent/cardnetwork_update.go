// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/fca-emu/ent/cardnetwork"
	"github.com/robinhuiser/fca-emu/ent/predicate"
)

// CardNetworkUpdate is the builder for updating CardNetwork entities.
type CardNetworkUpdate struct {
	config
	hooks    []Hook
	mutation *CardNetworkMutation
}

// Where adds a new predicate for the CardNetworkUpdate builder.
func (cnu *CardNetworkUpdate) Where(ps ...predicate.CardNetwork) *CardNetworkUpdate {
	cnu.mutation.predicates = append(cnu.mutation.predicates, ps...)
	return cnu
}

// SetName sets the "name" field.
func (cnu *CardNetworkUpdate) SetName(s string) *CardNetworkUpdate {
	cnu.mutation.SetName(s)
	return cnu
}

// SetCode sets the "code" field.
func (cnu *CardNetworkUpdate) SetCode(s string) *CardNetworkUpdate {
	cnu.mutation.SetCode(s)
	return cnu
}

// Mutation returns the CardNetworkMutation object of the builder.
func (cnu *CardNetworkUpdate) Mutation() *CardNetworkMutation {
	return cnu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cnu *CardNetworkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cnu.hooks) == 0 {
		affected, err = cnu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardNetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cnu.mutation = mutation
			affected, err = cnu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cnu.hooks) - 1; i >= 0; i-- {
			mut = cnu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cnu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cnu *CardNetworkUpdate) SaveX(ctx context.Context) int {
	affected, err := cnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cnu *CardNetworkUpdate) Exec(ctx context.Context) error {
	_, err := cnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnu *CardNetworkUpdate) ExecX(ctx context.Context) {
	if err := cnu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cnu *CardNetworkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardnetwork.Table,
			Columns: cardnetwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardnetwork.FieldID,
			},
		},
	}
	if ps := cnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cnu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cardnetwork.FieldName,
		})
	}
	if value, ok := cnu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cardnetwork.FieldCode,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cardnetwork.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CardNetworkUpdateOne is the builder for updating a single CardNetwork entity.
type CardNetworkUpdateOne struct {
	config
	hooks    []Hook
	mutation *CardNetworkMutation
}

// SetName sets the "name" field.
func (cnuo *CardNetworkUpdateOne) SetName(s string) *CardNetworkUpdateOne {
	cnuo.mutation.SetName(s)
	return cnuo
}

// SetCode sets the "code" field.
func (cnuo *CardNetworkUpdateOne) SetCode(s string) *CardNetworkUpdateOne {
	cnuo.mutation.SetCode(s)
	return cnuo
}

// Mutation returns the CardNetworkMutation object of the builder.
func (cnuo *CardNetworkUpdateOne) Mutation() *CardNetworkMutation {
	return cnuo.mutation
}

// Save executes the query and returns the updated CardNetwork entity.
func (cnuo *CardNetworkUpdateOne) Save(ctx context.Context) (*CardNetwork, error) {
	var (
		err  error
		node *CardNetwork
	)
	if len(cnuo.hooks) == 0 {
		node, err = cnuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardNetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cnuo.mutation = mutation
			node, err = cnuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cnuo.hooks) - 1; i >= 0; i-- {
			mut = cnuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cnuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cnuo *CardNetworkUpdateOne) SaveX(ctx context.Context) *CardNetwork {
	node, err := cnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cnuo *CardNetworkUpdateOne) Exec(ctx context.Context) error {
	_, err := cnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnuo *CardNetworkUpdateOne) ExecX(ctx context.Context) {
	if err := cnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cnuo *CardNetworkUpdateOne) sqlSave(ctx context.Context) (_node *CardNetwork, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardnetwork.Table,
			Columns: cardnetwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardnetwork.FieldID,
			},
		},
	}
	id, ok := cnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CardNetwork.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := cnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cnuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cardnetwork.FieldName,
		})
	}
	if value, ok := cnuo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cardnetwork.FieldCode,
		})
	}
	_node = &CardNetwork{config: cnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cardnetwork.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
