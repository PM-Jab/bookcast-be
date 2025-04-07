package services

import (
	"bookcast/modules/entities"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type usersSer struct {
	UsersRepo entities.UsersRepositery
}

func NewUsersService(usersRepo entities.UsersRepositery) entities.UsersService {
	return &usersSer{
		UsersRepo: usersRepo,
	}
}

func (u *usersSer) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	user, err := u.UsersRepo.Register(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return user, nil
}
