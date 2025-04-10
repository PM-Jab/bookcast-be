package databases

import (
	"bookcast/configs"
	"bookcast/pkg/utils"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewPostgreSQLDBConnection(cfg *configs.Configs) (*pgx.Conn, error) {
	postgresUrl, err := utils.ConnectionURLBuilder("postgresql", cfg)
	if err != nil {
		return nil, err
	}

	db, err := pgx.Connect(context.Background(), postgresUrl)
	if err != nil {
		defer db.Close(context.Background())
		log.Printf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	log.Println("postgreSQL database has been connected 🐘")
	return db, nil
}
