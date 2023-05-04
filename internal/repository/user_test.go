package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/Qwepo/InCryipt/internal"
	"github.com/Qwepo/InCryipt/internal/models"
	"github.com/Qwepo/InCryipt/internal/types"
	"github.com/Qwepo/InCryipt/pkg/database/postgres"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	ctx := context.Background()
	pgconf := internal.PostgresConfig{
		Database:       "test",
		Password:       "test",
		Username:       "test",
		Address:        "127.0.0.1",
		Port:           "5431",
		MigrationsPaht: "../../../cmd/app/migrations",
	}
	conf := internal.Config{
		Postgres: pgconf,
	}
	conn, err := postgres.NewClient(ctx, &conf)
	assert.NoError(t, err)

	err = conn.Migrate()
	assert.NoError(t, err)
	repo, err := NewRepo(ctx, conn.Pool)
	assert.NoError(t, err)

	testUserCreateAndGet(t, ctx, repo)
	testUserUpdatebById(t, ctx, repo)

}

func testUserCreateAndGet(t *testing.T, ctx context.Context, repo Repository) {
	u := models.User{
		Name:     "test",
		Email:    "test@mail.ru",
		Password: "test",
	}
	u.Created_at.SetNow()
	err := repo.UserCreate(ctx, &u)
	assert.NoError(t, err)

	user, err := repo.GetUserByID(ctx, types.PK(u.Id))
	assert.NoError(t, err)
	assert.Equal(t, u.Id, user.Id)
	assert.Equal(t, u.Name, user.Name)

}

func testUserUpdatebById(t *testing.T, ctx context.Context, repo Repository) {
	u := models.User{
		Name:     "test",
		Email:    "test@mail.ru",
		Password: "test",
	}
	u.Created_at.SetNow()
	err := repo.UserCreate(ctx, &u)
	assert.NoError(t, err)

	u.Name = "test2"
	u.Email = "test2@gmail"
	u.Password = "qqqq"
	err = repo.UserUpdateById(ctx, &u)
	assert.NoError(t, err)

	user, err := repo.GetUserByID(ctx, types.PK(u.Id))
	assert.NoError(t, err)

	assert.Equal(t, u.Id, user.Id)
	assert.Equal(t, u.Name, user.Name)

}

func testUserList(t *testing.T, ctx context.Context, repo Repository) {
	u1 := models.User{
		Name:     "test",
		Email:    "test@mail.ru",
		Password: "test",
	}
	u2 := models.User{
		Name:     "test2",
		Email:    "test2@mail.ru",
		Password: "test2",
	}
	u1.Created_at.SetNow()
	u2.Created_at.SetNow()

	err := repo.UserCreate(ctx, &u1)
	assert.NoError(t, err)

	err = repo.UserCreate(ctx, &u2)
	assert.NoError(t, err)

	users, err := repo.UserList(ctx)
	assert.NoError(t, err)

	for _, user := range users {
		if user.Id == u1.Id {
			assert.Equal(t, user.Id, u1.Id)
			assert.Equal(t, user.Name, u1.Name)

		}
		if user.Id == u2.Id {
			assert.Equal(t, user.Id, u2.Id)
			assert.Equal(t, user.Name, u2.Name)

		} else {
			err := errors.New("no user in list")
			assert.NoError(t, err)
		}
	}
}
