package User

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, email string, hash string) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO users(email, password_hash)
			VALUES ($1, $2)`, email, hash)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) LoginUser(ctx context.Context, email string) (*User, error) {
	var users User
	err := r.db.QueryRow(
		ctx,
		`SELECT id, email, password_hash
			FROM users
			WHERE email = $1`, email).Scan(
		&users.Id,
		&users.Email,
		&users.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &users, err
}

func (r *UserRepository) CheckRole(ctx context.Context, value int) (role string, err error) {
	var users User
	err = r.db.QueryRow(
		ctx,
		`SELECT role
				FROM users
				WHERE id = $1`, value).Scan(
		&users.Role)
	if err != nil {
		return "", err
	}
	return users.Role, nil
}

func (r *UserRepository) GetMe(ctx context.Context, userId int) (role string, email string, err error) {
	var users User
	err = r.db.QueryRow(
		ctx,
		`SELECT email, role
			FROM users
			WHERE id = $1`, userId).Scan(&users.Email, &users.Role)
	if err != nil {
		return "", "", err
	}
	return users.Role, users.Email, nil

}
