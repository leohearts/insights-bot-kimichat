// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/leohearts/insights-bot-kimichat/ent/internal"
	"github.com/leohearts/insights-bot-kimichat/ent/logchathistoriesrecap"
	"github.com/leohearts/insights-bot-kimichat/ent/predicate"
)

// LogChatHistoriesRecapDelete is the builder for deleting a LogChatHistoriesRecap entity.
type LogChatHistoriesRecapDelete struct {
	config
	hooks    []Hook
	mutation *LogChatHistoriesRecapMutation
}

// Where appends a list predicates to the LogChatHistoriesRecapDelete builder.
func (lchrd *LogChatHistoriesRecapDelete) Where(ps ...predicate.LogChatHistoriesRecap) *LogChatHistoriesRecapDelete {
	lchrd.mutation.Where(ps...)
	return lchrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lchrd *LogChatHistoriesRecapDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lchrd.sqlExec, lchrd.mutation, lchrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lchrd *LogChatHistoriesRecapDelete) ExecX(ctx context.Context) int {
	n, err := lchrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lchrd *LogChatHistoriesRecapDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(logchathistoriesrecap.Table, sqlgraph.NewFieldSpec(logchathistoriesrecap.FieldID, field.TypeUUID))
	_spec.Node.Schema = lchrd.schemaConfig.LogChatHistoriesRecap
	ctx = internal.NewSchemaConfigContext(ctx, lchrd.schemaConfig)
	if ps := lchrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lchrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lchrd.mutation.done = true
	return affected, err
}

// LogChatHistoriesRecapDeleteOne is the builder for deleting a single LogChatHistoriesRecap entity.
type LogChatHistoriesRecapDeleteOne struct {
	lchrd *LogChatHistoriesRecapDelete
}

// Where appends a list predicates to the LogChatHistoriesRecapDelete builder.
func (lchrdo *LogChatHistoriesRecapDeleteOne) Where(ps ...predicate.LogChatHistoriesRecap) *LogChatHistoriesRecapDeleteOne {
	lchrdo.lchrd.mutation.Where(ps...)
	return lchrdo
}

// Exec executes the deletion query.
func (lchrdo *LogChatHistoriesRecapDeleteOne) Exec(ctx context.Context) error {
	n, err := lchrdo.lchrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{logchathistoriesrecap.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lchrdo *LogChatHistoriesRecapDeleteOne) ExecX(ctx context.Context) {
	if err := lchrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
