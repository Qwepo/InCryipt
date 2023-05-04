package models

import (
	"github.com/Qwepo/InCryipt/internal/types"
)

type User struct {
	Id         int64      `json:"id"`
	Name       string     `json:"username"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Created_at types.Unix `json:"created_at"`
	Updated_at types.Unix `json:"updated_at"`
}
