package entities

type UsersRegisterRes struct {
	Id       uint64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

type UsersRegisterReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersRepositery interface {
	Register(req *UsersRegisterReq) (res *UsersRegisterRes, err error)
}

type UsersService interface {
	Register(req *UsersRegisterReq) (res *UsersRegisterRes, err error)
}
