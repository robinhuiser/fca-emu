// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/fca-emu/ent/entitycontactpoint"
	"github.com/robinhuiser/fca-emu/ent/predicate"
)

// EntityContactPointQuery is the builder for querying EntityContactPoint entities.
type EntityContactPointQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.EntityContactPoint
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EntityContactPointQuery builder.
func (ecpq *EntityContactPointQuery) Where(ps ...predicate.EntityContactPoint) *EntityContactPointQuery {
	ecpq.predicates = append(ecpq.predicates, ps...)
	return ecpq
}

// Limit adds a limit step to the query.
func (ecpq *EntityContactPointQuery) Limit(limit int) *EntityContactPointQuery {
	ecpq.limit = &limit
	return ecpq
}

// Offset adds an offset step to the query.
func (ecpq *EntityContactPointQuery) Offset(offset int) *EntityContactPointQuery {
	ecpq.offset = &offset
	return ecpq
}

// Order adds an order step to the query.
func (ecpq *EntityContactPointQuery) Order(o ...OrderFunc) *EntityContactPointQuery {
	ecpq.order = append(ecpq.order, o...)
	return ecpq
}

// First returns the first EntityContactPoint entity from the query.
// Returns a *NotFoundError when no EntityContactPoint was found.
func (ecpq *EntityContactPointQuery) First(ctx context.Context) (*EntityContactPoint, error) {
	nodes, err := ecpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entitycontactpoint.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecpq *EntityContactPointQuery) FirstX(ctx context.Context) *EntityContactPoint {
	node, err := ecpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EntityContactPoint ID from the query.
// Returns a *NotFoundError when no EntityContactPoint ID was found.
func (ecpq *EntityContactPointQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entitycontactpoint.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecpq *EntityContactPointQuery) FirstIDX(ctx context.Context) int {
	id, err := ecpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EntityContactPoint entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one EntityContactPoint entity is not found.
// Returns a *NotFoundError when no EntityContactPoint entities are found.
func (ecpq *EntityContactPointQuery) Only(ctx context.Context) (*EntityContactPoint, error) {
	nodes, err := ecpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entitycontactpoint.Label}
	default:
		return nil, &NotSingularError{entitycontactpoint.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecpq *EntityContactPointQuery) OnlyX(ctx context.Context) *EntityContactPoint {
	node, err := ecpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EntityContactPoint ID in the query.
// Returns a *NotSingularError when exactly one EntityContactPoint ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ecpq *EntityContactPointQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entitycontactpoint.Label}
	default:
		err = &NotSingularError{entitycontactpoint.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecpq *EntityContactPointQuery) OnlyIDX(ctx context.Context) int {
	id, err := ecpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EntityContactPoints.
func (ecpq *EntityContactPointQuery) All(ctx context.Context) ([]*EntityContactPoint, error) {
	if err := ecpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ecpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ecpq *EntityContactPointQuery) AllX(ctx context.Context) []*EntityContactPoint {
	nodes, err := ecpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EntityContactPoint IDs.
func (ecpq *EntityContactPointQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ecpq.Select(entitycontactpoint.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecpq *EntityContactPointQuery) IDsX(ctx context.Context) []int {
	ids, err := ecpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecpq *EntityContactPointQuery) Count(ctx context.Context) (int, error) {
	if err := ecpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ecpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ecpq *EntityContactPointQuery) CountX(ctx context.Context) int {
	count, err := ecpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecpq *EntityContactPointQuery) Exist(ctx context.Context) (bool, error) {
	if err := ecpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ecpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ecpq *EntityContactPointQuery) ExistX(ctx context.Context) bool {
	exist, err := ecpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EntityContactPointQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecpq *EntityContactPointQuery) Clone() *EntityContactPointQuery {
	if ecpq == nil {
		return nil
	}
	return &EntityContactPointQuery{
		config:     ecpq.config,
		limit:      ecpq.limit,
		offset:     ecpq.offset,
		order:      append([]OrderFunc{}, ecpq.order...),
		predicates: append([]predicate.EntityContactPoint{}, ecpq.predicates...),
		// clone intermediate query.
		sql:  ecpq.sql.Clone(),
		path: ecpq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Prefix string `json:"prefix,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EntityContactPoint.Query().
//		GroupBy(entitycontactpoint.FieldPrefix).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ecpq *EntityContactPointQuery) GroupBy(field string, fields ...string) *EntityContactPointGroupBy {
	group := &EntityContactPointGroupBy{config: ecpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ecpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ecpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Prefix string `json:"prefix,omitempty"`
//	}
//
//	client.EntityContactPoint.Query().
//		Select(entitycontactpoint.FieldPrefix).
//		Scan(ctx, &v)
//
func (ecpq *EntityContactPointQuery) Select(field string, fields ...string) *EntityContactPointSelect {
	ecpq.fields = append([]string{field}, fields...)
	return &EntityContactPointSelect{EntityContactPointQuery: ecpq}
}

func (ecpq *EntityContactPointQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ecpq.fields {
		if !entitycontactpoint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ecpq.path != nil {
		prev, err := ecpq.path(ctx)
		if err != nil {
			return err
		}
		ecpq.sql = prev
	}
	return nil
}

func (ecpq *EntityContactPointQuery) sqlAll(ctx context.Context) ([]*EntityContactPoint, error) {
	var (
		nodes   = []*EntityContactPoint{}
		withFKs = ecpq.withFKs
		_spec   = ecpq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, entitycontactpoint.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &EntityContactPoint{config: ecpq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ecpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ecpq *EntityContactPointQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecpq.querySpec()
	return sqlgraph.CountNodes(ctx, ecpq.driver, _spec)
}

func (ecpq *EntityContactPointQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ecpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ecpq *EntityContactPointQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entitycontactpoint.Table,
			Columns: entitycontactpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entitycontactpoint.FieldID,
			},
		},
		From:   ecpq.sql,
		Unique: true,
	}
	if fields := ecpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entitycontactpoint.FieldID)
		for i := range fields {
			if fields[i] != entitycontactpoint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ecpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ecpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ecpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ecpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, entitycontactpoint.ValidColumn)
			}
		}
	}
	return _spec
}

func (ecpq *EntityContactPointQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecpq.driver.Dialect())
	t1 := builder.Table(entitycontactpoint.Table)
	selector := builder.Select(t1.Columns(entitycontactpoint.Columns...)...).From(t1)
	if ecpq.sql != nil {
		selector = ecpq.sql
		selector.Select(selector.Columns(entitycontactpoint.Columns...)...)
	}
	for _, p := range ecpq.predicates {
		p(selector)
	}
	for _, p := range ecpq.order {
		p(selector, entitycontactpoint.ValidColumn)
	}
	if offset := ecpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ecpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EntityContactPointGroupBy is the group-by builder for EntityContactPoint entities.
type EntityContactPointGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecpgb *EntityContactPointGroupBy) Aggregate(fns ...AggregateFunc) *EntityContactPointGroupBy {
	ecpgb.fns = append(ecpgb.fns, fns...)
	return ecpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ecpgb *EntityContactPointGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ecpgb.path(ctx)
	if err != nil {
		return err
	}
	ecpgb.sql = query
	return ecpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ecpgb *EntityContactPointGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ecpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ecpgb *EntityContactPointGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ecpgb.fields) > 1 {
		return nil, errors.New("ent: EntityContactPointGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ecpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ecpgb *EntityContactPointGroupBy) StringsX(ctx context.Context) []string {
	v, err := ecpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ecpgb *EntityContactPointGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ecpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{entitycontactpoint.Label}
	default:
		err = fmt.Errorf("ent: EntityContactPointGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ecpgb *EntityContactPointGroupBy) StringX(ctx context.Context) string {
	v, err := ecpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ecpgb *EntityContactPointGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ecpgb.fields) > 1 {
		return nil, errors.New("ent: EntityContactPointGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ecpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ecpgb *EntityContactPointGroupBy) IntsX(ctx context.Context) []int {
	v, err := ecpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ecpgb *EntityContactPointGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ecpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{entitycontactpoint.Label}
	default:
		err = fmt.Errorf("ent: EntityContactPointGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ecpgb *EntityContactPointGroupBy) IntX(ctx context.Context) int {
	v, err := ecpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ecpgb *EntityContactPointGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ecpgb.fields) > 1 {
		return nil, errors.New("ent: EntityContactPointGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ecpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ecpgb *EntityContactPointGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ecpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ecpgb *EntityContactPointGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ecpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{entitycontactpoint.Label}
	default:
		err = fmt.Errorf("ent: EntityContactPointGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ecpgb *EntityContactPointGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ecpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ecpgb *EntityContactPointGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ecpgb.fields) > 1 {
		return nil, errors.New("ent: EntityContactPointGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ecpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ecpgb *EntityContactPointGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ecpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ecpgb *EntityContactPointGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ecpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{entitycontactpoint.Label}
	default:
		err = fmt.Errorf("ent: EntityContactPointGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ecpgb *EntityContactPointGroupBy) BoolX(ctx context.Context) bool {
	v, err := ecpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ecpgb *EntityContactPointGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ecpgb.fields {
		if !entitycontactpoint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ecpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ecpgb *EntityContactPointGroupBy) sqlQuery() *sql.Selector {
	selector := ecpgb.sql
	columns := make([]string, 0, len(ecpgb.fields)+len(ecpgb.fns))
	columns = append(columns, ecpgb.fields...)
	for _, fn := range ecpgb.fns {
		columns = append(columns, fn(selector, entitycontactpoint.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ecpgb.fields...)
}

// EntityContactPointSelect is the builder for selecting fields of EntityContactPoint entities.
type EntityContactPointSelect struct {
	*EntityContactPointQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ecps *EntityContactPointSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ecps.prepareQuery(ctx); err != nil {
		return err
	}
	ecps.sql = ecps.EntityContactPointQuery.sqlQuery(ctx)
	return ecps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ecps *EntityContactPointSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ecps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ecps *EntityContactPointSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ecps.fields) > 1 {
		return nil, errors.New("ent: EntityContactPointSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ecps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ecps *EntityContactPointSelect) StringsX(ctx context.Context) []string {
	v, err := ecps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ecps *EntityContactPointSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ecps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{entitycontactpoint.Label}
	default:
		err = fmt.Errorf("ent: EntityContactPointSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ecps *EntityContactPointSelect) StringX(ctx context.Context) string {
	v, err := ecps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ecps *EntityContactPointSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ecps.fields) > 1 {
		return nil, errors.New("ent: EntityContactPointSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ecps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ecps *EntityContactPointSelect) IntsX(ctx context.Context) []int {
	v, err := ecps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ecps *EntityContactPointSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ecps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{entitycontactpoint.Label}
	default:
		err = fmt.Errorf("ent: EntityContactPointSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ecps *EntityContactPointSelect) IntX(ctx context.Context) int {
	v, err := ecps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ecps *EntityContactPointSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ecps.fields) > 1 {
		return nil, errors.New("ent: EntityContactPointSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ecps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ecps *EntityContactPointSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ecps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ecps *EntityContactPointSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ecps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{entitycontactpoint.Label}
	default:
		err = fmt.Errorf("ent: EntityContactPointSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ecps *EntityContactPointSelect) Float64X(ctx context.Context) float64 {
	v, err := ecps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ecps *EntityContactPointSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ecps.fields) > 1 {
		return nil, errors.New("ent: EntityContactPointSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ecps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ecps *EntityContactPointSelect) BoolsX(ctx context.Context) []bool {
	v, err := ecps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ecps *EntityContactPointSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ecps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{entitycontactpoint.Label}
	default:
		err = fmt.Errorf("ent: EntityContactPointSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ecps *EntityContactPointSelect) BoolX(ctx context.Context) bool {
	v, err := ecps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ecps *EntityContactPointSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ecps.sqlQuery().Query()
	if err := ecps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ecps *EntityContactPointSelect) sqlQuery() sql.Querier {
	selector := ecps.sql
	selector.Select(selector.Columns(ecps.fields...)...)
	return selector
}
