package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Go/graph/generated"
	"Go/graph/model"
	"context"
)

func (r *queryResolver) Login(ctx context.Context, userEmail string, password string) (*model.User, error) {
	users := new(model.User)
	err := r.DB.Model(users).Where("email = ? AND password = ?", userEmail, password).Select()
	if err != nil{
		return nil, err
	}
	return users, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
