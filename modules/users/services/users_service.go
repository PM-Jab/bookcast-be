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

	// check if username already exists
	user, _ := u.UsersRepo.GetUserByUsername(req.Username)
	if user != nil {
		fmt.Println("user already exists")
		return nil, fmt.Errorf("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	regisUser, err := u.UsersRepo.Register(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return regisUser, nil
}

func (u *usersSer) Login(req *entities.UsersLoginReq) (*entities.UsersLoginRes, error) {
	user, err := u.UsersRepo.GetUserByUsername(req.Username)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("wrong password")
	}

	userLogin := &entities.UsersLoginRes{
		UserToken: "testtoken",
	}

	return userLogin, nil

}
