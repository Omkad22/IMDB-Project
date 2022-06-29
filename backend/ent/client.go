// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"imdbv2/ent/migrate"

	"imdbv2/ent/actor"
	"imdbv2/ent/comment"
	"imdbv2/ent/director"
	"imdbv2/ent/favorite"
	"imdbv2/ent/movie"
	"imdbv2/ent/review"
	"imdbv2/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Actor is the client for interacting with the Actor builders.
	Actor *ActorClient
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// Director is the client for interacting with the Director builders.
	Director *DirectorClient
	// Favorite is the client for interacting with the Favorite builders.
	Favorite *FavoriteClient
	// Movie is the client for interacting with the Movie builders.
	Movie *MovieClient
	// Review is the client for interacting with the Review builders.
	Review *ReviewClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Actor = NewActorClient(c.config)
	c.Comment = NewCommentClient(c.config)
	c.Director = NewDirectorClient(c.config)
	c.Favorite = NewFavoriteClient(c.config)
	c.Movie = NewMovieClient(c.config)
	c.Review = NewReviewClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Actor:    NewActorClient(cfg),
		Comment:  NewCommentClient(cfg),
		Director: NewDirectorClient(cfg),
		Favorite: NewFavoriteClient(cfg),
		Movie:    NewMovieClient(cfg),
		Review:   NewReviewClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Actor:    NewActorClient(cfg),
		Comment:  NewCommentClient(cfg),
		Director: NewDirectorClient(cfg),
		Favorite: NewFavoriteClient(cfg),
		Movie:    NewMovieClient(cfg),
		Review:   NewReviewClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Actor.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Actor.Use(hooks...)
	c.Comment.Use(hooks...)
	c.Director.Use(hooks...)
	c.Favorite.Use(hooks...)
	c.Movie.Use(hooks...)
	c.Review.Use(hooks...)
	c.User.Use(hooks...)
}

// ActorClient is a client for the Actor schema.
type ActorClient struct {
	config
}

