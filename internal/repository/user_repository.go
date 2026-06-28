package repository

import (
	"database/sql"
	"time"

	
)

type User struct {
	ID             int
	Username       string
	PasswordHash   string
	TwoFASecret    string
	TwoFAEnabled   bool
	FailedAttempts int
	LockedUntil    sql.NullTime
	CreatedAt      time.Time
	LastLogin      sql.NullTime
}

type Session struct {
	Token     string
	Username  string
	ExpiresAt time.Time
}