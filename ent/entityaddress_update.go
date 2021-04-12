// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/finite-mock-server/ent/entityaddress"
	"github.com/robinhuiser/finite-mock-server/ent/predicate"
)

// EntityAddressUpdate is the builder for updating EntityAddress entities.
type EntityAddressUpdate struct {
	config
	hooks    []Hook
	mutation *EntityAddressMutation
}

// Where adds a new predicate for the EntityAddressUpdate builder.
func (eau *EntityAddressUpdate) Where(ps ...predicate.EntityAddress) *EntityAddressUpdate {
	eau.mutation.predicates = append(eau.mutation.predicates, ps...)
	return eau
}

// SetCountry sets the "country" field.
func (eau *EntityAddressUpdate) SetCountry(s string) *EntityAddressUpdate {
	eau.mutation.SetCountry(s)
	return eau
}

// SetCity sets the "city" field.
func (eau *EntityAddressUpdate) SetCity(s string) *EntityAddressUpdate {
	eau.mutation.SetCity(s)
	return eau
}

// SetPostalCode sets the "postalCode" field.
func (eau *EntityAddressUpdate) SetPostalCode(s string) *EntityAddressUpdate {
	eau.mutation.SetPostalCode(s)
	return eau
}

// SetState sets the "state" field.
func (eau *EntityAddressUpdate) SetState(s string) *EntityAddressUpdate {
	eau.mutation.SetState(s)
	return eau
}

// SetType sets the "type" field.
func (eau *EntityAddressUpdate) SetType(e entityaddress.Type) *EntityAddressUpdate {
	eau.mutation.SetType(e)
	return eau
}

// SetLine1 sets the "line1" field.
func (eau *EntityAddressUpdate) SetLine1(s string) *EntityAddressUpdate {
	eau.mutation.SetLine1(s)
	return eau
}

// SetLine2 sets the "line2" field.
func (eau *EntityAddressUpdate) SetLine2(s string) *EntityAddressUpdate {
	eau.mutation.SetLine2(s)
	return eau
}

// SetNillableLine2 sets the "line2" field if the given value is not nil.
func (eau *EntityAddressUpdate) SetNillableLine2(s *string) *EntityAddressUpdate {
	if s != nil {
		eau.SetLine2(*s)
	}
	return eau
}

// ClearLine2 clears the value of the "line2" field.
func (eau *EntityAddressUpdate) ClearLine2() *EntityAddressUpdate {
	eau.mutation.ClearLine2()
	return eau
}

// SetLine3 sets the "line3" field.
func (eau *EntityAddressUpdate) SetLine3(s string) *EntityAddressUpdate {
	eau.mutation.SetLine3(s)
	return eau
}

// SetNillableLine3 sets the "line3" field if the given value is not nil.
func (eau *EntityAddressUpdate) SetNillableLine3(s *string) *EntityAddressUpdate {
	if s != nil {
		eau.SetLine3(*s)
	}
	return eau
}

// ClearLine3 clears the value of the "line3" field.
func (eau *EntityAddressUpdate) ClearLine3() *EntityAddressUpdate {
	eau.mutation.ClearLine3()
	return eau
}

// SetPrimary sets the "primary" field.
func (eau *EntityAddressUpdate) SetPrimary(b bool) *EntityAddressUpdate {
	eau.mutation.SetPrimary(b)
	return eau
}

// Mutation returns the EntityAddressMutation object of the builder.
func (eau *EntityAddressUpdate) Mutation() *EntityAddressMutation {
	return eau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eau *EntityAddressUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eau.hooks) == 0 {
		if err = eau.check(); err != nil {
			return 0, err
		}
		affected, err = eau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityAddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eau.check(); err != nil {
				return 0, err
			}
			eau.mutation = mutation
			affected, err = eau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eau.hooks) - 1; i >= 0; i-- {
			mut = eau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eau *EntityAddressUpdate) SaveX(ctx context.Context) int {
	affected, err := eau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eau *EntityAddressUpdate) Exec(ctx context.Context) error {
	_, err := eau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eau *EntityAddressUpdate) ExecX(ctx context.Context) {
	if err := eau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eau *EntityAddressUpdate) check() error {
	if v, ok := eau.mutation.GetType(); ok {
		if err := entityaddress.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (eau *EntityAddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entityaddress.Table,
			Columns: entityaddress.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entityaddress.FieldID,
			},
		},
	}
	if ps := eau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eau.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldCountry,
		})
	}
	if value, ok := eau.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldCity,
		})
	}
	if value, ok := eau.mutation.PostalCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldPostalCode,
		})
	}
	if value, ok := eau.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldState,
		})
	}
	if value, ok := eau.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: entityaddress.FieldType,
		})
	}
	if value, ok := eau.mutation.Line1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldLine1,
		})
	}
	if value, ok := eau.mutation.Line2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldLine2,
		})
	}
	if eau.mutation.Line2Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: entityaddress.FieldLine2,
		})
	}
	if value, ok := eau.mutation.Line3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldLine3,
		})
	}
	if eau.mutation.Line3Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: entityaddress.FieldLine3,
		})
	}
	if value, ok := eau.mutation.Primary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entityaddress.FieldPrimary,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entityaddress.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EntityAddressUpdateOne is the builder for updating a single EntityAddress entity.
type EntityAddressUpdateOne struct {
	config
	hooks    []Hook
	mutation *EntityAddressMutation
}

