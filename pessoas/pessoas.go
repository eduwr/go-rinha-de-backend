package pessoas

import (
	"fmt"

	"github.com/eduwr/go-rinha-de-backend/rinhaguard"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Stack []string

type Pessoa struct {
	Id         string `json:"id" db:"id" validate:"required"`
	Apelido    string `json:"apelido" db:"apelido" validate:"required,max=32"`
	Nome       string `json:"nome" db:"nome" validate:"required,max=100"`
	Nascimento string `json:"nascimento" db:"nascimento" validate:"datetime=2006-01-02"`
}

type PessoaWithStack struct {
	Pessoa
	Stack []string `json:"stack" db:"stack"`
}

var PessoaSchema = `
	CREATE TABLE IF NOT EXISTS pessoas (
		id uuid PRIMARY KEY,
		apelido varchar(32) NOT NULL UNIQUE,
		nome varchar(100) NOT NULL,
		nascimento date
	);

	CREATE TABLE IF NOT EXISTS stacks (
		pessoa_id uuid REFERENCES pessoas(id),
		stack_value varchar(32),
		FOREIGN KEY (pessoa_id) REFERENCES pessoas(id)
	);
`

func (p *PessoaWithStack) Create(db *sqlx.DB) error {
	id := uuid.New()
	p.Id = id.String()

	err := rinhaguard.Check(p)
	if err != nil {
		return err
	}

	_, err = db.NamedExec(`
	INSERT INTO pessoas (id, apelido, nome, nascimento)
	VALUES (:id, :apelido, :nome, :nascimento)`, p)

	if err != nil {
		return err
	}

	for _, stackValue := range p.Stack {
		_, err = db.Exec(`
			INSERT INTO stacks (pessoa_id, stack_value)
			VALUES ($1, $2)`, p.Id, stackValue)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PessoaWithStack) Show(db *sqlx.DB) error {
	query := `
		SELECT id, apelido, nome, nascimento, stack_value as stack
		FROM pessoas p
		LEFT JOIN stacks s
		ON p.id = s.pessoa_id
		WHERE p.id = $1
	`

	rows, err := db.Queryx(query, p.Id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for rows.Next() {
		var (
			Id,
			Apelido,
			Nome,
			Nascimento,
			Stack string
		)
		rows.Scan(&Id, &Apelido, &Nome, &Nascimento, &Stack)

		fmt.Println(Id,
			Apelido,
			Nome,
			Nascimento,
			Stack)

		p.Apelido = Apelido
		p.Nome = Nome
		p.Nascimento = Nascimento
		p.Stack = append(p.Stack, Stack)
	}

	return nil
}
