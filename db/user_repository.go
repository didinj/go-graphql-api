package db

import (
	"context"

	"github.com/didinj/go-graphql-api/graph/model"
)

func GetUsers(ctx context.Context) ([]*model.DBUser, error) {
	rows, err := Pool.Query(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.DBUser
	for rows.Next() {
		var u model.DBUser
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func GetUserByID(ctx context.Context, id int) (*model.DBUser, error) {
	var u model.DBUser
	err := Pool.QueryRow(ctx, "SELECT id, name, email FROM users WHERE id=$1", id).
		Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func CreateUser(ctx context.Context, name, email string) (*model.DBUser, error) {
	var u model.DBUser
	err := Pool.QueryRow(ctx,
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email",
		name, email,
	).Scan(&u.ID, &u.Name, &u.Email)

	if err != nil {
		return nil, err
	}
	return &u, nil
}
