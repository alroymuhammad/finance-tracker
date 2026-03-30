package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)
var db *sql.DB
func initDB(){
	host     := os.Getenv("DB_HOST")
	port     := os.Getenv("DB_PORT")
	user     := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname   := os.Getenv("DB_NAME")
	sslmode  := os.Getenv("DB_SSLMODE")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password,host, port, dbname, sslmode )

	var err error

	db, err = sql.Open("pgx", connStr)
	if err != nil{
		log.Fatal("Failed to setup connection: ", err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatal("Failed to ping database: ", err)
	}

	fmt.Println("Database terhubung!")
}

type HellowWorldResponse struct{
	Message string `json:"message"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	response := HellowWorldResponse{
		Message: "Hello world",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func DBCheckHandler(w http.ResponseWriter, r *http.Request){
	err := db.Ping()
	if err != nil {
		http.Error(w, "Database tidak terhubung", http.StatusInternalServerError)
		return
	}
	response := HellowWorldResponse{
		Message: "DB connected",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func EchoHandler(w http.ResponseWriter, r *http.Request){
	var body map[string]any

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil{
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	fmt.Println("Body input", body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

func main (){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("cant load env", err)
	}
	initDB()
	defer db.Close()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", HomeHandler)
	mux.HandleFunc("GET /db-check", DBCheckHandler)
	mux.HandleFunc("POST /echo", EchoHandler)
	fmt.Println("Running on server :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))

}