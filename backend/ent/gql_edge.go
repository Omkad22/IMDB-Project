// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (d *Director) Movies(ctx context.Context) ([]*Movie, error) {
	result, err := d.Edges.MoviesOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryMovies().All(ctx)
	}
	return result, err
}

func (m *Movie) Director(ctx context.Context) (*Director, error) {
	result, err := m.Edges.DirectorOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryDirector().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (m *Movie) Reviews(ctx context.Context) ([]*Review, error) {
	result, err := m.Edges.ReviewsOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryReviews().All(ctx)
	}
	return result, err
}

func (r *Review) Movie(ctx context.Context) (*Movie, error) {
	result, err := r.Edges.MovieOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryMovie().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (r *Review) User(ctx context.Context) (*User, error) {
	result, err := r.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Reviews(ctx context.Context) ([]*Review, error) {
	result, err := u.Edges.ReviewsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryReviews().All(ctx)
	}
	return result, err
}
