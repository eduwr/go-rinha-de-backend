package main

import (
	"fmt"
	"log"

	"github.com/eduwr/go-rinha-de-backend/app"
	"github.com/eduwr/go-rinha-de-backend/pessoas"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("HELLO RINHA DE BACKEND")

	db, err := sqlx.Connect("postgres", "host=localhost port=5432 dbname=rinha user=user password=pass sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(pessoas.PessoaSchema)

	a := app.NewApp("Go! Rinha de Backend")
	a.Setup()
	a.Serve(":3333")
}
