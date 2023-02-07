// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (t *Temple) Children(ctx context.Context) (result []*Temple, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.ChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryChildren().All(ctx)
	}
	return result, err
}

func (t *Temple) Parent(ctx context.Context) (*Temple, error) {
	result, err := t.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Temple) Owner(ctx context.Context) (*User, error) {
	result, err := t.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Todos(ctx context.Context) (result []*Temple, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedTodos(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.TodosOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryTodos().All(ctx)
	}
	return result, err
}
