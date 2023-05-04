package repository

import "github.com/jackc/pgx/v5"

const (
	userTable       = "users"
	sessionsTable   = "sessions"
	chatsTable      = "chats"
	chatsUsersTable = "chats_users"
	messageTable    = "message"
	walletTable     = "wallet"
)

var ErrNoRows = pgx.ErrNoRows
