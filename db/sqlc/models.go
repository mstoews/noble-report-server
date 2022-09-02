// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type ExternalRef struct {
	TrdRecordno      int32  `json:"trd_recordno"`
	ExtRefExtreftype string `json:"ext_ref_extreftype"`
	ExtRefExtref     string `json:"ext_ref_extref"`
}

type RefDate struct {
	TrdRecordno     int32     `json:"trd_recordno"`
	Datetype        string    `json:"datetype"`
	Datewil         time.Time `json:"datewil"`
	Refdatetime     string    `json:"refdatetime"`
	Dateversionuser string    `json:"dateversionuser"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Trade struct {
	TrdRecordno         int32     `json:"trd_recordno"`
	TrdGlosstraderef    int32     `json:"trd_glosstraderef"`
	TrdVersiono         int32     `json:"trd_versiono"`
	TrdOrigin           string    `json:"trd_origin"`
	TrdTradetype        string    `json:"trd_tradetype"`
	TrdSettlementstatus string    `json:"trd_settlementstatus"`
	TrdTradestatus      string    `json:"trd_tradestatus"`
	TrdOriginversion    int32     `json:"trd_originversion"`
	TrdDate             time.Time `json:"trd_date"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
