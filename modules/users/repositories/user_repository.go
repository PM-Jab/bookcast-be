package repositories

import (
	"bookcast/modules/entities"
	"context"

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

	var id int8
	var username string
	r.Db.QueryRow(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, username",
		req.Username, req.Password).Scan(&id, &username)

	user := &entities.UsersRegisterRes{
		Username: username,
		Id:       id,
	}

	return user, nil
}
