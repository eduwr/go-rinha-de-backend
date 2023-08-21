package pessoas

import (
	"errors"
	"strings"

	"github.com/eduwr/go-rinha-de-backend/rinhaguard"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Stack []string

func NewStackFromString(str string) Stack {
	return Stack(strings.Split(str, ","))
}

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

func Create(p PessoaWithStack, db *sqlx.DB) (*PessoaWithStack, error) {
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

	return &p, nil
}

func Show(id string, db *sqlx.DB) (*PessoaWithStack, error) {
	query := `
		SELECT
			id,
			apelido,
			nome,
			TO_CHAR(nascimento, 'YYYY-MM-DD') AS nascimento,
			stack_value AS stack
		FROM
			pessoas p
		LEFT JOIN stacks s
		ON p.id = s.pessoa_id
		WHERE
			p.id = $1
	`
	p := new(PessoaWithStack)

	rows, err := db.Queryx(query, id)
	if err != nil {
		return nil, err
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

		p.Id = Id
		p.Apelido = Apelido
		p.Nome = Nome
		p.Nascimento = Nascimento
		p.Stack = append(p.Stack, Stack)
	}

	err = rinhaguard.Check(p)
	if err != nil {
		return nil, errors.New("not found")
	}

	return p, nil
}

func Index(t string, db *sqlx.DB) ([]PessoaWithStack, error) {
	pessoas := []PessoaWithStack{}

	query := `
		SELECT
			id,
			apelido,
			nome,
			TO_CHAR(nascimento, 'YYYY-MM-DD') AS nascimento,
			string_agg(s.stack_value, ',') AS stack
		FROM
			pessoas p
		LEFT JOIN stacks s ON
			p.id = s.pessoa_id
		WHERE
			p.nome ILIKE $1 OR
			p.apelido ILIKE $1 OR
			EXISTS (
		        SELECT 1
		        FROM stacks s
		        WHERE s.pessoa_id = p.id AND s.stack_value ILIKE $1
		    )
		GROUP BY
			p.id,
			p.apelido,
			p.nome,
			p.nascimento
		LIMIT 50
	`

	rows, err := db.Queryx(query, "%"+t+"%")

	if err != nil {
		return nil, err
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

		s := NewStackFromString(Stack)
		p := Pessoa{Id: Id, Apelido: Apelido, Nome: Nome, Nascimento: Nascimento}

		pessoas = append(pessoas, PessoaWithStack{
			Pessoa: p,
			Stack:  s,
		})
	}

	return pessoas, nil
}

func Count(db *sqlx.DB) int {
	query := `
		SELECT count(*) from pessoas;
	`
	var count int
	db.Get(&count, query)
	return count
}
