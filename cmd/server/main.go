package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alroymuhammad/finance-tracker/internal/handler"
	"github.com/alroymuhammad/finance-tracker/internal/storage"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main (){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("cant load env", err)
	}
	db := storage.NewDB()
	defer db.Close()
	h := &handler.Handler{DB:db}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /",h.HomeHandler)
	mux.HandleFunc("GET /db-check", h.DBCheckHandler)
	mux.HandleFunc("POST /echo", h.EchoHandler)
	fmt.Println("Running on server :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))

}