package repository

import (
	"database/sql"
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/models"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(userId int) (*models.User, error)
	IsEmailRegistered(email string) (bool, error)
	//AddTokenToBlacklist(token string, expiration time.Duration) error
	//IsTokenBlacklisted(token string) (bool, error)
}

type PqUserRepository struct {
	DB *sql.DB
	//Redis *redis.Client
	//Ctx   context.Context
}

func NewPgUserRepository(db *sql.DB) *PqUserRepository {
	return &PqUserRepository{DB: db}
}

func (r *PqUserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (email, password_hash, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.Exec(query, user.Email, user.PasswordHash, user.Name, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *PqUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = $1`
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *PqUserRepository) GetUserById(userId int) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.DB.QueryRow(query, userId).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *PqUserRepository) IsEmailRegistered(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := r.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

//func (r *PqUserRepository) AddTokenToBlacklist(token string, expiration time.Duration) error {
//	return r.Redis.Set(r.Ctx, token, "blacklisted", expiration).Err()
//}
//
//func (r *PqUserRepository) IsTokenBlacklisted(token string) (bool, error) {
//	result, err := r.Redis.Get(r.Ctx, token).Result()
//	if err == redis.Nil {
//		return false, nil
//	} else if err != nil {
//		return false, err
//	}
//	return result == "blacklisted", nil
//}
