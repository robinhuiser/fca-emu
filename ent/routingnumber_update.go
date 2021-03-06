// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/fca-emu/ent/predicate"
	"github.com/robinhuiser/fca-emu/ent/routingnumber"
)

// RoutingNumberUpdate is the builder for updating RoutingNumber entities.
type RoutingNumberUpdate struct {
	config
	hooks    []Hook
	mutation *RoutingNumberMutation
}

// Where adds a new predicate for the RoutingNumberUpdate builder.
func (rnu *RoutingNumberUpdate) Where(ps ...predicate.RoutingNumber) *RoutingNumberUpdate {
	rnu.mutation.predicates = append(rnu.mutation.predicates, ps...)
	return rnu
}

// SetNumber sets the "number" field.
func (rnu *RoutingNumberUpdate) SetNumber(s string) *RoutingNumberUpdate {
	rnu.mutation.SetNumber(s)
	return rnu
}

// SetType sets the "type" field.
func (rnu *RoutingNumberUpdate) SetType(r routingnumber.Type) *RoutingNumberUpdate {
	rnu.mutation.SetType(r)
	return rnu
}

// Mutation returns the RoutingNumberMutation object of the builder.
func (rnu *RoutingNumberUpdate) Mutation() *RoutingNumberMutation {
	return rnu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rnu *RoutingNumberUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rnu.hooks) == 0 {
		if err = rnu.check(); err != nil {
			return 0, err
		}
		affected, err = rnu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoutingNumberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rnu.check(); err != nil {
				return 0, err
			}
			rnu.mutation = mutation
			affected, err = rnu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rnu.hooks) - 1; i >= 0; i-- {
			mut = rnu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rnu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rnu *RoutingNumberUpdate) SaveX(ctx context.Context) int {
	affected, err := rnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rnu *RoutingNumberUpdate) Exec(ctx context.Context) error {
	_, err := rnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rnu *RoutingNumberUpdate) ExecX(ctx context.Context) {
	if err := rnu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rnu *RoutingNumberUpdate) check() error {
	if v, ok := rnu.mutation.GetType(); ok {
		if err := routingnumber.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (rnu *RoutingNumberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   routingnumber.Table,
			Columns: routingnumber.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: routingnumber.FieldID,
			},
		},
	}
	if ps := rnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rnu.mutation.Number(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: routingnumber.FieldNumber,
		})
	}
	if value, ok := rnu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: routingnumber.FieldType,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routingnumber.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoutingNumberUpdateOne is the builder for updating a single RoutingNumber entity.
type RoutingNumberUpdateOne struct {
	config
	hooks    []Hook
	mutation *RoutingNumberMutation
}

// SetNumber sets the "number" field.
func (rnuo *RoutingNumberUpdateOne) SetNumber(s string) *RoutingNumberUpdateOne {
	rnuo.mutation.SetNumber(s)
	return rnuo
}

// SetType sets the "type" field.
func (rnuo *RoutingNumberUpdateOne) SetType(r routingnumber.Type) *RoutingNumberUpdateOne {
	rnuo.mutation.SetType(r)
	return rnuo
}

// Mutation returns the RoutingNumberMutation object of the builder.
func (rnuo *RoutingNumberUpdateOne) Mutation() *RoutingNumberMutation {
	return rnuo.mutation
}

// Save executes the query and returns the updated RoutingNumber entity.
func (rnuo *RoutingNumberUpdateOne) Save(ctx context.Context) (*RoutingNumber, error) {
	var (
		err  error
		node *RoutingNumber
	)
	if len(rnuo.hooks) == 0 {
		if err = rnuo.check(); err != nil {
			return nil, err
		}
		node, err = rnuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoutingNumberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rnuo.check(); err != nil {
				return nil, err
			}
			rnuo.mutation = mutation
			node, err = rnuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rnuo.hooks) - 1; i >= 0; i-- {
			mut = rnuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rnuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rnuo *RoutingNumberUpdateOne) SaveX(ctx context.Context) *RoutingNumber {
	node, err := rnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rnuo *RoutingNumberUpdateOne) Exec(ctx context.Context) error {
	_, err := rnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rnuo *RoutingNumberUpdateOne) ExecX(ctx context.Context) {
	if err := rnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rnuo *RoutingNumberUpdateOne) check() error {
	if v, ok := rnuo.mutation.GetType(); ok {
		if err := routingnumber.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (rnuo *RoutingNumberUpdateOne) sqlSave(ctx context.Context) (_node *RoutingNumber, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   routingnumber.Table,
			Columns: routingnumber.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: routingnumber.FieldID,
			},
		},
	}
	id, ok := rnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing RoutingNumber.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := rnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rnuo.mutation.Number(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: routingnumber.FieldNumber,
		})
	}
	if value, ok := rnuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: routingnumber.FieldType,
		})
	}
	_node = &RoutingNumber{config: rnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routingnumber.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