// NewActorClient returns a client for the Actor from the given config.
func NewActorClient(c config) *ActorClient {
	return &ActorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `actor.Hooks(f(g(h())))`.
func (c *ActorClient) Use(hooks ...Hook) {
	c.hooks.Actor = append(c.hooks.Actor, hooks...)
}

// Create returns a create builder for Actor.
func (c *ActorClient) Create() *ActorCreate {
	mutation := newActorMutation(c.config, OpCreate)
	return &ActorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Actor entities.
func (c *ActorClient) CreateBulk(builders ...*ActorCreate) *ActorCreateBulk {
	return &ActorCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Actor.
func (c *ActorClient) Update() *ActorUpdate {
	mutation := newActorMutation(c.config, OpUpdate)
	return &ActorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ActorClient) UpdateOne(a *Actor) *ActorUpdateOne {
	mutation := newActorMutation(c.config, OpUpdateOne, withActor(a))
	return &ActorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ActorClient) UpdateOneID(id int) *ActorUpdateOne {
	mutation := newActorMutation(c.config, OpUpdateOne, withActorID(id))
	return &ActorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Actor.
func (c *ActorClient) Delete() *ActorDelete {
	mutation := newActorMutation(c.config, OpDelete)
	return &ActorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ActorClient) DeleteOne(a *Actor) *ActorDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ActorClient) DeleteOneID(id int) *ActorDeleteOne {
	builder := c.Delete().Where(actor.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ActorDeleteOne{builder}
}

// Query returns a query builder for Actor.
func (c *ActorClient) Query() *ActorQuery {
	return &ActorQuery{
		config: c.config,
	}
}

// Get returns a Actor entity by its id.
func (c *ActorClient) Get(ctx context.Context, id int) (*Actor, error) {
	return c.Query().Where(actor.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ActorClient) GetX(ctx context.Context, id int) *Actor {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryActors queries the actors edge of a Actor.
func (c *ActorClient) QueryActors(a *Actor) *MovieQuery {
	query := &MovieQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, actor.ActorsTable, actor.ActorsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ActorClient) Hooks() []Hook {
	return c.hooks.Actor
}

// CommentClient is a client for the Comment schema.
type CommentClient struct {
	config
}

// NewCommentClient returns a client for the Comment from the given config.
func NewCommentClient(c config) *CommentClient {
	return &CommentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comment.Hooks(f(g(h())))`.
func (c *CommentClient) Use(hooks ...Hook) {
	c.hooks.Comment = append(c.hooks.Comment, hooks...)
}

// Create returns a create builder for Comment.
func (c *CommentClient) Create() *CommentCreate {
	mutation := newCommentMutation(c.config, OpCreate)
	return &CommentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comment entities.
func (c *CommentClient) CreateBulk(builders ...*CommentCreate) *CommentCreateBulk {
	return &CommentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comment.
func (c *CommentClient) Update() *CommentUpdate {
	mutation := newCommentMutation(c.config, OpUpdate)
	return &CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentClient) UpdateOne(co *Comment) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withComment(co))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentClient) UpdateOneID(id int) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withCommentID(id))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comment.
func (c *CommentClient) Delete() *CommentDelete {
	mutation := newCommentMutation(c.config, OpDelete)
	return &CommentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CommentClient) DeleteOne(co *Comment) *CommentDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CommentClient) DeleteOneID(id int) *CommentDeleteOne {
	builder := c.Delete().Where(comment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentDeleteOne{builder}
}

// Query returns a query builder for Comment.
func (c *CommentClient) Query() *CommentQuery {
	return &CommentQuery{
		config: c.config,
	}
}

// Get returns a Comment entity by its id.
func (c *CommentClient) Get(ctx context.Context, id int) (*Comment, error) {
	return c.Query().Where(comment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentClient) GetX(ctx context.Context, id int) *Comment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryReview queries the review edge of a Comment.
func (c *CommentClient) QueryReview(co *Comment) *ReviewQuery {
	query := &ReviewQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, comment.ReviewTable, comment.ReviewPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CommentClient) Hooks() []Hook {
	return c.hooks.Comment
}

// DirectorClient is a client for the Director schema.
type DirectorClient struct {
	config
}

// NewDirectorClient returns a client for the Director from the given config.
func NewDirectorClient(c config) *DirectorClient {
	return &DirectorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `director.Hooks(f(g(h())))`.
func (c *DirectorClient) Use(hooks ...Hook) {
	c.hooks.Director = append(c.hooks.Director, hooks...)
}

// Create returns a create builder for Director.
func (c *DirectorClient) Create() *DirectorCreate {
	mutation := newDirectorMutation(c.config, OpCreate)
	return &DirectorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Director entities.
func (c *DirectorClient) CreateBulk(builders ...*DirectorCreate) *DirectorCreateBulk {
	return &DirectorCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Director.
func (c *DirectorClient) Update() *DirectorUpdate {
	mutation := newDirectorMutation(c.config, OpUpdate)
	return &DirectorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DirectorClient) UpdateOne(d *Director) *DirectorUpdateOne {
	mutation := newDirectorMutation(c.config, OpUpdateOne, withDirector(d))
	return &DirectorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DirectorClient) UpdateOneID(id int) *DirectorUpdateOne {
	mutation := newDirectorMutation(c.config, OpUpdateOne, withDirectorID(id))
	return &DirectorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Director.
func (c *DirectorClient) Delete() *DirectorDelete {
	mutation := newDirectorMutation(c.config, OpDelete)
	return &DirectorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DirectorClient) DeleteOne(d *Director) *DirectorDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DirectorClient) DeleteOneID(id int) *DirectorDeleteOne {
	builder := c.Delete().Where(director.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DirectorDeleteOne{builder}
}

// Query returns a query builder for Director.
func (c *DirectorClient) Query() *DirectorQuery {
	return &DirectorQuery{
		config: c.config,
	}
}

// Get returns a Director entity by its id.
func (c *DirectorClient) Get(ctx context.Context, id int) (*Director, error) {
	return c.Query().Where(director.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DirectorClient) GetX(ctx context.Context, id int) *Director {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMovies queries the movies edge of a Director.
func (c *DirectorClient) QueryMovies(d *Director) *MovieQuery {
	query := &MovieQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(director.Table, director.FieldID, id),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, director.MoviesTable, director.MoviesColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DirectorClient) Hooks() []Hook {
	return c.hooks.Director
}

// FavoriteClient is a client for the Favorite schema.
type FavoriteClient struct {
	config
}

// NewFavoriteClient returns a client for the Favorite from the given config.
func NewFavoriteClient(c config) *FavoriteClient {
	return &FavoriteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `favorite.Hooks(f(g(h())))`.
func (c *FavoriteClient) Use(hooks ...Hook) {
	c.hooks.Favorite = append(c.hooks.Favorite, hooks...)
}

// Create returns a create builder for Favorite.
func (c *FavoriteClient) Create() *FavoriteCreate {
	mutation := newFavoriteMutation(c.config, OpCreate)
	return &FavoriteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Favorite entities.
func (c *FavoriteClient) CreateBulk(builders ...*FavoriteCreate) *FavoriteCreateBulk {
	return &FavoriteCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Favorite.
func (c *FavoriteClient) Update() *FavoriteUpdate {
	mutation := newFavoriteMutation(c.config, OpUpdate)
	return &FavoriteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FavoriteClient) UpdateOne(f *Favorite) *FavoriteUpdateOne {
	mutation := newFavoriteMutation(c.config, OpUpdateOne, withFavorite(f))
	return &FavoriteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FavoriteClient) UpdateOneID(id int) *FavoriteUpdateOne {
	mutation := newFavoriteMutation(c.config, OpUpdateOne, withFavoriteID(id))
	return &FavoriteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Favorite.
func (c *FavoriteClient) Delete() *FavoriteDelete {
	mutation := newFavoriteMutation(c.config, OpDelete)
	return &FavoriteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FavoriteClient) DeleteOne(f *Favorite) *FavoriteDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FavoriteClient) DeleteOneID(id int) *FavoriteDeleteOne {
	builder := c.Delete().Where(favorite.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FavoriteDeleteOne{builder}
}

// Query returns a query builder for Favorite.
func (c *FavoriteClient) Query() *FavoriteQuery {
	return &FavoriteQuery{
		config: c.config,
	}
}

// Get returns a Favorite entity by its id.
func (c *FavoriteClient) Get(ctx context.Context, id int) (*Favorite, error) {
	return c.Query().Where(favorite.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FavoriteClient) GetX(ctx context.Context, id int) *Favorite {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FavoriteClient) Hooks() []Hook {
	return c.hooks.Favorite
}

// MovieClient is a client for the Movie schema.
type MovieClient struct {
	config
}

// NewMovieClient returns a client for the Movie from the given config.
func NewMovieClient(c config) *MovieClient {
	return &MovieClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `movie.Hooks(f(g(h())))`.
func (c *MovieClient) Use(hooks ...Hook) {
	c.hooks.Movie = append(c.hooks.Movie, hooks...)
}

// Create returns a create builder for Movie.
func (c *MovieClient) Create() *MovieCreate {
	mutation := newMovieMutation(c.config, OpCreate)
	return &MovieCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Movie entities.
func (c *MovieClient) CreateBulk(builders ...*MovieCreate) *MovieCreateBulk {
	return &MovieCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Movie.
func (c *MovieClient) Update() *MovieUpdate {
	mutation := newMovieMutation(c.config, OpUpdate)
	return &MovieUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MovieClient) UpdateOne(m *Movie) *MovieUpdateOne {
	mutation := newMovieMutation(c.config, OpUpdateOne, withMovie(m))
	return &MovieUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MovieClient) UpdateOneID(id int) *MovieUpdateOne {
	mutation := newMovieMutation(c.config, OpUpdateOne, withMovieID(id))
	return &MovieUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Movie.
func (c *MovieClient) Delete() *MovieDelete {
	mutation := newMovieMutation(c.config, OpDelete)
	return &MovieDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MovieClient) DeleteOne(m *Movie) *MovieDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MovieClient) DeleteOneID(id int) *MovieDeleteOne {
	builder := c.Delete().Where(movie.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MovieDeleteOne{builder}
}

// Query returns a query builder for Movie.
func (c *MovieClient) Query() *MovieQuery {
	return &MovieQuery{
		config: c.config,
	}
}

// Get returns a Movie entity by its id.
func (c *MovieClient) Get(ctx context.Context, id int) (*Movie, error) {
	return c.Query().Where(movie.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MovieClient) GetX(ctx context.Context, id int) *Movie {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDirector queries the director edge of a Movie.
func (c *MovieClient) QueryDirector(m *Movie) *DirectorQuery {
	query := &DirectorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(director.Table, director.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, movie.DirectorTable, movie.DirectorColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReviews queries the reviews edge of a Movie.
func (c *MovieClient) QueryReviews(m *Movie) *ReviewQuery {
	query := &ReviewQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, movie.ReviewsTable, movie.ReviewsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryActor queries the actor edge of a Movie.
func (c *MovieClient) QueryActor(m *Movie) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, movie.ActorTable, movie.ActorPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MovieClient) Hooks() []Hook {
	return c.hooks.Movie
}

// ReviewClient is a client for the Review schema.
type ReviewClient struct {
	config
}

// NewReviewClient returns a client for the Review from the given config.
func NewReviewClient(c config) *ReviewClient {
	return &ReviewClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `review.Hooks(f(g(h())))`.
func (c *ReviewClient) Use(hooks ...Hook) {
	c.hooks.Review = append(c.hooks.Review, hooks...)
}

// Create returns a create builder for Review.
func (c *ReviewClient) Create() *ReviewCreate {
	mutation := newReviewMutation(c.config, OpCreate)
	return &ReviewCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Review entities.
func (c *ReviewClient) CreateBulk(builders ...*ReviewCreate) *ReviewCreateBulk {
	return &ReviewCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Review.
func (c *ReviewClient) Update() *ReviewUpdate {
	mutation := newReviewMutation(c.config, OpUpdate)
	return &ReviewUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReviewClient) UpdateOne(r *Review) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReview(r))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReviewClient) UpdateOneID(id int) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReviewID(id))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Review.
func (c *ReviewClient) Delete() *ReviewDelete {
	mutation := newReviewMutation(c.config, OpDelete)
	return &ReviewDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ReviewClient) DeleteOne(r *Review) *ReviewDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ReviewClient) DeleteOneID(id int) *ReviewDeleteOne {
	builder := c.Delete().Where(review.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReviewDeleteOne{builder}
}

// Query returns a query builder for Review.
func (c *ReviewClient) Query() *ReviewQuery {
	return &ReviewQuery{
		config: c.config,
	}
}

// Get returns a Review entity by its id.
func (c *ReviewClient) Get(ctx context.Context, id int) (*Review, error) {
	return c.Query().Where(review.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReviewClient) GetX(ctx context.Context, id int) *Review {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMovie queries the movie edge of a Review.
func (c *ReviewClient) QueryMovie(r *Review) *MovieQuery {
	query := &MovieQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, id),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, review.MovieTable, review.MovieColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Review.
func (c *ReviewClient) QueryUser(r *Review) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, review.UserTable, review.UserColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComments queries the comments edge of a Review.
func (c *ReviewClient) QueryComments(r *Review) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, review.CommentsTable, review.CommentsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ReviewClient) Hooks() []Hook {
	return c.hooks.Review
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryReviews queries the reviews edge of a User.
func (c *UserClient) QueryReviews(u *User) *ReviewQuery {
	query := &ReviewQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ReviewsTable, user.ReviewsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
