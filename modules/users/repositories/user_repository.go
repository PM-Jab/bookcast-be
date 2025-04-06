package repositories

import (
	"bookcast/modules/entities"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type usersRepo struct {
	Db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *usersRepo {
	return &usersRepo{
		Db: db,
	}
}

func (r *usersRepo) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	query := `
	INSERT INTO "users"(
		"username",
		"password"
	)
		VALUES ($1, $2)
		RETURNING "id", "username";
	`

	user := new(entities.UsersRegisterRes)

	rows, err := r.Db.Queryx(query, req.Username, req.Password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(user); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}

	defer r.Db.Close()

	return user, nil

}
