package request

import (
	"fmt"
	"time"
)

type AnggotaRequest struct {
	Id           uint      `json:"id" db:"ID"`
	Nama         string    `json:"nama" binding:"required" db:"NAMA"`
	AngkaFavorit int       `json:"angkaFavorit" binding:"required" db:"ANGKA_FAVORIT"`
	TanggalLahir time.Time `json:"tanggalLahir" binding:"required" db:"TANGGAL_LAHIR"`
	CreatedAt    time.Time `json:"createdAt" db:"CREATED_AT"`
}

func (a *AnggotaRequest) ToString() string {
	return fmt.Sprintf("Id: %d, Nama: %s, Angka Favorit: %d, Tanggal Lahir: %s", a.Id, a.Nama, a.AngkaFavorit, a.TanggalLahir)
}
