package models

import (
	"time"
)

type Base struct {
	Id        int       `json:"id,string" sql:",pk"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
