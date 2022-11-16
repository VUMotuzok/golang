package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserModel struct {
	bun.BaseModel `bun:"table:user"`

	Id     uuid.UUID `bun:"id,pk,type:uuid"`
	Amount int       `bun:"amount,notnull"`
}
