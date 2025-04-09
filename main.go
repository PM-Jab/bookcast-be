package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	log.Println("Server running on port:", port)
	router.Run(":" + port)
}
