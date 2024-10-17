package services

import (
	"errors"
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/models"
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/repository"
	"github.com/doryngal/CourseFlow/backend/internal/shared/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (a *AuthService) Register(user *models.User) error {
	isRegistered, err := a.Repo.IsEmailRegistered(user.Email)
	if err != nil {
		return err
	}
	if isRegistered {
		return errors.New("email already registered")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	return a.Repo.CreateUser(user)
}

func (a *AuthService) Login(email, password string) (string, error) {
	user, err := a.Repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("password is not correct")
	}

	token, err := utils.GenerateNewAccessToken(*user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthService) Logout(token string) error {
	//expiration := time.Hour * 24
	//err := a.Repo.AddTokenToBlacklist(token, expiration)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (a *AuthService) GetProfile(userId int) (*models.User, error) {
	user, err := a.Repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
