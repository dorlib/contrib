// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithpackagename"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithPackageNameQuery is the builder for querying MessageWithPackageName entities.
type MessageWithPackageNameQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MessageWithPackageName
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageWithPackageNameQuery builder.
func (mwpnq *MessageWithPackageNameQuery) Where(ps ...predicate.MessageWithPackageName) *MessageWithPackageNameQuery {
	mwpnq.predicates = append(mwpnq.predicates, ps...)
	return mwpnq
}

// Limit adds a limit step to the query.
func (mwpnq *MessageWithPackageNameQuery) Limit(limit int) *MessageWithPackageNameQuery {
	mwpnq.limit = &limit
	return mwpnq
}

// Offset adds an offset step to the query.
func (mwpnq *MessageWithPackageNameQuery) Offset(offset int) *MessageWithPackageNameQuery {
	mwpnq.offset = &offset
	return mwpnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mwpnq *MessageWithPackageNameQuery) Unique(unique bool) *MessageWithPackageNameQuery {
	mwpnq.unique = &unique
	return mwpnq
}

// Order adds an order step to the query.
func (mwpnq *MessageWithPackageNameQuery) Order(o ...OrderFunc) *MessageWithPackageNameQuery {
	mwpnq.order = append(mwpnq.order, o...)
	return mwpnq
}

