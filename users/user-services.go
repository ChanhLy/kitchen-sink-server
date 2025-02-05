package users

import (
	"context"
	"database/sql"
	"errors"
	database "go-server/db"
)

func GetUsers(ctx context.Context) ([]database.User, error) {
	queries := database.GetQueries()
	users, err := queries.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(id int64, ctx context.Context) (database.User, error) {
	queries := database.GetQueries()
	user, err := queries.GetUserById(ctx, id)

	if errors.Is(err, sql.ErrNoRows) {
		return database.User{}, nil
	}

	if err != nil {
		return database.User{}, err
	}

	return user, nil
}

func CreateUser(data database.CreateUserParams, ctx context.Context) (database.User, error) {
	queries := database.GetQueries()

	user, err := queries.CreateUser(ctx, data)
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}
