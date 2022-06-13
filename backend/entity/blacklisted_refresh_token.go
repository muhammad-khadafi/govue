package entity

import "time"

type BlacklistedRefreshToken struct {
	Id        uint      `json:"id" db:"ID"`
	Token     string    `json:"token" db:"TOKEN"`
	CreatedAt time.Time `json:"createdAt" db:"CREATED_AT"`
	ExpiredAt time.Time `json:"expiredAt" db:"EXPIRED_AT"`
}
