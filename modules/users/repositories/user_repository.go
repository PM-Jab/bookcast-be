package repositories

import (
	"bookcast/modules/entities"

	"github.com/jackc/pgx/v5"
)

type usersRepo struct {
	Db *pgx.Conn
}

func NewUsersRepository(db *pgx.Conn) *usersRepo {
	return &usersRepo{
		Db: db,
	}
}

func (r *usersRepo) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	// query := `
	// INSERT INTO "users"(
	// 	"username",
	// 	"password"
	// )
	// 	VALUES ($1, $2)
	// 	RETURNING "id", "username";
	// `

	// user := new(entities.UsersRegisterRes)

	return nil, nil

}
