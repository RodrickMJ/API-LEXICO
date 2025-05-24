package main

import (
	"api_go/internal/application"
	"api_go/internal/infrastructure/env"
	myhttp "api_go/internal/infrastructure/http" 
	"api_go/internal/infrastructure/mysql"
	"database/sql"
	"log"
	"net/http" 

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := env.LoadDBEnv()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	repo := mysql.NewMySQLRepo(db)
	service := application.NewClienteService(repo)
	handler := myhttp.NewHandler(service) 

	mux := http.NewServeMux() 
	handler.RegisterRoutes(mux)

	log.Println("API corriendo en http://localhost:3000") 
	http.ListenAndServe(":3000", mux)
}