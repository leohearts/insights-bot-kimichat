// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/leohearts/insights-bot-kimichat/ent/internal"
	"github.com/leohearts/insights-bot-kimichat/ent/logsummarizations"
	"github.com/leohearts/insights-bot-kimichat/ent/predicate"
)

// LogSummarizationsQuery is the builder for querying LogSummarizations entities.
type LogSummarizationsQuery struct {
	config
	ctx        *QueryContext
	order      []logsummarizations.OrderOption
	inters     []Interceptor
	predicates []predicate.LogSummarizations
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LogSummarizationsQuery builder.
func (lsq *LogSummarizationsQuery) Where(ps ...predicate.LogSummarizations) *LogSummarizationsQuery {
	lsq.predicates = append(lsq.predicates, ps...)
	return lsq
}

// Limit the number of records to be returned by this query.
func (lsq *LogSummarizationsQuery) Limit(limit int) *LogSummarizationsQuery {
	lsq.ctx.Limit = &limit
	return lsq
}

// Offset to start from.
func (lsq *LogSummarizationsQuery) Offset(offset int) *LogSummarizationsQuery {
	lsq.ctx.Offset = &offset
	return lsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lsq *LogSummarizationsQuery) Unique(unique bool) *LogSummarizationsQuery {
	lsq.ctx.Unique = &unique
	return lsq
}

// Order specifies how the records should be ordered.
func (lsq *LogSummarizationsQuery) Order(o ...logsummarizations.OrderOption) *LogSummarizationsQuery {
	lsq.order = append(lsq.order, o...)
	return lsq
}

// First returns the first LogSummarizations entity from the query.
// Returns a *NotFoundError when no LogSummarizations was found.
func (lsq *LogSummarizationsQuery) First(ctx context.Context) (*LogSummarizations, error) {
	nodes, err := lsq.Limit(1).All(setContextOp(ctx, lsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{logsummarizations.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lsq *LogSummarizationsQuery) FirstX(ctx context.Context) *LogSummarizations {
	node, err := lsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LogSummarizations ID from the query.
// Returns a *NotFoundError when no LogSummarizations ID was found.
func (lsq *LogSummarizationsQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lsq.Limit(1).IDs(setContextOp(ctx, lsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{logsummarizations.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lsq *LogSummarizationsQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := lsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LogSummarizations entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LogSummarizations entity is found.
// Returns a *NotFoundError when no LogSummarizations entities are found.
func (lsq *LogSummarizationsQuery) Only(ctx context.Context) (*LogSummarizations, error) {
	nodes, err := lsq.Limit(2).All(setContextOp(ctx, lsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{logsummarizations.Label}
	default:
		return nil, &NotSingularError{logsummarizations.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lsq *LogSummarizationsQuery) OnlyX(ctx context.Context) *LogSummarizations {
	node, err := lsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LogSummarizations ID in the query.
// Returns a *NotSingularError when more than one LogSummarizations ID is found.
// Returns a *NotFoundError when no entities are found.
func (lsq *LogSummarizationsQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lsq.Limit(2).IDs(setContextOp(ctx, lsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{logsummarizations.Label}
	default:
		err = &NotSingularError{logsummarizations.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lsq *LogSummarizationsQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := lsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LogSummarizationsSlice.
func (lsq *LogSummarizationsQuery) All(ctx context.Context) ([]*LogSummarizations, error) {
	ctx = setContextOp(ctx, lsq.ctx, "All")
	if err := lsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LogSummarizations, *LogSummarizationsQuery]()
	return withInterceptors[[]*LogSummarizations](ctx, lsq, qr, lsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lsq *LogSummarizationsQuery) AllX(ctx context.Context) []*LogSummarizations {
	nodes, err := lsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LogSummarizations IDs.
func (lsq *LogSummarizationsQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if lsq.ctx.Unique == nil && lsq.path != nil {
		lsq.Unique(true)
	}
	ctx = setContextOp(ctx, lsq.ctx, "IDs")
	if err = lsq.Select(logsummarizations.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lsq *LogSummarizationsQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := lsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lsq *LogSummarizationsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lsq.ctx, "Count")
	if err := lsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lsq, querierCount[*LogSummarizationsQuery](), lsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lsq *LogSummarizationsQuery) CountX(ctx context.Context) int {
	count, err := lsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lsq *LogSummarizationsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lsq.ctx, "Exist")
	switch _, err := lsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lsq *LogSummarizationsQuery) ExistX(ctx context.Context) bool {
	exist, err := lsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LogSummarizationsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lsq *LogSummarizationsQuery) Clone() *LogSummarizationsQuery {
	if lsq == nil {
		return nil
	}
	return &LogSummarizationsQuery{
		config:     lsq.config,
		ctx:        lsq.ctx.Clone(),
		order:      append([]logsummarizations.OrderOption{}, lsq.order...),
		inters:     append([]Interceptor{}, lsq.inters...),
		predicates: append([]predicate.LogSummarizations{}, lsq.predicates...),
		// clone intermediate query.
		sql:  lsq.sql.Clone(),
		path: lsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ContentURL string `json:"content_url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LogSummarizations.Query().
//		GroupBy(logsummarizations.FieldContentURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lsq *LogSummarizationsQuery) GroupBy(field string, fields ...string) *LogSummarizationsGroupBy {
	lsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LogSummarizationsGroupBy{build: lsq}
	grbuild.flds = &lsq.ctx.Fields
	grbuild.label = logsummarizations.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ContentURL string `json:"content_url,omitempty"`
//	}
//
//	client.LogSummarizations.Query().
//		Select(logsummarizations.FieldContentURL).
//		Scan(ctx, &v)
func (lsq *LogSummarizationsQuery) Select(fields ...string) *LogSummarizationsSelect {
	lsq.ctx.Fields = append(lsq.ctx.Fields, fields...)
	sbuild := &LogSummarizationsSelect{LogSummarizationsQuery: lsq}
	sbuild.label = logsummarizations.Label
	sbuild.flds, sbuild.scan = &lsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LogSummarizationsSelect configured with the given aggregations.
func (lsq *LogSummarizationsQuery) Aggregate(fns ...AggregateFunc) *LogSummarizationsSelect {
	return lsq.Select().Aggregate(fns...)
}

func (lsq *LogSummarizationsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lsq); err != nil {
				return err
			}
		}
	}
	for _, f := range lsq.ctx.Fields {
		if !logsummarizations.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lsq.path != nil {
		prev, err := lsq.path(ctx)
		if err != nil {
			return err
		}
		lsq.sql = prev
	}
	return nil
}

func (lsq *LogSummarizationsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LogSummarizations, error) {
	var (
		nodes = []*LogSummarizations{}
		_spec = lsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LogSummarizations).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LogSummarizations{config: lsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = lsq.schemaConfig.LogSummarizations
	ctx = internal.NewSchemaConfigContext(ctx, lsq.schemaConfig)
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lsq *LogSummarizationsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lsq.querySpec()
	_spec.Node.Schema = lsq.schemaConfig.LogSummarizations
	ctx = internal.NewSchemaConfigContext(ctx, lsq.schemaConfig)
	_spec.Node.Columns = lsq.ctx.Fields
	if len(lsq.ctx.Fields) > 0 {
		_spec.Unique = lsq.ctx.Unique != nil && *lsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lsq.driver, _spec)
}

func (lsq *LogSummarizationsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(logsummarizations.Table, logsummarizations.Columns, sqlgraph.NewFieldSpec(logsummarizations.FieldID, field.TypeUUID))
	_spec.From = lsq.sql
	if unique := lsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lsq.path != nil {
		_spec.Unique = true
	}
	if fields := lsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logsummarizations.FieldID)
		for i := range fields {
			if fields[i] != logsummarizations.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lsq *LogSummarizationsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lsq.driver.Dialect())
	t1 := builder.Table(logsummarizations.Table)
	columns := lsq.ctx.Fields
	if len(columns) == 0 {
		columns = logsummarizations.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lsq.sql != nil {
		selector = lsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lsq.ctx.Unique != nil && *lsq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(lsq.schemaConfig.LogSummarizations)
	ctx = internal.NewSchemaConfigContext(ctx, lsq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range lsq.predicates {
		p(selector)
	}
	for _, p := range lsq.order {
		p(selector)
	}
	if offset := lsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LogSummarizationsGroupBy is the group-by builder for LogSummarizations entities.
type LogSummarizationsGroupBy struct {
	selector
	build *LogSummarizationsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lsgb *LogSummarizationsGroupBy) Aggregate(fns ...AggregateFunc) *LogSummarizationsGroupBy {
	lsgb.fns = append(lsgb.fns, fns...)
	return lsgb
}

// Scan applies the selector query and scans the result into the given value.
func (lsgb *LogSummarizationsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lsgb.build.ctx, "GroupBy")
	if err := lsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogSummarizationsQuery, *LogSummarizationsGroupBy](ctx, lsgb.build, lsgb, lsgb.build.inters, v)
}

func (lsgb *LogSummarizationsGroupBy) sqlScan(ctx context.Context, root *LogSummarizationsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lsgb.fns))
	for _, fn := range lsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lsgb.flds)+len(lsgb.fns))
		for _, f := range *lsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LogSummarizationsSelect is the builder for selecting fields of LogSummarizations entities.
type LogSummarizationsSelect struct {
	*LogSummarizationsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lss *LogSummarizationsSelect) Aggregate(fns ...AggregateFunc) *LogSummarizationsSelect {
	lss.fns = append(lss.fns, fns...)
	return lss
}

// Scan applies the selector query and scans the result into the given value.
func (lss *LogSummarizationsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lss.ctx, "Select")
	if err := lss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogSummarizationsQuery, *LogSummarizationsSelect](ctx, lss.LogSummarizationsQuery, lss, lss.inters, v)
}

func (lss *LogSummarizationsSelect) sqlScan(ctx context.Context, root *LogSummarizationsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lss.fns))
	for _, fn := range lss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
