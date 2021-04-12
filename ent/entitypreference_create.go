// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/finite-mock-server/ent/entitypreference"
)

// EntityPreferenceCreate is the builder for creating a EntityPreference entity.
type EntityPreferenceCreate struct {
	config
	mutation *EntityPreferenceMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (epc *EntityPreferenceCreate) SetName(s string) *EntityPreferenceCreate {
	epc.mutation.SetName(s)
	return epc
}

// SetValue sets the "value" field.
func (epc *EntityPreferenceCreate) SetValue(s string) *EntityPreferenceCreate {
	epc.mutation.SetValue(s)
	return epc
}

// Mutation returns the EntityPreferenceMutation object of the builder.
func (epc *EntityPreferenceCreate) Mutation() *EntityPreferenceMutation {
	return epc.mutation
}

// Save creates the EntityPreference in the database.
func (epc *EntityPreferenceCreate) Save(ctx context.Context) (*EntityPreference, error) {
	var (
		err  error
		node *EntityPreference
	)
	if len(epc.hooks) == 0 {
		if err = epc.check(); err != nil {
			return nil, err
		}
		node, err = epc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityPreferenceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = epc.check(); err != nil {
				return nil, err
			}
			epc.mutation = mutation
			node, err = epc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(epc.hooks) - 1; i >= 0; i-- {
			mut = epc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, epc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (epc *EntityPreferenceCreate) SaveX(ctx context.Context) *EntityPreference {
	v, err := epc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (epc *EntityPreferenceCreate) check() error {
	if _, ok := epc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := epc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New("ent: missing required field \"value\"")}
	}
	return nil
}

func (epc *EntityPreferenceCreate) sqlSave(ctx context.Context) (*EntityPreference, error) {
	_node, _spec := epc.createSpec()
	if err := sqlgraph.CreateNode(ctx, epc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (epc *EntityPreferenceCreate) createSpec() (*EntityPreference, *sqlgraph.CreateSpec) {
	var (
		_node = &EntityPreference{config: epc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: entitypreference.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entitypreference.FieldID,
			},
		}
	)
	if value, ok := epc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entitypreference.FieldName,
		})
		_node.Name = value
	}
	if value, ok := epc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entitypreference.FieldValue,
		})
		_node.Value = value
	}
	return _node, _spec
}

// EntityPreferenceCreateBulk is the builder for creating many EntityPreference entities in bulk.
type EntityPreferenceCreateBulk struct {
	config
	builders []*EntityPreferenceCreate
}

// Save creates the EntityPreference entities in the database.
func (epcb *EntityPreferenceCreateBulk) Save(ctx context.Context) ([]*EntityPreference, error) {
	specs := make([]*sqlgraph.CreateSpec, len(epcb.builders))
	nodes := make([]*EntityPreference, len(epcb.builders))
	mutators := make([]Mutator, len(epcb.builders))
	for i := range epcb.builders {
		func(i int, root context.Context) {
			builder := epcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntityPreferenceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, epcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, epcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, epcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (epcb *EntityPreferenceCreateBulk) SaveX(ctx context.Context) []*EntityPreference {
	v, err := epcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
