package repository

import (
	"database/sql"
	"log"
)

type IAuthRepository struct {
	UserRepo UserRepository
}

func NewAuthRepository(db *sql.DB) *IAuthRepository {
	return &IAuthRepository{
		UserRepo: NewPgUserRepository(db),
	}
}

func IAuthRepositories(db *sql.DB) *IAuthRepository {
	repos := NewAuthRepository(db)
	log.Println("Repository have been initialized")
	return repos
}
