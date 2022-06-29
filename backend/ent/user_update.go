// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imdbv2/ent/comment"
	"imdbv2/ent/predicate"
	"imdbv2/ent/review"
	"imdbv2/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetFirstname sets the "firstname" field.
func (uu *UserUpdate) SetFirstname(s string) *UserUpdate {
	uu.mutation.SetFirstname(s)
	return uu
}

// SetLastname sets the "lastname" field.
func (uu *UserUpdate) SetLastname(s string) *UserUpdate {
	uu.mutation.SetLastname(s)
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetDescription sets the "description" field.
func (uu *UserUpdate) SetDescription(s string) *UserUpdate {
	uu.mutation.SetDescription(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetBirthDay sets the "birthDay" field.
func (uu *UserUpdate) SetBirthDay(s string) *UserUpdate {
	uu.mutation.SetBirthDay(s)
	return uu
}

// SetProfile sets the "profile" field.
func (uu *UserUpdate) SetProfile(s string) *UserUpdate {
	uu.mutation.SetProfile(s)
	return uu
}

// SetCountry sets the "country" field.
func (uu *UserUpdate) SetCountry(s string) *UserUpdate {
	uu.mutation.SetCountry(s)
	return uu
}

// SetGender sets the "gender" field.
func (uu *UserUpdate) SetGender(s string) *UserUpdate {
	uu.mutation.SetGender(s)
	return uu
}

// SetSignupAt sets the "signup_at" field.
func (uu *UserUpdate) SetSignupAt(s string) *UserUpdate {
	uu.mutation.SetSignupAt(s)
	return uu
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (uu *UserUpdate) AddReviewIDs(ids ...int) *UserUpdate {
	uu.mutation.AddReviewIDs(ids...)
	return uu
}

// AddReviews adds the "reviews" edges to the Review entity.
func (uu *UserUpdate) AddReviews(r ...*Review) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddReviewIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (uu *UserUpdate) AddCommentIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCommentIDs(ids...)
	return uu
}

// AddComments adds the "comments" edges to the Comment entity.
func (uu *UserUpdate) AddComments(c ...*Comment) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCommentIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearReviews clears all "reviews" edges to the Review entity.
func (uu *UserUpdate) ClearReviews() *UserUpdate {
	uu.mutation.ClearReviews()
	return uu
}

// RemoveReviewIDs removes the "reviews" edge to Review entities by IDs.
func (uu *UserUpdate) RemoveReviewIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveReviewIDs(ids...)
	return uu
}

// RemoveReviews removes "reviews" edges to Review entities.
func (uu *UserUpdate) RemoveReviews(r ...*Review) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveReviewIDs(ids...)
}

// ClearComments clears all "comments" edges to the Comment entity.
func (uu *UserUpdate) ClearComments() *UserUpdate {
	uu.mutation.ClearComments()
	return uu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (uu *UserUpdate) RemoveCommentIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCommentIDs(ids...)
	return uu
}

// RemoveComments removes "comments" edges to Comment entities.
func (uu *UserUpdate) RemoveComments(c ...*Comment) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCommentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Firstname(); ok {
		if err := user.FirstnameValidator(v); err != nil {
			return &ValidationError{Name: "firstname", err: fmt.Errorf(`ent: validator failed for field "User.firstname": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Lastname(); ok {
		if err := user.LastnameValidator(v); err != nil {
			return &ValidationError{Name: "lastname", err: fmt.Errorf(`ent: validator failed for field "User.lastname": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Firstname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFirstname,
		})
	}
	if value, ok := uu.mutation.Lastname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastname,
		})
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNickname,
		})
	}
	if value, ok := uu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDescription,
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.BirthDay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBirthDay,
		})
	}
	if value, ok := uu.mutation.Profile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldProfile,
		})
	}
	if value, ok := uu.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCountry,
		})
	}
	if value, ok := uu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldGender,
		})
	}
	if value, ok := uu.mutation.SignupAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSignupAt,
		})
	}
	if uu.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReviewsTable,
			Columns: []string{user.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !uu.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReviewsTable,
			Columns: []string{user.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReviewsTable,
			Columns: []string{user.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: user.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !uu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: user.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: user.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetFirstname sets the "firstname" field.
func (uuo *UserUpdateOne) SetFirstname(s string) *UserUpdateOne {
	uuo.mutation.SetFirstname(s)
	return uuo
}

// SetLastname sets the "lastname" field.
func (uuo *UserUpdateOne) SetLastname(s string) *UserUpdateOne {
	uuo.mutation.SetLastname(s)
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetDescription sets the "description" field.
func (uuo *UserUpdateOne) SetDescription(s string) *UserUpdateOne {
	uuo.mutation.SetDescription(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetBirthDay sets the "birthDay" field.
func (uuo *UserUpdateOne) SetBirthDay(s string) *UserUpdateOne {
	uuo.mutation.SetBirthDay(s)
	return uuo
}

// SetProfile sets the "profile" field.
func (uuo *UserUpdateOne) SetProfile(s string) *UserUpdateOne {
	uuo.mutation.SetProfile(s)
	return uuo
}

// SetCountry sets the "country" field.
func (uuo *UserUpdateOne) SetCountry(s string) *UserUpdateOne {
	uuo.mutation.SetCountry(s)
	return uuo
}

// SetGender sets the "gender" field.
func (uuo *UserUpdateOne) SetGender(s string) *UserUpdateOne {
	uuo.mutation.SetGender(s)
	return uuo
}

// SetSignupAt sets the "signup_at" field.
func (uuo *UserUpdateOne) SetSignupAt(s string) *UserUpdateOne {
	uuo.mutation.SetSignupAt(s)
	return uuo
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (uuo *UserUpdateOne) AddReviewIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddReviewIDs(ids...)
	return uuo
}

// AddReviews adds the "reviews" edges to the Review entity.
func (uuo *UserUpdateOne) AddReviews(r ...*Review) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddReviewIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (uuo *UserUpdateOne) AddCommentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCommentIDs(ids...)
	return uuo
}

// AddComments adds the "comments" edges to the Comment entity.
func (uuo *UserUpdateOne) AddComments(c ...*Comment) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCommentIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearReviews clears all "reviews" edges to the Review entity.
func (uuo *UserUpdateOne) ClearReviews() *UserUpdateOne {
	uuo.mutation.ClearReviews()
	return uuo
}

// RemoveReviewIDs removes the "reviews" edge to Review entities by IDs.
func (uuo *UserUpdateOne) RemoveReviewIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveReviewIDs(ids...)
	return uuo
}

// RemoveReviews removes "reviews" edges to Review entities.
func (uuo *UserUpdateOne) RemoveReviews(r ...*Review) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveReviewIDs(ids...)
}

// ClearComments clears all "comments" edges to the Comment entity.
func (uuo *UserUpdateOne) ClearComments() *UserUpdateOne {
	uuo.mutation.ClearComments()
	return uuo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (uuo *UserUpdateOne) RemoveCommentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCommentIDs(ids...)
	return uuo
}

// RemoveComments removes "comments" edges to Comment entities.
func (uuo *UserUpdateOne) RemoveComments(c ...*Comment) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCommentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Firstname(); ok {
		if err := user.FirstnameValidator(v); err != nil {
			return &ValidationError{Name: "firstname", err: fmt.Errorf(`ent: validator failed for field "User.firstname": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Lastname(); ok {
		if err := user.LastnameValidator(v); err != nil {
			return &ValidationError{Name: "lastname", err: fmt.Errorf(`ent: validator failed for field "User.lastname": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Firstname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFirstname,
		})
	}
	if value, ok := uuo.mutation.Lastname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastname,
		})
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNickname,
		})
	}
	if value, ok := uuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDescription,
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.BirthDay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBirthDay,
		})
	}
	if value, ok := uuo.mutation.Profile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldProfile,
		})
	}
	if value, ok := uuo.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCountry,
		})
	}
	if value, ok := uuo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldGender,
		})
	}
	if value, ok := uuo.mutation.SignupAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSignupAt,
		})
	}
	if uuo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReviewsTable,
			Columns: []string{user.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !uuo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReviewsTable,
			Columns: []string{user.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReviewsTable,
			Columns: []string{user.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: user.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !uuo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: user.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: user.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
