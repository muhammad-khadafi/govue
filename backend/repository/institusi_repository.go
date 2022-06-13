package repository

import (
	"backend/entity"
	"backend/request"

	"github.com/jmoiron/sqlx"
)

type InstitusiRepository interface {
	FindByID(id uint) (entity.Institusi, error)
	List() ([]entity.Institusi, error)
	Insert(institusi request.InstitusiRequest) error
	Update(institusi request.InstitusiRequest) error
	DeleteById(id uint) int
}

type institusiRepository struct {
	DB *sqlx.DB
}

func NewInstitusiRepository(db *sqlx.DB) *institusiRepository {
	return &institusiRepository{DB: db}
}

func (repository *institusiRepository) FindByID(id uint) (entity.Institusi, error) {
	institusi := entity.Institusi{}
	script := `select * from institusi where id = :1`
	err := repository.DB.Get(&institusi, script, id)
	return institusi, err
}

func (repository *institusiRepository) List() ([]entity.Institusi, error) {
	institusi := []entity.Institusi{}
	script := `select * from institusi`
	err := repository.DB.Select(&institusi, script)
	return institusi, err
}

func (repository *institusiRepository) Insert(institusi request.InstitusiRequest) error {
	script := `INSERT INTO institusi (nama, alamat, email, created_at) VALUES (:NAMA, :ALAMAT, :EMAIL, :CREATED_AT)`
	_, err := repository.DB.NamedExec(script, &institusi)
	return err
}

func (repository *institusiRepository) Update(institusi request.InstitusiRequest) error {
	script := `UPDATE institusi SET nama = :NAMA, alamat = :ALAMAT, email = :EMAIL WHERE ID = :ID`
	_, err := repository.DB.NamedExec(script, &institusi)
	return err
}

func (repository *institusiRepository) DeleteById(id uint) int {
	script := `DELETE FROM institusi WHERE id = :1`
	del := repository.DB.MustExec(script, id)
	count, _ := del.RowsAffected()
	return int(count)
}
