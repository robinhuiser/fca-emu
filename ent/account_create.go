// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/robinhuiser/fca-emu/ent/account"
	"github.com/robinhuiser/fca-emu/ent/branch"
	"github.com/robinhuiser/fca-emu/ent/entity"
	"github.com/robinhuiser/fca-emu/ent/preference"
	"github.com/robinhuiser/fca-emu/ent/routingnumber"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (ac *AccountCreate) SetType(s string) *AccountCreate {
	ac.mutation.SetType(s)
	return ac
}

// SetNumber sets the "number" field.
func (ac *AccountCreate) SetNumber(s string) *AccountCreate {
	ac.mutation.SetNumber(s)
	return ac
}

// SetParentId sets the "parentId" field.
func (ac *AccountCreate) SetParentId(u uuid.UUID) *AccountCreate {
	ac.mutation.SetParentId(u)
	return ac
}

// SetName sets the "name" field.
func (ac *AccountCreate) SetName(s string) *AccountCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetTitle sets the "title" field.
func (ac *AccountCreate) SetTitle(s string) *AccountCreate {
	ac.mutation.SetTitle(s)
	return ac
}

// SetDateCreated sets the "dateCreated" field.
func (ac *AccountCreate) SetDateCreated(t time.Time) *AccountCreate {
	ac.mutation.SetDateCreated(t)
	return ac
}

// SetDateOpened sets the "dateOpened" field.
func (ac *AccountCreate) SetDateOpened(t time.Time) *AccountCreate {
	ac.mutation.SetDateOpened(t)
	return ac
}

// SetDateLastUpdated sets the "dateLastUpdated" field.
func (ac *AccountCreate) SetDateLastUpdated(t time.Time) *AccountCreate {
	ac.mutation.SetDateLastUpdated(t)
	return ac
}

// SetDateClosed sets the "dateClosed" field.
func (ac *AccountCreate) SetDateClosed(t time.Time) *AccountCreate {
	ac.mutation.SetDateClosed(t)
	return ac
}

// SetNillableDateClosed sets the "dateClosed" field if the given value is not nil.
func (ac *AccountCreate) SetNillableDateClosed(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetDateClosed(*t)
	}
	return ac
}

// SetCurrencyCode sets the "currencyCode" field.
func (ac *AccountCreate) SetCurrencyCode(s string) *AccountCreate {
	ac.mutation.SetCurrencyCode(s)
	return ac
}

// SetStatus sets the "status" field.
func (ac *AccountCreate) SetStatus(s string) *AccountCreate {
	ac.mutation.SetStatus(s)
	return ac
}

// SetSource sets the "source" field.
func (ac *AccountCreate) SetSource(s string) *AccountCreate {
	ac.mutation.SetSource(s)
	return ac
}

// SetInterestReporting sets the "interestReporting" field.
func (ac *AccountCreate) SetInterestReporting(b bool) *AccountCreate {
	ac.mutation.SetInterestReporting(b)
	return ac
}

// SetCurrentBalance sets the "currentBalance" field.
func (ac *AccountCreate) SetCurrentBalance(f float32) *AccountCreate {
	ac.mutation.SetCurrentBalance(f)
	return ac
}

// SetAvailableBalance sets the "availableBalance" field.
func (ac *AccountCreate) SetAvailableBalance(f float32) *AccountCreate {
	ac.mutation.SetAvailableBalance(f)
	return ac
}

