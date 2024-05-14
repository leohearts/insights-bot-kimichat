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
	"github.com/leohearts/insights-bot-kimichat/ent/chathistories"
	"github.com/leohearts/insights-bot-kimichat/ent/internal"
	"github.com/leohearts/insights-bot-kimichat/ent/predicate"
)

// ChatHistoriesQuery is the builder for querying ChatHistories entities.
type ChatHistoriesQuery struct {
	config
	ctx        *QueryContext
	order      []chathistories.OrderOption
	inters     []Interceptor
	predicates []predicate.ChatHistories
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChatHistoriesQuery builder.
func (chq *ChatHistoriesQuery) Where(ps ...predicate.ChatHistories) *ChatHistoriesQuery {
	chq.predicates = append(chq.predicates, ps...)
	return chq
}

// Limit the number of records to be returned by this query.
func (chq *ChatHistoriesQuery) Limit(limit int) *ChatHistoriesQuery {
	chq.ctx.Limit = &limit
	return chq
}

// Offset to start from.
func (chq *ChatHistoriesQuery) Offset(offset int) *ChatHistoriesQuery {
	chq.ctx.Offset = &offset
	return chq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (chq *ChatHistoriesQuery) Unique(unique bool) *ChatHistoriesQuery {
	chq.ctx.Unique = &unique
	return chq
}

// Order specifies how the records should be ordered.
func (chq *ChatHistoriesQuery) Order(o ...chathistories.OrderOption) *ChatHistoriesQuery {
	chq.order = append(chq.order, o...)
	return chq
}

// First returns the first ChatHistories entity from the query.
// Returns a *NotFoundError when no ChatHistories was found.
func (chq *ChatHistoriesQuery) First(ctx context.Context) (*ChatHistories, error) {
	nodes, err := chq.Limit(1).All(setContextOp(ctx, chq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chathistories.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (chq *ChatHistoriesQuery) FirstX(ctx context.Context) *ChatHistories {
	node, err := chq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ChatHistories ID from the query.
// Returns a *NotFoundError when no ChatHistories ID was found.
func (chq *ChatHistoriesQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = chq.Limit(1).IDs(setContextOp(ctx, chq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chathistories.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (chq *ChatHistoriesQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := chq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ChatHistories entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ChatHistories entity is found.
// Returns a *NotFoundError when no ChatHistories entities are found.
func (chq *ChatHistoriesQuery) Only(ctx context.Context) (*ChatHistories, error) {
	nodes, err := chq.Limit(2).All(setContextOp(ctx, chq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chathistories.Label}
	default:
		return nil, &NotSingularError{chathistories.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (chq *ChatHistoriesQuery) OnlyX(ctx context.Context) *ChatHistories {
	node, err := chq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ChatHistories ID in the query.
// Returns a *NotSingularError when more than one ChatHistories ID is found.
// Returns a *NotFoundError when no entities are found.
func (chq *ChatHistoriesQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = chq.Limit(2).IDs(setContextOp(ctx, chq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chathistories.Label}
	default:
		err = &NotSingularError{chathistories.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (chq *ChatHistoriesQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := chq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ChatHistoriesSlice.
func (chq *ChatHistoriesQuery) All(ctx context.Context) ([]*ChatHistories, error) {
	ctx = setContextOp(ctx, chq.ctx, "All")
	if err := chq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ChatHistories, *ChatHistoriesQuery]()
	return withInterceptors[[]*ChatHistories](ctx, chq, qr, chq.inters)
}

// AllX is like All, but panics if an error occurs.
func (chq *ChatHistoriesQuery) AllX(ctx context.Context) []*ChatHistories {
	nodes, err := chq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ChatHistories IDs.
func (chq *ChatHistoriesQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if chq.ctx.Unique == nil && chq.path != nil {
		chq.Unique(true)
	}
	ctx = setContextOp(ctx, chq.ctx, "IDs")
	if err = chq.Select(chathistories.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (chq *ChatHistoriesQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := chq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (chq *ChatHistoriesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, chq.ctx, "Count")
	if err := chq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, chq, querierCount[*ChatHistoriesQuery](), chq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (chq *ChatHistoriesQuery) CountX(ctx context.Context) int {
	count, err := chq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (chq *ChatHistoriesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, chq.ctx, "Exist")
	switch _, err := chq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (chq *ChatHistoriesQuery) ExistX(ctx context.Context) bool {
	exist, err := chq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChatHistoriesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (chq *ChatHistoriesQuery) Clone() *ChatHistoriesQuery {
	if chq == nil {
		return nil
	}
	return &ChatHistoriesQuery{
		config:     chq.config,
		ctx:        chq.ctx.Clone(),
		order:      append([]chathistories.OrderOption{}, chq.order...),
		inters:     append([]Interceptor{}, chq.inters...),
		predicates: append([]predicate.ChatHistories{}, chq.predicates...),
		// clone intermediate query.
		sql:  chq.sql.Clone(),
		path: chq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ChatID int64 `json:"chat_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ChatHistories.Query().
//		GroupBy(chathistories.FieldChatID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (chq *ChatHistoriesQuery) GroupBy(field string, fields ...string) *ChatHistoriesGroupBy {
	chq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChatHistoriesGroupBy{build: chq}
	grbuild.flds = &chq.ctx.Fields
	grbuild.label = chathistories.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ChatID int64 `json:"chat_id,omitempty"`
//	}
//
//	client.ChatHistories.Query().
//		Select(chathistories.FieldChatID).
//		Scan(ctx, &v)
func (chq *ChatHistoriesQuery) Select(fields ...string) *ChatHistoriesSelect {
	chq.ctx.Fields = append(chq.ctx.Fields, fields...)
	sbuild := &ChatHistoriesSelect{ChatHistoriesQuery: chq}
	sbuild.label = chathistories.Label
	sbuild.flds, sbuild.scan = &chq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChatHistoriesSelect configured with the given aggregations.
func (chq *ChatHistoriesQuery) Aggregate(fns ...AggregateFunc) *ChatHistoriesSelect {
	return chq.Select().Aggregate(fns...)
}

func (chq *ChatHistoriesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range chq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, chq); err != nil {
				return err
			}
		}
	}
	for _, f := range chq.ctx.Fields {
		if !chathistories.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if chq.path != nil {
		prev, err := chq.path(ctx)
		if err != nil {
			return err
		}
		chq.sql = prev
	}
	return nil
}

func (chq *ChatHistoriesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ChatHistories, error) {
	var (
		nodes = []*ChatHistories{}
		_spec = chq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ChatHistories).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ChatHistories{config: chq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = chq.schemaConfig.ChatHistories
	ctx = internal.NewSchemaConfigContext(ctx, chq.schemaConfig)
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, chq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (chq *ChatHistoriesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := chq.querySpec()
	_spec.Node.Schema = chq.schemaConfig.ChatHistories
	ctx = internal.NewSchemaConfigContext(ctx, chq.schemaConfig)
	_spec.Node.Columns = chq.ctx.Fields
	if len(chq.ctx.Fields) > 0 {
		_spec.Unique = chq.ctx.Unique != nil && *chq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, chq.driver, _spec)
}

func (chq *ChatHistoriesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(chathistories.Table, chathistories.Columns, sqlgraph.NewFieldSpec(chathistories.FieldID, field.TypeUUID))
	_spec.From = chq.sql
	if unique := chq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if chq.path != nil {
		_spec.Unique = true
	}
	if fields := chq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chathistories.FieldID)
		for i := range fields {
			if fields[i] != chathistories.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := chq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := chq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := chq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := chq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (chq *ChatHistoriesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(chq.driver.Dialect())
	t1 := builder.Table(chathistories.Table)
	columns := chq.ctx.Fields
	if len(columns) == 0 {
		columns = chathistories.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if chq.sql != nil {
		selector = chq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if chq.ctx.Unique != nil && *chq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(chq.schemaConfig.ChatHistories)
	ctx = internal.NewSchemaConfigContext(ctx, chq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range chq.predicates {
		p(selector)
	}
	for _, p := range chq.order {
		p(selector)
	}
	if offset := chq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := chq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChatHistoriesGroupBy is the group-by builder for ChatHistories entities.
type ChatHistoriesGroupBy struct {
	selector
	build *ChatHistoriesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (chgb *ChatHistoriesGroupBy) Aggregate(fns ...AggregateFunc) *ChatHistoriesGroupBy {
	chgb.fns = append(chgb.fns, fns...)
	return chgb
}

// Scan applies the selector query and scans the result into the given value.
func (chgb *ChatHistoriesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, chgb.build.ctx, "GroupBy")
	if err := chgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChatHistoriesQuery, *ChatHistoriesGroupBy](ctx, chgb.build, chgb, chgb.build.inters, v)
}

func (chgb *ChatHistoriesGroupBy) sqlScan(ctx context.Context, root *ChatHistoriesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(chgb.fns))
	for _, fn := range chgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*chgb.flds)+len(chgb.fns))
		for _, f := range *chgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*chgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChatHistoriesSelect is the builder for selecting fields of ChatHistories entities.
type ChatHistoriesSelect struct {
	*ChatHistoriesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (chs *ChatHistoriesSelect) Aggregate(fns ...AggregateFunc) *ChatHistoriesSelect {
	chs.fns = append(chs.fns, fns...)
	return chs
}

// Scan applies the selector query and scans the result into the given value.
func (chs *ChatHistoriesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, chs.ctx, "Select")
	if err := chs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChatHistoriesQuery, *ChatHistoriesSelect](ctx, chs.ChatHistoriesQuery, chs, chs.inters, v)
}

func (chs *ChatHistoriesSelect) sqlScan(ctx context.Context, root *ChatHistoriesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(chs.fns))
	for _, fn := range chs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*chs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
