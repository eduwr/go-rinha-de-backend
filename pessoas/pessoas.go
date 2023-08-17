package pessoas

import (
	"github.com/eduwr/go-rinha-de-backend/rinhaguard"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Pessoa struct {
	Id         string   `json:"id" db:"id" validate:"required"`
	Apelido    string   `json:"apelido" db:"apelido" validate:"required,max=32"`
	Nome       string   `json:"nome" db:"nome" validate:"required,max=100"`
	Nascimento string   `json:"nascimento" db:"nascimento" validate:"datetime=2006-01-02"`
	Stack      []string `json:"stack" db:"stack"`
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

func (p *Pessoa) Create(db *sqlx.DB) (*Pessoa, error) {
	id := uuid.New()
	p.Id = id.String()

	err := rinhaguard.Check(p)
	if err != nil {
		return nil, err
	}

	_, err = db.NamedExec(`
	INSERT INTO pessoas (id, apelido, nome, nascimento)
	VALUES (:id, :apelido, :nome, :nascimento)`, p)

	if err != nil {
		return nil, err
	}

	for _, stackValue := range p.Stack {
		_, err = db.Exec(`
			INSERT INTO stacks (pessoa_id, stack_value)
			VALUES ($1, $2)`, p.Id, stackValue)

		if err != nil {
			return nil, err
		}
	}

	return p, nil
}