// First returns the first MessageWithPackageName entity from the query.
// Returns a *NotFoundError when no MessageWithPackageName was found.
func (mwpnq *MessageWithPackageNameQuery) First(ctx context.Context) (*MessageWithPackageName, error) {
	nodes, err := mwpnq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{messagewithpackagename.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) FirstX(ctx context.Context) *MessageWithPackageName {
	node, err := mwpnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MessageWithPackageName ID from the query.
// Returns a *NotFoundError when no MessageWithPackageName ID was found.
func (mwpnq *MessageWithPackageNameQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mwpnq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{messagewithpackagename.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) FirstIDX(ctx context.Context) int {
	id, err := mwpnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MessageWithPackageName entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one MessageWithPackageName entity is not found.
// Returns a *NotFoundError when no MessageWithPackageName entities are found.
func (mwpnq *MessageWithPackageNameQuery) Only(ctx context.Context) (*MessageWithPackageName, error) {
	nodes, err := mwpnq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{messagewithpackagename.Label}
	default:
		return nil, &NotSingularError{messagewithpackagename.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) OnlyX(ctx context.Context) *MessageWithPackageName {
	node, err := mwpnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MessageWithPackageName ID in the query.
// Returns a *NotSingularError when exactly one MessageWithPackageName ID is not found.
// Returns a *NotFoundError when no entities are found.
func (mwpnq *MessageWithPackageNameQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mwpnq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{messagewithpackagename.Label}
	default:
		err = &NotSingularError{messagewithpackagename.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) OnlyIDX(ctx context.Context) int {
	id, err := mwpnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MessageWithPackageNames.
func (mwpnq *MessageWithPackageNameQuery) All(ctx context.Context) ([]*MessageWithPackageName, error) {
	if err := mwpnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mwpnq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) AllX(ctx context.Context) []*MessageWithPackageName {
	nodes, err := mwpnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MessageWithPackageName IDs.
func (mwpnq *MessageWithPackageNameQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mwpnq.Select(messagewithpackagename.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) IDsX(ctx context.Context) []int {
	ids, err := mwpnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mwpnq *MessageWithPackageNameQuery) Count(ctx context.Context) (int, error) {
	if err := mwpnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mwpnq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) CountX(ctx context.Context) int {
	count, err := mwpnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mwpnq *MessageWithPackageNameQuery) Exist(ctx context.Context) (bool, error) {
	if err := mwpnq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mwpnq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) ExistX(ctx context.Context) bool {
	exist, err := mwpnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageWithPackageNameQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mwpnq *MessageWithPackageNameQuery) Clone() *MessageWithPackageNameQuery {
	if mwpnq == nil {
		return nil
	}
	return &MessageWithPackageNameQuery{
		config:     mwpnq.config,
		limit:      mwpnq.limit,
		offset:     mwpnq.offset,
		order:      append([]OrderFunc{}, mwpnq.order...),
		predicates: append([]predicate.MessageWithPackageName{}, mwpnq.predicates...),
		// clone intermediate query.
		sql:  mwpnq.sql.Clone(),
		path: mwpnq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MessageWithPackageName.Query().
//		GroupBy(messagewithpackagename.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mwpnq *MessageWithPackageNameQuery) GroupBy(field string, fields ...string) *MessageWithPackageNameGroupBy {
	group := &MessageWithPackageNameGroupBy{config: mwpnq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mwpnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mwpnq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.MessageWithPackageName.Query().
//		Select(messagewithpackagename.FieldName).
//		Scan(ctx, &v)
//
func (mwpnq *MessageWithPackageNameQuery) Select(field string, fields ...string) *MessageWithPackageNameSelect {
	mwpnq.fields = append([]string{field}, fields...)
	return &MessageWithPackageNameSelect{MessageWithPackageNameQuery: mwpnq}
}

func (mwpnq *MessageWithPackageNameQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mwpnq.fields {
		if !messagewithpackagename.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mwpnq.path != nil {
		prev, err := mwpnq.path(ctx)
		if err != nil {
			return err
		}
		mwpnq.sql = prev
	}
	return nil
}

func (mwpnq *MessageWithPackageNameQuery) sqlAll(ctx context.Context) ([]*MessageWithPackageName, error) {
	var (
		nodes = []*MessageWithPackageName{}
		_spec = mwpnq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MessageWithPackageName{config: mwpnq.config}
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
	if err := sqlgraph.QueryNodes(ctx, mwpnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mwpnq *MessageWithPackageNameQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mwpnq.querySpec()
	return sqlgraph.CountNodes(ctx, mwpnq.driver, _spec)
}

func (mwpnq *MessageWithPackageNameQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mwpnq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mwpnq *MessageWithPackageNameQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messagewithpackagename.Table,
			Columns: messagewithpackagename.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithpackagename.FieldID,
			},
		},
		From:   mwpnq.sql,
		Unique: true,
	}
	if unique := mwpnq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mwpnq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithpackagename.FieldID)
		for i := range fields {
			if fields[i] != messagewithpackagename.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mwpnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mwpnq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mwpnq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mwpnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mwpnq *MessageWithPackageNameQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mwpnq.driver.Dialect())
	t1 := builder.Table(messagewithpackagename.Table)
	columns := mwpnq.fields
	if len(columns) == 0 {
		columns = messagewithpackagename.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mwpnq.sql != nil {
		selector = mwpnq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range mwpnq.predicates {
		p(selector)
	}
	for _, p := range mwpnq.order {
		p(selector)
	}
	if offset := mwpnq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mwpnq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MessageWithPackageNameGroupBy is the group-by builder for MessageWithPackageName entities.
type MessageWithPackageNameGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mwpngb *MessageWithPackageNameGroupBy) Aggregate(fns ...AggregateFunc) *MessageWithPackageNameGroupBy {
	mwpngb.fns = append(mwpngb.fns, fns...)
	return mwpngb
}

// Scan applies the group-by query and scans the result into the given value.
func (mwpngb *MessageWithPackageNameGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mwpngb.path(ctx)
	if err != nil {
		return err
	}
	mwpngb.sql = query
	return mwpngb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mwpngb *MessageWithPackageNameGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mwpngb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mwpngb *MessageWithPackageNameGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mwpngb.fields) > 1 {
		return nil, errors.New("ent: MessageWithPackageNameGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mwpngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mwpngb *MessageWithPackageNameGroupBy) StringsX(ctx context.Context) []string {
	v, err := mwpngb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mwpngb *MessageWithPackageNameGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mwpngb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithpackagename.Label}
	default:
		err = fmt.Errorf("ent: MessageWithPackageNameGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mwpngb *MessageWithPackageNameGroupBy) StringX(ctx context.Context) string {
	v, err := mwpngb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mwpngb *MessageWithPackageNameGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mwpngb.fields) > 1 {
		return nil, errors.New("ent: MessageWithPackageNameGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mwpngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mwpngb *MessageWithPackageNameGroupBy) IntsX(ctx context.Context) []int {
	v, err := mwpngb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mwpngb *MessageWithPackageNameGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mwpngb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithpackagename.Label}
	default:
		err = fmt.Errorf("ent: MessageWithPackageNameGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mwpngb *MessageWithPackageNameGroupBy) IntX(ctx context.Context) int {
	v, err := mwpngb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mwpngb *MessageWithPackageNameGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mwpngb.fields) > 1 {
		return nil, errors.New("ent: MessageWithPackageNameGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mwpngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mwpngb *MessageWithPackageNameGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mwpngb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mwpngb *MessageWithPackageNameGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mwpngb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithpackagename.Label}
	default:
		err = fmt.Errorf("ent: MessageWithPackageNameGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mwpngb *MessageWithPackageNameGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mwpngb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mwpngb *MessageWithPackageNameGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mwpngb.fields) > 1 {
		return nil, errors.New("ent: MessageWithPackageNameGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mwpngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mwpngb *MessageWithPackageNameGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mwpngb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mwpngb *MessageWithPackageNameGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mwpngb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithpackagename.Label}
	default:
		err = fmt.Errorf("ent: MessageWithPackageNameGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mwpngb *MessageWithPackageNameGroupBy) BoolX(ctx context.Context) bool {
	v, err := mwpngb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mwpngb *MessageWithPackageNameGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mwpngb.fields {
		if !messagewithpackagename.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mwpngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mwpngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mwpngb *MessageWithPackageNameGroupBy) sqlQuery() *sql.Selector {
	selector := mwpngb.sql.Select()
	aggregation := make([]string, 0, len(mwpngb.fns))
	for _, fn := range mwpngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mwpngb.fields)+len(mwpngb.fns))
		for _, f := range mwpngb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mwpngb.fields...)...)
}

// MessageWithPackageNameSelect is the builder for selecting fields of MessageWithPackageName entities.
type MessageWithPackageNameSelect struct {
	*MessageWithPackageNameQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mwpns *MessageWithPackageNameSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mwpns.prepareQuery(ctx); err != nil {
		return err
	}
	mwpns.sql = mwpns.MessageWithPackageNameQuery.sqlQuery(ctx)
	return mwpns.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mwpns *MessageWithPackageNameSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mwpns.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mwpns *MessageWithPackageNameSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mwpns.fields) > 1 {
		return nil, errors.New("ent: MessageWithPackageNameSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mwpns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mwpns *MessageWithPackageNameSelect) StringsX(ctx context.Context) []string {
	v, err := mwpns.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mwpns *MessageWithPackageNameSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mwpns.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithpackagename.Label}
	default:
		err = fmt.Errorf("ent: MessageWithPackageNameSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mwpns *MessageWithPackageNameSelect) StringX(ctx context.Context) string {
	v, err := mwpns.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mwpns *MessageWithPackageNameSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mwpns.fields) > 1 {
		return nil, errors.New("ent: MessageWithPackageNameSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mwpns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mwpns *MessageWithPackageNameSelect) IntsX(ctx context.Context) []int {
	v, err := mwpns.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mwpns *MessageWithPackageNameSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mwpns.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithpackagename.Label}
	default:
		err = fmt.Errorf("ent: MessageWithPackageNameSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mwpns *MessageWithPackageNameSelect) IntX(ctx context.Context) int {
	v, err := mwpns.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mwpns *MessageWithPackageNameSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mwpns.fields) > 1 {
		return nil, errors.New("ent: MessageWithPackageNameSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mwpns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mwpns *MessageWithPackageNameSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mwpns.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mwpns *MessageWithPackageNameSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mwpns.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithpackagename.Label}
	default:
		err = fmt.Errorf("ent: MessageWithPackageNameSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mwpns *MessageWithPackageNameSelect) Float64X(ctx context.Context) float64 {
	v, err := mwpns.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mwpns *MessageWithPackageNameSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mwpns.fields) > 1 {
		return nil, errors.New("ent: MessageWithPackageNameSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mwpns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mwpns *MessageWithPackageNameSelect) BoolsX(ctx context.Context) []bool {
	v, err := mwpns.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mwpns *MessageWithPackageNameSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mwpns.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithpackagename.Label}
	default:
		err = fmt.Errorf("ent: MessageWithPackageNameSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mwpns *MessageWithPackageNameSelect) BoolX(ctx context.Context) bool {
	v, err := mwpns.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mwpns *MessageWithPackageNameSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mwpns.sql.Query()
	if err := mwpns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}