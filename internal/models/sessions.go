package models

import "github.com/Qwepo/InCryipt/internal/types"

type Sessions struct {
	Id        types.PK   `json:"id"`
	UserId    int64      `json:"user_id"`
	Token     string     `json:"token"`
	CreatedAt types.Unix `json:"created_at"`
	ExpiresAt types.Unix `json:"expires_at"`
}
