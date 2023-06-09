package usersRepositories

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/peedans/GoEcommerce/modules/users"
	"github.com/peedans/GoEcommerce/modules/users/usersPatterns"
	"time"
)

type IUserRepository interface {
	InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error)
	FindOneUserByEmail(email string) (*users.UserCredentialCheck, error)
	InsertOauth(req *users.UserPassport) error
	FineOneOauth(refreshToken string) (*users.Oauth, error)
	UpdateOauth(req *users.UserToken) error
	GetProfile(userId string) (*users.User, error)
	DeleteOauth(oauthId string) error
}

type usersRepository struct {
	db *sqlx.DB
}

func UsersRepository(db *sqlx.DB) IUserRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error) {
	result := usersPatterns.InsertUser(r.db, req, isAdmin)

	var err error
	if isAdmin {
		result, err = result.Admin()
		if err != nil {
			return nil, err
		}
	} else {
		result, err = result.Customer()
		if err != nil {
			return nil, err
		}
	}

	user, err := result.Result()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *usersRepository) FindOneUserByEmail(email string) (*users.UserCredentialCheck, error) {
	query := `
	SELECT
		"id",
		"email",
		"password",
		"username",
		"role_id"
	FROM "users"
	WHERE "email" = $1;`
	user := new(users.UserCredentialCheck)
	if err := r.db.Get(user, query, email); err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *usersRepository) InsertOauth(req *users.UserPassport) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
	INSERT INTO "oauth" (
		"user_id",
		"refresh_token",
		"access_token",
    	"is_deleted",
		"created_by",
	    "created_at",
 	    "updated_by",
 	    "updated_at"
)
	VALUES ($1, $2, $3, false, $4, NOW(), $4, NOW())
	RETURNING "id";`

	if err := r.db.QueryRowContext(
		ctx,
		query,
		req.User.Id,
		req.Token.RefreshToken,
		req.Token.AccessToken,
		req.User.CreatedBy,
	).Scan(&req.Token.Id); err != nil {
		return fmt.Errorf("insert oauth failed: %v", err)
	}
	return nil
}

func (r *usersRepository) FineOneOauth(refreshToken string) (*users.Oauth, error) {
	//fmt.Println(refreshToken)
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbXMiOnsiaWQiOiJVMDAwMDA2Iiwicm9sZSI6MX0sImlzcyI6ImdvRWNvbW1lcmNlLWFwaSIsInN1YiI6InJlZnJlc2gtdG9rZW4iLCJhdWQiOlsiY3VzdG9tZXIiLCJhZG1pbiJdLCJleHAiOjE2ODY4MzI5NzksIm5iZiI6MTY4NjIyODE3OSwiaWF0IjoxNjg2MjI4MTc5fQ.8C5TfsMNIx_uQIm2zrX8ttA3pE4jX7QdRRMlr5_nqEw
	query := `
	SELECT
    	"id",
		"user_id"
	FROM "oauth"
	WHERE "refresh_token" = $1;`
	oauth := new(users.Oauth)

	if err := r.db.Get(oauth, query, refreshToken); err != nil {
		return nil, fmt.Errorf("oauth not found")
	}
	fmt.Println("UserId", oauth.UserId)
	fmt.Println("Id", oauth.Id)
	return oauth, nil
}

func (r *usersRepository) UpdateOauth(req *users.UserToken) error {
	query := `
	UPDATE "oauth" SET
		   "access_token" = :access_token,
		   "refresh_token" = :refresh_token
    WHERE "id" = :id`
	if _, err := r.db.NamedExecContext(context.Background(), query, req); err != nil {
		return fmt.Errorf("update oauth failed: %v", err)
	}
	return nil
}

func (r *usersRepository) GetProfile(userId string) (*users.User, error) {
	query := `
	SELECT
		"id",
		"email",
		"username",
		"role_id"
	FROM "users"
	WHERE "id" = $1;`
	profile := new(users.User)
	if err := r.db.Get(profile, query, userId); err != nil {
		return nil, fmt.Errorf("get user failed: %v", err)
	}

	return profile, nil
}

func (r *usersRepository) DeleteOauth(oauthId string) error {
	query := `DELETE FROM "oauth" WHERE "id" = $1;`

	if _, err := r.db.ExecContext(context.Background(), query, oauthId); err != nil {
		return fmt.Errorf("oauth not found")
	}
	return nil
}
