package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.76

import (
	"context"
	"fmt"
	"strconv"

	"github.com/didinj/go-graphql-api/db"
	"github.com/didinj/go-graphql-api/graph/generated"
	"github.com/didinj/go-graphql-api/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string) (*model.User, error) {
	dbUser, err := db.CreateUser(ctx, name, email)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    fmt.Sprintf("%d", dbUser.ID),
		Name:  dbUser.Name,
		Email: dbUser.Email,
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	dbUsers, err := db.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var gqlUsers []*model.User
	for _, u := range dbUsers {
		gqlUsers = append(gqlUsers, &model.User{
			ID:    fmt.Sprintf("%d", u.ID),
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return gqlUsers, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format")
	}

	dbUser, err := db.GetUserByID(ctx, intID)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    fmt.Sprintf("%d", dbUser.ID),
		Name:  dbUser.Name,
		Email: dbUser.Email,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
