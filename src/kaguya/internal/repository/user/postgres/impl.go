package postgres

import (
	"context"
	"database/sql"
	"errors"
	userDomain "github.com/aryahmph/wattpad-clone/src/kaguya/internal/domain/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type postgresUserRepositoryImpl struct {
	db *sql.DB
}

func NewPostgresUserRepositoryImpl(db *sql.DB) postgresUserRepositoryImpl {
	return postgresUserRepositoryImpl{db: db}
}

func (repository postgresUserRepositoryImpl) Insert(ctx context.Context, user userDomain.User) (rid string, err error) {
	row := repository.db.QueryRowContext(ctx,
		"INSERT INTO users(username, email, password_hash) VALUES ($1, $2, $3) RETURNING id;",
		user.Username,
		user.Email,
		user.PasswordHash,
	)
	err = row.Scan(&rid)
	if errors.Is(err, sql.ErrNoRows) {
		return rid, status.Error(codes.NotFound, "user not found")
	}
	return rid, err
}

func (repository postgresUserRepositoryImpl) FindByID(ctx context.Context, id string) (user userDomain.User, err error) {
	row := repository.db.QueryRowContext(ctx,
		"SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE id = $1 LIMIT 1;",
		id)

	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return user, status.Error(codes.NotFound, "user not found")
	}
	return user, err
}

func (repository postgresUserRepositoryImpl) FindByUsername(ctx context.Context, username string) (user userDomain.User, err error) {
	row := repository.db.QueryRowContext(ctx,
		"SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE username = $1 LIMIT 1;",
		username)

	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return user, status.Error(codes.NotFound, "user not found")
	}
	return user, err
}

func (repository postgresUserRepositoryImpl) FindByEmail(ctx context.Context, email string) (user userDomain.User, err error) {
	row := repository.db.QueryRowContext(ctx,
		"SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE email = $1 LIMIT 1;",
		email)

	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return user, status.Error(codes.NotFound, "user not found")
	}
	return user, err
}

func (repository postgresUserRepositoryImpl) UpdatePassword(ctx context.Context, id string, password string) (rid string, err error) {
	//TODO implement me
	panic("implement me")
}
