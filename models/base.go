package models

import (
	"time"
)

type Base struct {
	Id        int       `json:"id" sql:",pk"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
