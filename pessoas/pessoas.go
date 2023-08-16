package pessoas

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Pessoa struct {
	Id         string   `json:"id" db:"id" validate:"required"`
	Apelido    string   `json:"apelido" db:"apelido" validate:"required"`
	Nome       string   `json:"nome" db:"nome" validate:"required,datetime=2006-01-02"`
	Nascimento string   `json:"nascimento" db:"nascimento"`
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

func (p *Pessoa) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		// Convert validation errors to a single error message
		var errMsg string
		for _, err := range err.(validator.ValidationErrors) {
			errMsg += err.Field() + " is invalid; "
		}
		return errors.New(errMsg)
	}

	return nil
}
