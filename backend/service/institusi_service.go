package service

import (
	"backend/entity"
	"backend/repository"
	"backend/request"
)

type InstitusiService interface {
	FindInstitusiByID(id uint) (entity.Institusi, error)
	ListInstitusi() ([]entity.Institusi, error)
	InsertInstitusi(institusi request.InstitusiRequest) error
	UpdateInstitusi(institusi request.InstitusiRequest) error
	DeleteInstitusiById(id uint) int
}

type institusiService struct {
	institusiRepository repository.InstitusiRepository
}

func NewInstitusiService(institusiRepository repository.InstitusiRepository) *institusiService {
	return &institusiService{institusiRepository}
}

func (service *institusiService) FindInstitusiByID(id uint) (entity.Institusi, error) {
	institusi, err := service.institusiRepository.FindByID(id)
	return institusi, err
}

func (service *institusiService) ListInstitusi() ([]entity.Institusi, error) {
	institusi, err := service.institusiRepository.List()
	return institusi, err
}

func (service *institusiService) InsertInstitusi(institusi request.InstitusiRequest) error {
	err := service.institusiRepository.Insert(institusi)
	return err
}

func (service *institusiService) UpdateInstitusi(institusi request.InstitusiRequest) error {
	err := service.institusiRepository.Update(institusi)
	return err
}

func (service *institusiService) DeleteInstitusiById(id uint) int {
	err := service.institusiRepository.DeleteById(id)
	return err
}
