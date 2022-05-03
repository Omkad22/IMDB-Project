package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"imdbv2/ent"
	"imdbv2/ent/director"
	"imdbv2/ent/movie"
)

func (r *mutationResolver) CreateMovie(ctx context.Context, movie MovieInput) (*ent.Movie, error) {
	mov, err := r.client.Movie.Create().SetTitle(movie.Title).SetGenre(movie.Genre).SetDescription(movie.Description).SetRank(movie.Rank).SetDirectorID(movie.DirectorID).Save(ctx)
	if err != nil {
		return nil, ent.MaskNotFound(err)
	}

	review, error := r.client.Review.Create().SetTopic(movie.Topic).SetText(movie.Text).SetRank(movie.Rank).SetMovieID(mov.ID).Save(ctx)
	fmt.Println("new review made", review)
	if error != nil {
		return nil, ent.MaskNotFound(err)
	}

	return mov, err
}
func (r *mutationResolver) CreateDirector(ctx context.Context, director DirectorInput) (*ent.Director, error) {
	return r.client.Director.Create().
		SetName(director.Name).
		SetProfileImage("https://hope.be/wp-content/uploads/2015/05/no-user-image.gif").
		SetBornAt("1.1.1111").
		Save(ctx)
}

func (r *mutationResolver) CreateReview(ctx context.Context, text string, rank int, movieID int, topic string) (*ent.Review, error) {
	return r.client.Review.Create().
		SetTopic(topic).
		SetText(text).
		SetRank(rank).
		SetMovieID(movieID).
		Save(ctx)
}

func (r *mutationResolver) CreateUser(ctx context.Context, firstname string, lastname string, nickname string, description string, email string, profile string, birthday string, password string) (*ent.User, error) {
	return r.client.User.Create().
		SetFirstname(firstname).
		SetLastname(lastname).
		SetNickname(nickname).
		SetEmail(email).
		SetPassword(password).
		SetDescription(description).
		SetBirthDay(birthday).
		SetProfile(profile).
		Save(ctx)
}

func (r *queryResolver) Movies(ctx context.Context) ([]*ent.Movie, error) {
	return r.client.Movie.Query().All(ctx)
}

func (r *queryResolver) Directors(ctx context.Context) ([]*ent.Director, error) {
	return r.client.Director.Query().All(ctx)
}

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DirectorIDByName(ctx context.Context, name string) (*int, error) {
	id, err := r.client.Director.Query().Where(director.Name(name)).OnlyID(ctx)
	if err != nil {
		return nil, ent.MaskNotFound(err)
	}
	return &id, nil
}

func (r *mutationResolver) CreateMovieAndDirector(ctx context.Context, title string, description string, rank int, genre string, directorName string, image string, topic string, text string, profileImage string, bornAt string) (*ent.Movie, error) {
	newDirector, err := r.client.Director.Create().
		SetName(directorName).
		SetProfileImage(profileImage).
		SetBornAt(bornAt).
		Save(ctx)
	if err != nil {
		return nil, ent.MaskNotFound(err)
	}

	fmt.Println("new director made", newDirector)

	movie, error := r.client.Movie.Create().
		SetTitle(title).
		SetGenre(genre).
		SetDescription(description).
		SetRank(rank).
		SetDirectorID(newDirector.ID).
		SetImage(image).
		Save(ctx)
	if error != nil {
		return nil, ent.MaskNotFound(err)
	}

	review, e := r.client.Review.Create().SetTopic(topic).SetText(text).SetRank(rank).SetMovieID(movie.ID).Save(ctx)
	if e != nil {
		return nil, ent.MaskNotFound(err)
	}

	fmt.Println("new review made", review)

	return movie, error
}

// Mutation returns imdbv2.MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns imdbv2.QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *queryResolver) MovieByID(ctx context.Context, id int) ([]*ent.Movie, error) {
	data, err := r.client.Movie.Query().Where(movie.ID(id)).All(ctx)
	if err != nil {
		return nil, ent.MaskNotFound(err)
	}
	return data, nil
}

func (r *queryResolver) DirectorByID(ctx context.Context, id int) ([]*ent.Director, error) {
	return r.client.Director.Query().Where(director.ID(id)).All(ctx)
}

func (r *queryResolver) ReviewsOfMovie(ctx context.Context, movieID int) ([]*ent.Review, error) {
	data, err := r.client.Movie.Query().Where(movie.ID(movieID)).QueryReviews().All(ctx)
	if err != nil {
		return nil, ent.MaskNotFound(err)
	}
	return data, nil
}

func (r *queryResolver) Top10Movies(ctx context.Context) ([]*ent.Movie, error) {
	data, err := r.client.Movie.Query().Order(ent.Desc(movie.FieldRank)).Limit(10).All(ctx)
	if err != nil {
		return nil, ent.MaskNotFound(err)
	}
	return data, nil
}

func (r *mutationResolver) UpdateRank(ctx context.Context, id int, rank int) (*ent.Movie, error) {
	data, err := r.client.Movie.UpdateOneID(id).SetRank(rank).Save(ctx)
	if err != nil {
		return nil, ent.MaskNotFound(err)
	}
	return data, nil
}

func (r *mutationResolver) UpdateDirectorDetails(ctx context.Context, id int, bornAt string, profileImage string, description string) (*ent.Director, error) {
	data, err := r.client.Director.UpdateOneID(id).SetBornAt(bornAt).SetProfileImage(profileImage).SetDescription(description).Save(ctx)
	if err != nil {
		return nil, ent.MaskNotFound(err)
	}
	return data, nil
}
