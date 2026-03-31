package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)
func NewDB()  *sql.DB{
	host     := os.Getenv("DB_HOST")
	port     := os.Getenv("DB_PORT")
	user     := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname   := os.Getenv("DB_NAME")
	sslmode  := os.Getenv("DB_SSLMODE")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password,host, port, dbname, sslmode )

	var err error

	db, err := sql.Open("pgx", connStr)
	if err != nil{
		log.Fatal("Failed to setup connection: ", err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatal("Failed to ping database: ", err)
	}

	fmt.Println("Database terhubung!")
	return db
}