package repository

import (
	"backend/entity"
	"backend/request"
	"github.com/jmoiron/sqlx"
	"time"
)

type AnggotaRepository interface {
	FindByID(id uint) (entity.Anggota, error)
	List() ([]entity.Anggota, error)
	Insert(anggota request.AnggotaRequest) error
	Update(anggota request.AnggotaRequest) error
	DeleteById(id uint) int
}

type anggotaRepository struct {
	DB *sqlx.DB
}

func NewAnggotaRepository(db *sqlx.DB) *anggotaRepository {
	return &anggotaRepository{DB: db}
}

func (repository *anggotaRepository) FindByID(id uint) (entity.Anggota, error) {
	anggota := entity.Anggota{}
	script := `select * from anggota where id = :1`
	err := repository.DB.Get(&anggota, script, id)
	return anggota, err
}

func (repository *anggotaRepository) List() ([]entity.Anggota, error) {
	anggota := []entity.Anggota{}
	script := `select * from anggota`
	err := repository.DB.Select(&anggota, script)
	return anggota, err
}

func (repository *anggotaRepository) Insert(anggota request.AnggotaRequest) error {
	anggota.CreatedAt = time.Now()
	script := `INSERT INTO anggota (nama, angka_favorit, tanggal_lahir, created_at) VALUES (:NAMA, :ANGKA_FAVORIT, :TANGGAL_LAHIR, :CREATED_AT)`
	_, err := repository.DB.NamedExec(script, &anggota)
	return err
}

func (repository *anggotaRepository) Update(anggota request.AnggotaRequest) error {
	script := `UPDATE anggota SET nama = :NAMA, angka_favorit = :ANGKA_FAVORIT, tanggal_lahir = :TANGGAL_LAHIR WHERE ID = :ID`
	_, err := repository.DB.NamedExec(script, &anggota)
	return err
}

func (repository *anggotaRepository) DeleteById(id uint) int {
	script := `DELETE FROM anggota WHERE id = :1`
	del := repository.DB.MustExec(script, id)
	count, _ := del.RowsAffected()
	return int(count)
}
