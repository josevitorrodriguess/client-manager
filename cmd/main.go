package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // Biblioteca para carregar variáveis de ambiente de um arquivo .env
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load()

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error to connect db:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Could not ping DB:", err)
	}
	fmt.Println("Connection with DB created successfully!")

}
