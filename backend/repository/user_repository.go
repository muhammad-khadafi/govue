package repository

/*
 * Created by muhammad.khadafi on 27/05/2022
 */

import (
	"backend/entity"
	"backend/request"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	FindByEmail(email string) (entity.User, error)
	FindByID(ID uint) (entity.User, error)
	FindByUsernameOrEmail(input string) (entity.User, error)
	Insert(user request.UserRequest) error
	FindByUsername(username string) (entity.User, error)
}

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{DB: db}
}

func (repository *userRepository) FindByEmail(email string) (entity.User, error) {
	user := entity.User{}
	script := `select * from "USER" where email = :1`
	err := repository.DB.Get(&user, script, email)
	user.Password = ""
	return user, err
}

func (repository *userRepository) FindByID(ID uint) (entity.User, error) {
	user := entity.User{}
	script := `select * from "USER" where id = :1`
	err := repository.DB.Get(&user, script, ID)
	user.Password = ""
	return user, err
}

func (repository *userRepository) FindByUsernameOrEmail(input string) (entity.User, error) {
	user := entity.User{}
	script := `select * from "USER" where username = :1 or email = :2`
	err := repository.DB.Get(&user, script, input, input)
	user.Password = ""
	return user, err
}

func (repository *userRepository) Insert(user request.UserRequest) error {
	script := `INSERT INTO "USER" u (username, password, email) VALUES (:USERNAME, :PASSWORD, :EMAIL)`
	_, err := repository.DB.NamedExec(script, user)
	return err
}

func (repository *userRepository) FindByUsername(username string) (entity.User, error) {
	user := entity.User{}
	script := `select * from "USER" where username = :1`
	err := repository.DB.Get(&user, script, username)
	user.Password = ""
	return user, err
}
