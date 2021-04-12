// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/finite-mock-server/ent/product"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetType sets the "type" field.
func (pc *ProductCreate) SetType(pr product.Type) *ProductCreate {
	pc.mutation.SetType(pr)
	return pc
}

// SetTypeName sets the "typeName" field.
func (pc *ProductCreate) SetTypeName(s string) *ProductCreate {
	pc.mutation.SetTypeName(s)
	return pc
}

// SetSubType sets the "subType" field.
func (pc *ProductCreate) SetSubType(s string) *ProductCreate {
	pc.mutation.SetSubType(s)
	return pc
}

// SetSubTypeName sets the "subTypeName" field.
func (pc *ProductCreate) SetSubTypeName(s string) *ProductCreate {
	pc.mutation.SetSubTypeName(s)
	return pc
}

// SetURL sets the "url" field.
func (pc *ProductCreate) SetURL(s string) *ProductCreate {
	pc.mutation.SetURL(s)
	return pc
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	var (
		err  error
		node *Product
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := pc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	if v, ok := pc.mutation.GetType(); ok {
		if err := product.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := pc.mutation.TypeName(); !ok {
		return &ValidationError{Name: "typeName", err: errors.New("ent: missing required field \"typeName\"")}
	}
	if _, ok := pc.mutation.SubType(); !ok {
		return &ValidationError{Name: "subType", err: errors.New("ent: missing required field \"subType\"")}
	}
	if _, ok := pc.mutation.SubTypeName(); !ok {
		return &ValidationError{Name: "subTypeName", err: errors.New("ent: missing required field \"subTypeName\"")}
	}
	if _, ok := pc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New("ent: missing required field \"url\"")}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: product.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: product.FieldType,
		})
		_node.Type = value
	}
	if value, ok := pc.mutation.TypeName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldTypeName,
		})
		_node.TypeName = value
	}
	if value, ok := pc.mutation.SubType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldSubType,
		})
		_node.SubType = value
	}
	if value, ok := pc.mutation.SubTypeName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldSubTypeName,
		})
		_node.SubTypeName = value
	}
	if value, ok := pc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldURL,
		})
		_node.URL = value
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
