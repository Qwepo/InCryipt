package service

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/Qwepo/InCryipt/internal/models"
	"github.com/Qwepo/InCryipt/internal/repository"
	"github.com/Qwepo/InCryipt/internal/types"
	"github.com/Qwepo/InCryipt/pkg/utils"
)

var ErrUserNotFound = errors.New("user not found")
var ErrPasswordMismatch = errors.New("password mismatch")

type sAuth interface {
	CreateSessions(ctx context.Context, password, email string) (*models.Sessions, error)
}

func (s *service) CreateSessions(ctx context.Context, email, password string) (*models.Sessions, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrNoRows) {
			return nil, ErrUserNotFound
		}
	}

	if err = utils.ComparePasswords(user.Password, password); err != nil {
		return nil, err
	}

	var sessions models.Sessions
	token, err := utils.GenerateRandomToken()
	if err != nil {
		return nil, err
	}

	sessions.UserId = user.Id
	sessions.Token = token
	sessions.CreatedAt.SetNow()
	sessions.ExpiresAt.Add(s.conf.Sessions.LifeHours)

	jsessions, err := json.Marshal(&sessions)
	if err != nil {
		return nil, err
	}
	if err = s.redis.HSet(ctx, key(types.PK(sessions.UserId)), sessions.Token, jsessions).Err(); err != nil {
		return nil, err
	}
	return &sessions, nil
}

func (s *service) GetSessions(ctx context.Context, id types.PK, token string) (*models.Sessions, error) {
	jsessions, err := s.redis.HGet(ctx, key(id), token).Result()
	if err != nil {
		return nil, err
	}

	var sessions models.Sessions
	if err = json.Unmarshal([]byte(jsessions), &sessions); err != nil {
		return nil, err
	}
	return &sessions, nil
}

func (s *service) UpdateSession(ctx context.Context, id types.PK, token string) error {
	session, err := s.redis.HGet(ctx, key(id), token).Result()
	if err != nil {
		return err
	}
	var sessionData models.Sessions
	if err = json.Unmarshal([]byte(session), &session); err != nil {
		return err
	}

	sessionData.ExpiresAt.Add(s.conf.Sessions.LifeHours)
	sessionJSON, err := json.Marshal(&session)
	if err != nil {
		return err
	}

	return s.redis.HSet(ctx, key(id), token, string(sessionJSON)).Err()

}

func (s *service) DeleteSession(ctx context.Context, id types.PK, token string) error {
	return s.redis.HDel(ctx, key(id), token).Err()
}

func key(id types.PK) string {
	return "sessions:" + strconv.FormatInt(int64(id), 10)
}
