// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/fca-emu/ent/bank"
	"github.com/robinhuiser/fca-emu/ent/branch"
)

// BankCreate is the builder for creating a Bank entity.
type BankCreate struct {
	config
	mutation *BankMutation
	hooks    []Hook
}

// SetBankCode sets the "bankCode" field.
func (bc *BankCreate) SetBankCode(s string) *BankCreate {
	bc.mutation.SetBankCode(s)
	return bc
}

// SetBankName sets the "bankName" field.
func (bc *BankCreate) SetBankName(s string) *BankCreate {
	bc.mutation.SetBankName(s)
	return bc
}

// SetURL sets the "url" field.
func (bc *BankCreate) SetURL(s string) *BankCreate {
	bc.mutation.SetURL(s)
	return bc
}

// SetSwift sets the "swift" field.
func (bc *BankCreate) SetSwift(s string) *BankCreate {
	bc.mutation.SetSwift(s)
	return bc
}

// AddBranchIDs adds the "branches" edge to the Branch entity by IDs.
func (bc *BankCreate) AddBranchIDs(ids ...int) *BankCreate {
	bc.mutation.AddBranchIDs(ids...)
	return bc
}

// AddBranches adds the "branches" edges to the Branch entity.
func (bc *BankCreate) AddBranches(b ...*Branch) *BankCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddBranchIDs(ids...)
}

// Mutation returns the BankMutation object of the builder.
func (bc *BankCreate) Mutation() *BankMutation {
	return bc.mutation
}

// Save creates the Bank in the database.
func (bc *BankCreate) Save(ctx context.Context) (*Bank, error) {
	var (
		err  error
		node *Bank
	)
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BankMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BankCreate) SaveX(ctx context.Context) *Bank {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (bc *BankCreate) check() error {
	if _, ok := bc.mutation.BankCode(); !ok {
		return &ValidationError{Name: "bankCode", err: errors.New("ent: missing required field \"bankCode\"")}
	}
	if _, ok := bc.mutation.BankName(); !ok {
		return &ValidationError{Name: "bankName", err: errors.New("ent: missing required field \"bankName\"")}
	}
	if _, ok := bc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New("ent: missing required field \"url\"")}
	}
	if _, ok := bc.mutation.Swift(); !ok {
		return &ValidationError{Name: "swift", err: errors.New("ent: missing required field \"swift\"")}
	}
	return nil
}

func (bc *BankCreate) sqlSave(ctx context.Context) (*Bank, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BankCreate) createSpec() (*Bank, *sqlgraph.CreateSpec) {
	var (
		_node = &Bank{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bank.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bank.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.BankCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bank.FieldBankCode,
		})
		_node.BankCode = value
	}
	if value, ok := bc.mutation.BankName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bank.FieldBankName,
		})
		_node.BankName = value
	}
	if value, ok := bc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bank.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := bc.mutation.Swift(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bank.FieldSwift,
		})
		_node.Swift = value
	}
	if nodes := bc.mutation.BranchesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bank.BranchesTable,
			Columns: []string{bank.BranchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: branch.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BankCreateBulk is the builder for creating many Bank entities in bulk.
type BankCreateBulk struct {
	config
	builders []*BankCreate
}

// Save creates the Bank entities in the database.
func (bcb *BankCreateBulk) Save(ctx context.Context) ([]*Bank, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Bank, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BankMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BankCreateBulk) SaveX(ctx context.Context) []*Bank {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}