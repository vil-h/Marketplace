package User

import (
	"context"
	"errors"
	"net/mail"
	"time"
	"unicode/utf8"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo   *UserRepository
	jwtSec string
}

func NewUserService(repo *UserRepository, jwtSec string) *UserService {
	return &UserService{repo: repo, jwtSec: jwtSec}
}

func (u *UserService) Register(ctx context.Context, email string, password string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}
	if utf8.RuneCountInString(password) < 3 {
		return errors.New("password must be at least 3 characters")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Error password hash")
	}
	return u.repo.CreateUser(ctx, email, string(hash))
}

func (u *UserService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := u.repo.LoginUser(ctx, email)
	if err != nil {
		return "", errors.New("No users Login")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("No login error password or email")
	}

	claims := jwt.MapClaims{
		"sub": user.Id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToke, err := token.SignedString([]byte(u.jwtSec))
	if err != nil {
		return "", err
	}
	return signedToke, nil
}

func (u *UserService) IsSeller(ctx context.Context, value int) (string, error) {
	role, err := u.repo.CheckRole(ctx, value)
	if err != nil {
		return "", err
	}
	return role, nil

}

func (u *UserService) Me(ctx context.Context, userId int) (role string, email string, err error) {
	role, email, err = u.repo.GetMe(ctx, userId)
	if err != nil {
		return "", "", err
	}
	return role, email, nil

}
