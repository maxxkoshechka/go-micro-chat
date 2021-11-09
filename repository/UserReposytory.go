package repository

import (
	"chat/models"
	"context"
	"fmt"
	pgx "github.com/jackc/pgx/v4"
	"os"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepo (db *pgx.Conn) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Select(query string) (*models.User, error)  {
	var name string
	var email string
	err = r.db.QueryRow(context.Background(), "select name, email from user where id=$1", 1).Scan(&name, &email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return &models.User{Name: name, Email: email}, nil
}

func (r *UserRepo) Create(user *models.User) error  {
	return nil
}

func (r *UserRepo) Update(user *models.User, query string) (*models.User, error)  {
	return &models.User{}, nil
}

func (r *UserRepo) Delete(user *models.User) error  {
	return nil
}
