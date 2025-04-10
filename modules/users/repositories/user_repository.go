package repositories

import (
	"bookcast/modules/entities"
	"context"
	"fmt"

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

	// runnning number of users
	amount, err := r.GetUserLength()
	if err != nil {
		return nil, err
	}

	var id string
	var username string
	err = r.Db.QueryRow(context.Background(), "INSERT INTO users (id, username, password) VALUES ($1, $2, $3) RETURNING id, username",
		fmt.Sprint(100+amount+1), req.Username, req.Password).Scan(&id, &username)

	if err != nil {
		fmt.Println("error inserting user")
		return nil, err
	}

	user := &entities.UsersRegisterRes{
		Id:       id,
		Username: username,
	}

	return user, nil
}

func (r *usersRepo) GetUserByUsername(username string) (*entities.UsersByUsernameRes, error) {
	var id string
	var password string
	err := r.Db.QueryRow(context.Background(), "SELECT id, password FROM users WHERE username = $1",
		username).Scan(&id, &password)

	if err != nil {
		fmt.Println("error getting user by username")
		return nil, err
	}
	user := &entities.UsersByUsernameRes{
		Id:       id,
		Username: username,
		Password: password,
	}

	return user, nil
}

func (r *usersRepo) GetUserLength() (int8, error) {
	var count int8
	err := r.Db.QueryRow(context.Background(), "SELECT COUNT(*) FROM users").Scan(&count)

	if err != nil {
		fmt.Println("error getting user length")
		return 0, err
	}

	return count, nil
}
