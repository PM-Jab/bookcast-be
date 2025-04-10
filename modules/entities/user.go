package entities

type UsersRegisterRes struct {
	Id       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

type UsersRegisterReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersRepositery interface {
	Register(req *UsersRegisterReq) (res *UsersRegisterRes, err error)
	GetUserByUsername(username string) (res *UsersByUsernameRes, err error)
}

type UsersService interface {
	Register(req *UsersRegisterReq) (res *UsersRegisterRes, err error)
	Login(req *UsersLoginReq) (res *UsersLoginRes, err error)
}

type UsersLoginReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersLoginRes struct {
	UserToken string `json:"usertoken" db:"usertoken"`
}

type UsersByUsernameRes struct {
	Id       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
