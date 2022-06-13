package entity

//time format harus RFC3339, ex: 2017-01-08T00:00:00Z
import (
	"time"
)

type Anggota struct {
	Id           uint      `json:"id" db:"ID"`
	Nama         string    `json:"nama" db:"NAMA"`
	AngkaFavorit int       `json:"angkaFavorit" db:"ANGKA_FAVORIT"`
	TanggalLahir time.Time `json:"tanggalLahir" db:"TANGGAL_LAHIR"`
	CreatedAt    time.Time `json:"createdAt" db:"CREATED_AT"`
}
