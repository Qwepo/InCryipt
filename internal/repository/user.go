package repository

import (
	"context"
	"fmt"

	"github.com/Qwepo/InCryipt/internal/models"
	"github.com/Qwepo/InCryipt/internal/types"
)

type dbUser interface {
	UserCreate(ctx context.Context, u *models.User) error
	GetUserByID(ctx context.Context, id types.PK) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UserUpdateById(ctx context.Context, user *models.User) error
	UserList(ctx context.Context) (users []*models.User, err error)
}

func (db *client) UserCreate(ctx context.Context, u *models.User) error {
	u.Created_at.SetNow()
	query := fmt.Sprintf(`
	INSERT INTO %s (
			username, 
			email, 
			password, 
			created_at, 
			updated_at) 
		VALUES
			($1, $2, $3, $4, 0) 
		RETURNING id`, userTable)

	row := db.QueryRow(ctx, query, u.Name, u.Email, u.Password, u.Created_at)
	return row.Scan(&u.Id)

}

func (db *client) GetUserByID(ctx context.Context, id types.PK) (*models.User, error) {
	var user models.User
	query := fmt.Sprintf(`
	SELECT 
		id, 
		username, 
		created_at, 
		updated_at 
	FROM %s 
	WHERE id = $1`, userTable)

	row := db.QueryRow(ctx, query, id)
	err := row.Scan(&user.Id, &user.Name, &user.Created_at, &user.Updated_at)
	if err != nil {
		return nil, err
	}
	return &user, err
}
func (db *client) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := fmt.Sprintf(`
	SELECT 
		id, 
		username, 
		email,
		password
		created_at, 
		updated_at 
	FROM %s 
	WHERE email = $1`, userTable)

	row := db.QueryRow(ctx, query, email)
	err := row.Scan(&user.Id, &user.Name, &user.Created_at, &user.Updated_at)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (db *client) UserUpdateById(ctx context.Context, user *models.User) error {
	user.Updated_at.SetNow()
	query := fmt.Sprintf(`
	UPDATE %s SET 
		username = $1, 
		email = $2, 
		password = $3, 
		updated_at = $4  
	WHERE id = $5`, userTable)

	_, err := db.Exec(ctx, query, user.Name, user.Email, user.Password, user.Updated_at, user.Id)
	return err

}

func (db *client) UserList(ctx context.Context) (users []*models.User, err error) {
	query := fmt.Sprintf(`
	SELECT 
		id, 
		username, 
		created_at, 
		updated_at   F
	FROM %s `, userTable)

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.Id, &u.Name, &u.Created_at, &u.Updated_at)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return
}
