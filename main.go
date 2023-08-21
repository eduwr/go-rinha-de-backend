package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eduwr/go-rinha-de-backend/app"
	"github.com/eduwr/go-rinha-de-backend/dbconfig"
	"github.com/joho/godotenv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("HELLO RINHA DE BACKEND")
	production := os.Getenv("GO_ENVIRONMENT") == "production"

	if !production {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db, err := sqlx.Connect(dbconfig.NewDBConfig("postgres").GetConnString())
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	a := app.NewApp("Go! Rinha de Backend")
	a.Setup(db)
	a.Serve(":3333")
}
