package user

import "github.com/google/uuid"

type User struct {
	ID *uuid.UUID `json:"id"`
}

type AccountToken struct {
	UserID       *uuid.UUID `json:"id"`
	PrivateToken *string    `json:"privateToken"`
	ItemID       *string    `json:"itemId"`
}