// SetURL sets the "url" field.
func (ac *AccountCreate) SetURL(s string) *AccountCreate {
	ac.mutation.SetURL(s)
	return ac
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (ac *AccountCreate) SetNillableURL(s *string) *AccountCreate {
	if s != nil {
		ac.SetURL(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AccountCreate) SetID(u uuid.UUID) *AccountCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetBranchID sets the "branch" edge to the Branch entity by ID.
func (ac *AccountCreate) SetBranchID(id int) *AccountCreate {
	ac.mutation.SetBranchID(id)
	return ac
}

// SetNillableBranchID sets the "branch" edge to the Branch entity by ID if the given value is not nil.
func (ac *AccountCreate) SetNillableBranchID(id *int) *AccountCreate {
	if id != nil {
		ac = ac.SetBranchID(*id)
	}
	return ac
}

// SetBranch sets the "branch" edge to the Branch entity.
func (ac *AccountCreate) SetBranch(b *Branch) *AccountCreate {
	return ac.SetBranchID(b.ID)
}

// AddOwnerIDs adds the "owner" edge to the Entity entity by IDs.
func (ac *AccountCreate) AddOwnerIDs(ids ...uuid.UUID) *AccountCreate {
	ac.mutation.AddOwnerIDs(ids...)
	return ac
}

// AddOwner adds the "owner" edges to the Entity entity.
func (ac *AccountCreate) AddOwner(e ...*Entity) *AccountCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ac.AddOwnerIDs(ids...)
}

// AddPreferenceIDs adds the "preference" edge to the Preference entity by IDs.
func (ac *AccountCreate) AddPreferenceIDs(ids ...int) *AccountCreate {
	ac.mutation.AddPreferenceIDs(ids...)
	return ac
}

// AddPreference adds the "preference" edges to the Preference entity.
func (ac *AccountCreate) AddPreference(p ...*Preference) *AccountCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddPreferenceIDs(ids...)
}

// AddRoutingnumberIDs adds the "routingnumber" edge to the RoutingNumber entity by IDs.
func (ac *AccountCreate) AddRoutingnumberIDs(ids ...int) *AccountCreate {
	ac.mutation.AddRoutingnumberIDs(ids...)
	return ac
}

// AddRoutingnumber adds the "routingnumber" edges to the RoutingNumber entity.
func (ac *AccountCreate) AddRoutingnumber(r ...*RoutingNumber) *AccountCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ac.AddRoutingnumberIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.ID(); !ok {
		v := account.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	if _, ok := ac.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New("ent: missing required field \"number\"")}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := ac.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if _, ok := ac.mutation.DateCreated(); !ok {
		return &ValidationError{Name: "dateCreated", err: errors.New("ent: missing required field \"dateCreated\"")}
	}
	if _, ok := ac.mutation.DateOpened(); !ok {
		return &ValidationError{Name: "dateOpened", err: errors.New("ent: missing required field \"dateOpened\"")}
	}
	if _, ok := ac.mutation.DateLastUpdated(); !ok {
		return &ValidationError{Name: "dateLastUpdated", err: errors.New("ent: missing required field \"dateLastUpdated\"")}
	}
	if _, ok := ac.mutation.CurrencyCode(); !ok {
		return &ValidationError{Name: "currencyCode", err: errors.New("ent: missing required field \"currencyCode\"")}
	}
	if _, ok := ac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := ac.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New("ent: missing required field \"source\"")}
	}
	if _, ok := ac.mutation.InterestReporting(); !ok {
		return &ValidationError{Name: "interestReporting", err: errors.New("ent: missing required field \"interestReporting\"")}
	}
	if _, ok := ac.mutation.CurrentBalance(); !ok {
		return &ValidationError{Name: "currentBalance", err: errors.New("ent: missing required field \"currentBalance\"")}
	}
	if _, ok := ac.mutation.AvailableBalance(); !ok {
		return &ValidationError{Name: "availableBalance", err: errors.New("ent: missing required field \"availableBalance\"")}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: account.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldType,
		})
		_node.Type = value
	}
	if value, ok := ac.mutation.Number(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldNumber,
		})
		_node.Number = value
	}
	if value, ok := ac.mutation.ParentId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: account.FieldParentId,
		})
		_node.ParentId = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := ac.mutation.DateCreated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldDateCreated,
		})
		_node.DateCreated = value
	}
	if value, ok := ac.mutation.DateOpened(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldDateOpened,
		})
		_node.DateOpened = value
	}
	if value, ok := ac.mutation.DateLastUpdated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldDateLastUpdated,
		})
		_node.DateLastUpdated = value
	}
	if value, ok := ac.mutation.DateClosed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldDateClosed,
		})
		_node.DateClosed = value
	}
	if value, ok := ac.mutation.CurrencyCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldCurrencyCode,
		})
		_node.CurrencyCode = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := ac.mutation.Source(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldSource,
		})
		_node.Source = value
	}
	if value, ok := ac.mutation.InterestReporting(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldInterestReporting,
		})
		_node.InterestReporting = value
	}
	if value, ok := ac.mutation.CurrentBalance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: account.FieldCurrentBalance,
		})
		_node.CurrentBalance = value
	}
	if value, ok := ac.mutation.AvailableBalance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: account.FieldAvailableBalance,
		})
		_node.AvailableBalance = value
	}
	if value, ok := ac.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldURL,
		})
		_node.URL = value
	}
	if nodes := ac.mutation.BranchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   account.BranchTable,
			Columns: []string{account.BranchColumn},
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
		_node.account_branch = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.OwnerTable,
			Columns: account.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.PreferenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PreferenceTable,
			Columns: []string{account.PreferenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: preference.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.RoutingnumberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.RoutingnumberTable,
			Columns: []string{account.RoutingnumberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: routingnumber.FieldID,
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

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	builders []*AccountCreate
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
