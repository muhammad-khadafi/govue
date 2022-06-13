package service

import (
	"backend/entity"
	"backend/repository"
	"backend/request"
)

/*
 * Created by muhammad.khadafi on 27/05/2022
 */

type UserService interface {
	FindUserByID(id uint) (entity.User, error)
	FindUserByUsernameOrEmail(input string) (entity.User, error)
	FindUserByEmail(email string) (entity.User, error)
	InsertUser(user request.UserRequest) error
	FindUserByUsername(username string) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) FindUserByID(ID uint) (entity.User, error) {
	user, err := s.userRepository.FindByID(ID)
	return user, err
}

func (s *userService) FindUserByUsernameOrEmail(input string) (entity.User, error) {
	user, err := s.userRepository.FindByUsernameOrEmail(input)
	return user, err
}

func (s *userService) FindUserByEmail(email string) (entity.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	return user, err
}

func (s *userService) InsertUser(user request.UserRequest) error {
	err := s.userRepository.Insert(user)
	return err
}

func (s *userService) FindUserByUsername(username string) (entity.User, error) {
	user, err := s.userRepository.FindByUsername(username)
	return user, err
}
