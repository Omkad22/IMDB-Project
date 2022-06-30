// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *ActorQuery) CollectFields(ctx context.Context, satisfies ...string) *ActorQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		a = a.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return a
}

func (a *ActorQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ActorQuery {
	return a
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CommentQuery) CollectFields(ctx context.Context, satisfies ...string) *CommentQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return c
}

func (c *CommentQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *CommentQuery {
	return c
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (d *DirectorQuery) CollectFields(ctx context.Context, satisfies ...string) *DirectorQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		d = d.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return d
}

func (d *DirectorQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *DirectorQuery {
	return d
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (f *FavoriteQuery) CollectFields(ctx context.Context, satisfies ...string) *FavoriteQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		f = f.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return f
}

func (f *FavoriteQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *FavoriteQuery {
	return f
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (l *LikeQuery) CollectFields(ctx context.Context, satisfies ...string) *LikeQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		l = l.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return l
}

func (l *LikeQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *LikeQuery {
	return l
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (m *MovieQuery) CollectFields(ctx context.Context, satisfies ...string) *MovieQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		m = m.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return m
}

func (m *MovieQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *MovieQuery {
	return m
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *ReviewQuery) CollectFields(ctx context.Context, satisfies ...string) *ReviewQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		r = r.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return r
}

func (r *ReviewQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ReviewQuery {
	return r
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	return u
}
