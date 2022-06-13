package entity

//time format harus RFC3339, ex: 2017-01-08T00:00:00Z
import (
	"time"
)

type Institusi struct {
	Id        uint      `json:"id" db:"ID"`
	Nama      string    `json:"nama" db:"NAMA"`
	Alamat    string    `json:"alamat" db:"ALAMAT"`
	Email     string    `json:"email" db:"EMAIL"`
	CreatedAt time.Time `json:"createdAt" db:"CREATED_AT"`
}
