// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/finite-mock-server/ent/entitytaxinformation"
)

// EntityTaxInformationCreate is the builder for creating a EntityTaxInformation entity.
type EntityTaxInformationCreate struct {
	config
	mutation *EntityTaxInformationMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (etic *EntityTaxInformationCreate) SetType(e entitytaxinformation.Type) *EntityTaxInformationCreate {
	etic.mutation.SetType(e)
	return etic
}

// SetTaxId sets the "taxId" field.
func (etic *EntityTaxInformationCreate) SetTaxId(s string) *EntityTaxInformationCreate {
	etic.mutation.SetTaxId(s)
	return etic
}

// Mutation returns the EntityTaxInformationMutation object of the builder.
func (etic *EntityTaxInformationCreate) Mutation() *EntityTaxInformationMutation {
	return etic.mutation
}

// Save creates the EntityTaxInformation in the database.
func (etic *EntityTaxInformationCreate) Save(ctx context.Context) (*EntityTaxInformation, error) {
	var (
		err  error
		node *EntityTaxInformation
	)
	if len(etic.hooks) == 0 {
		if err = etic.check(); err != nil {
			return nil, err
		}
		node, err = etic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityTaxInformationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = etic.check(); err != nil {
				return nil, err
			}
			etic.mutation = mutation
			node, err = etic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(etic.hooks) - 1; i >= 0; i-- {
			mut = etic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, etic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (etic *EntityTaxInformationCreate) SaveX(ctx context.Context) *EntityTaxInformation {
	v, err := etic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (etic *EntityTaxInformationCreate) check() error {
	if _, ok := etic.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	if v, ok := etic.mutation.GetType(); ok {
		if err := entitytaxinformation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := etic.mutation.TaxId(); !ok {
		return &ValidationError{Name: "taxId", err: errors.New("ent: missing required field \"taxId\"")}
	}
	return nil
}

func (etic *EntityTaxInformationCreate) sqlSave(ctx context.Context) (*EntityTaxInformation, error) {
	_node, _spec := etic.createSpec()
	if err := sqlgraph.CreateNode(ctx, etic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (etic *EntityTaxInformationCreate) createSpec() (*EntityTaxInformation, *sqlgraph.CreateSpec) {
	var (
		_node = &EntityTaxInformation{config: etic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: entitytaxinformation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entitytaxinformation.FieldID,
			},
		}
	)
	if value, ok := etic.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: entitytaxinformation.FieldType,
		})
		_node.Type = value
	}
	if value, ok := etic.mutation.TaxId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entitytaxinformation.FieldTaxId,
		})
		_node.TaxId = value
	}
	return _node, _spec
}

// EntityTaxInformationCreateBulk is the builder for creating many EntityTaxInformation entities in bulk.
type EntityTaxInformationCreateBulk struct {
	config
	builders []*EntityTaxInformationCreate
}

// Save creates the EntityTaxInformation entities in the database.
func (eticb *EntityTaxInformationCreateBulk) Save(ctx context.Context) ([]*EntityTaxInformation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eticb.builders))
	nodes := make([]*EntityTaxInformation, len(eticb.builders))
	mutators := make([]Mutator, len(eticb.builders))
	for i := range eticb.builders {
		func(i int, root context.Context) {
			builder := eticb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntityTaxInformationMutation)
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
					_, err = mutators[i+1].Mutate(root, eticb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eticb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eticb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eticb *EntityTaxInformationCreateBulk) SaveX(ctx context.Context) []*EntityTaxInformation {
	v, err := eticb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