// SetCountry sets the "country" field.
func (eauo *EntityAddressUpdateOne) SetCountry(s string) *EntityAddressUpdateOne {
	eauo.mutation.SetCountry(s)
	return eauo
}

// SetCity sets the "city" field.
func (eauo *EntityAddressUpdateOne) SetCity(s string) *EntityAddressUpdateOne {
	eauo.mutation.SetCity(s)
	return eauo
}

// SetPostalCode sets the "postalCode" field.
func (eauo *EntityAddressUpdateOne) SetPostalCode(s string) *EntityAddressUpdateOne {
	eauo.mutation.SetPostalCode(s)
	return eauo
}

// SetState sets the "state" field.
func (eauo *EntityAddressUpdateOne) SetState(s string) *EntityAddressUpdateOne {
	eauo.mutation.SetState(s)
	return eauo
}

// SetType sets the "type" field.
func (eauo *EntityAddressUpdateOne) SetType(e entityaddress.Type) *EntityAddressUpdateOne {
	eauo.mutation.SetType(e)
	return eauo
}

// SetLine1 sets the "line1" field.
func (eauo *EntityAddressUpdateOne) SetLine1(s string) *EntityAddressUpdateOne {
	eauo.mutation.SetLine1(s)
	return eauo
}

// SetLine2 sets the "line2" field.
func (eauo *EntityAddressUpdateOne) SetLine2(s string) *EntityAddressUpdateOne {
	eauo.mutation.SetLine2(s)
	return eauo
}

// SetNillableLine2 sets the "line2" field if the given value is not nil.
func (eauo *EntityAddressUpdateOne) SetNillableLine2(s *string) *EntityAddressUpdateOne {
	if s != nil {
		eauo.SetLine2(*s)
	}
	return eauo
}

// ClearLine2 clears the value of the "line2" field.
func (eauo *EntityAddressUpdateOne) ClearLine2() *EntityAddressUpdateOne {
	eauo.mutation.ClearLine2()
	return eauo
}

// SetLine3 sets the "line3" field.
func (eauo *EntityAddressUpdateOne) SetLine3(s string) *EntityAddressUpdateOne {
	eauo.mutation.SetLine3(s)
	return eauo
}

// SetNillableLine3 sets the "line3" field if the given value is not nil.
func (eauo *EntityAddressUpdateOne) SetNillableLine3(s *string) *EntityAddressUpdateOne {
	if s != nil {
		eauo.SetLine3(*s)
	}
	return eauo
}

// ClearLine3 clears the value of the "line3" field.
func (eauo *EntityAddressUpdateOne) ClearLine3() *EntityAddressUpdateOne {
	eauo.mutation.ClearLine3()
	return eauo
}

// SetPrimary sets the "primary" field.
func (eauo *EntityAddressUpdateOne) SetPrimary(b bool) *EntityAddressUpdateOne {
	eauo.mutation.SetPrimary(b)
	return eauo
}

// Mutation returns the EntityAddressMutation object of the builder.
func (eauo *EntityAddressUpdateOne) Mutation() *EntityAddressMutation {
	return eauo.mutation
}

// Save executes the query and returns the updated EntityAddress entity.
func (eauo *EntityAddressUpdateOne) Save(ctx context.Context) (*EntityAddress, error) {
	var (
		err  error
		node *EntityAddress
	)
	if len(eauo.hooks) == 0 {
		if err = eauo.check(); err != nil {
			return nil, err
		}
		node, err = eauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityAddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eauo.check(); err != nil {
				return nil, err
			}
			eauo.mutation = mutation
			node, err = eauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(eauo.hooks) - 1; i >= 0; i-- {
			mut = eauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (eauo *EntityAddressUpdateOne) SaveX(ctx context.Context) *EntityAddress {
	node, err := eauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eauo *EntityAddressUpdateOne) Exec(ctx context.Context) error {
	_, err := eauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eauo *EntityAddressUpdateOne) ExecX(ctx context.Context) {
	if err := eauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eauo *EntityAddressUpdateOne) check() error {
	if v, ok := eauo.mutation.GetType(); ok {
		if err := entityaddress.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (eauo *EntityAddressUpdateOne) sqlSave(ctx context.Context) (_node *EntityAddress, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entityaddress.Table,
			Columns: entityaddress.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entityaddress.FieldID,
			},
		},
	}
	id, ok := eauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing EntityAddress.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := eauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eauo.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldCountry,
		})
	}
	if value, ok := eauo.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldCity,
		})
	}
	if value, ok := eauo.mutation.PostalCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldPostalCode,
		})
	}
	if value, ok := eauo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldState,
		})
	}
	if value, ok := eauo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: entityaddress.FieldType,
		})
	}
	if value, ok := eauo.mutation.Line1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldLine1,
		})
	}
	if value, ok := eauo.mutation.Line2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldLine2,
		})
	}
	if eauo.mutation.Line2Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: entityaddress.FieldLine2,
		})
	}
	if value, ok := eauo.mutation.Line3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entityaddress.FieldLine3,
		})
	}
	if eauo.mutation.Line3Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: entityaddress.FieldLine3,
		})
	}
	if value, ok := eauo.mutation.Primary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entityaddress.FieldPrimary,
		})
	}
	_node = &EntityAddress{config: eauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entityaddress.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
