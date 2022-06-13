package entity

import "time"

type TokenEntity struct {
	UserId      uint
	ExpiredTime time.Time
}
