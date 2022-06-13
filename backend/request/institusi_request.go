package request

//time format harus RFC3339, ex: 2017-01-08T00:00:00Z
import (
	"fmt"
	"time"
)

type InstitusiRequest struct {
	Id        uint      `json:"id" db:"ID"`
	Nama      string    `json:"nama" binding:"required" db:"NAMA"`
	Alamat    string    `json:"alamat" db:"ALAMAT"`
	Email     string    `json:"email" binding:"email" db:"EMAIL"`
	CreatedAt time.Time `json:"createdAt" db:"CREATED_AT"`
}

func (i *InstitusiRequest) ToString() string {
	return fmt.Sprintf("Id: %d, Nama: %s, Alamat: %s, Email: %s", i.Id, i.Nama, i.Alamat, i.Email)
}
