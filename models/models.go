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

type Transaction struct {
	bun.BaseModel `bun:"table:transaction"`

	Id        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Amount    int       `bun:"amount,notnull"`
	OrderId   uuid.UUID `bun:"order_id,notnull,type:uuid"`
	ServiceId uuid.UUID `bun:"service_id,notnull,type:uuid"`
	Status    string    `bun:"status,notnull,type:transaction_status"`
	UserId    uuid.UUID `bun:"user_id,notnull,type:uuid"`
}
