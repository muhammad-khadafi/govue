package service

import (
	"backend/entity"
	"backend/repository"
	"backend/request"
)

type AnggotaService interface {
	FindAnggotaByID(id uint) (entity.Anggota, error)
	ListAnggota() ([]entity.Anggota, error)
	InsertAnggota(anggota request.AnggotaRequest) error
	UpdateAnggota(anggota request.AnggotaRequest) error
	DeleteAnggotaById(id uint) int
}

type anggotaService struct {
	anggotaRepository repository.AnggotaRepository
}

func NewAnggotaService(anggotaRepository repository.AnggotaRepository) *anggotaService {
	return &anggotaService{anggotaRepository}
}

func (service *anggotaService) FindAnggotaByID(id uint) (entity.Anggota, error) {
	anggota, err := service.anggotaRepository.FindByID(id)
	return anggota, err
}

func (service *anggotaService) ListAnggota() ([]entity.Anggota, error) {
	anggota, err := service.anggotaRepository.List()
	return anggota, err
}

func (service *anggotaService) InsertAnggota(anggota request.AnggotaRequest) error {
	err := service.anggotaRepository.Insert(anggota)
	return err
}

func (service *anggotaService) UpdateAnggota(anggota request.AnggotaRequest) error {
	err := service.anggotaRepository.Update(anggota)
	return err
}

func (service *anggotaService) DeleteAnggotaById(id uint) int {
	err := service.anggotaRepository.DeleteById(id)
	return err
}
