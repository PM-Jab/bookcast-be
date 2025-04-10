package main

import (
	"bookcast/configs"
	"bookcast/modules/servers"
	"bookcast/pkg/databases"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// conn, err := pgx.Connect(context.Background(), "postgresql://postgres.pgtomljzmebxwfvxelvz:dbuITCU1vTGsmkIv@aws-0-us-east-1.pooler.supabase.com:5432/postgres?sslmode=require")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	// conn.QueryRow(context.Background(), "SELECT 1").Scan()

	// fmt.Println("Successfully connected to database")

	// fmt.Println(userId)

	// Load dotenv config
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Configs)

	// Fiber configs
	cfg.App.Host = os.Getenv("GIN_HOST")
	cfg.App.Port = os.Getenv("GIN_PORT")

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	db, err := databases.NewPostgreSQLDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close(context.Background())

	s := servers.NewServer(cfg, db)
	s.Start()
}
