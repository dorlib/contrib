// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/invalidfieldmessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InvalidFieldMessageDelete is the builder for deleting a InvalidFieldMessage entity.
type InvalidFieldMessageDelete struct {
	config
	hooks    []Hook
	mutation *InvalidFieldMessageMutation
}

// Where appends a list predicates to the InvalidFieldMessageDelete builder.
func (ifmd *InvalidFieldMessageDelete) Where(ps ...predicate.InvalidFieldMessage) *InvalidFieldMessageDelete {
	ifmd.mutation.Where(ps...)
	return ifmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ifmd *InvalidFieldMessageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ifmd.hooks) == 0 {
		affected, err = ifmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InvalidFieldMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ifmd.mutation = mutation
			affected, err = ifmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ifmd.hooks) - 1; i >= 0; i-- {
			if ifmd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ifmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ifmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifmd *InvalidFieldMessageDelete) ExecX(ctx context.Context) int {
	n, err := ifmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ifmd *InvalidFieldMessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: invalidfieldmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invalidfieldmessage.FieldID,
			},
		},
	}
	if ps := ifmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ifmd.driver, _spec)
}

// InvalidFieldMessageDeleteOne is the builder for deleting a single InvalidFieldMessage entity.
type InvalidFieldMessageDeleteOne struct {
	ifmd *InvalidFieldMessageDelete
}

// Exec executes the deletion query.
func (ifmdo *InvalidFieldMessageDeleteOne) Exec(ctx context.Context) error {
	n, err := ifmdo.ifmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{invalidfieldmessage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ifmdo *InvalidFieldMessageDeleteOne) ExecX(ctx context.Context) {
	ifmdo.ifmd.ExecX(ctx)
}