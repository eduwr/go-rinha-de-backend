package pessoas

import (
	"strings"

	"github.com/eduwr/go-rinha-de-backend/rinhaguard"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Stack []string

func NewStackFromString(str string) Stack {
	return Stack(strings.Split(str, ","))
}

type Pessoa struct {
	Id         string         `json:"id" db:"id" validate:"required"`
	Apelido    string         `json:"apelido" db:"apelido" validate:"required,max=32"`
	Nome       string         `json:"nome" db:"nome" validate:"required,max=100"`
	Nascimento string         `json:"nascimento" db:"nascimento" validate:"datetime=2006-01-02"`
	Stack      pq.StringArray `json:"stack" db:"stacks"`
}

func Create(p Pessoa, db *sqlx.DB) (*Pessoa, error) {
	p.Id = uuid.New().String()

	err := rinhaguard.Check(p)
	if err != nil {
		return nil, err
	}

	_, err = db.NamedExec(`
	INSERT INTO pessoas (id, apelido, nome, nascimento, stacks)
	VALUES (:id, :apelido, :nome, :nascimento, :stacks)`, p)

	if err != nil {
		return nil, err
	}

	addPessoaToCache(&p)

	return &p, nil
}

func Show(id string, db *sqlx.DB) (Pessoa, error) {
	if cachedPessoa := getPessoaFromCache(id); cachedPessoa != nil {
		return *cachedPessoa, nil
	}

	p := Pessoa{}
	validationErr := rinhaguard.CheckUUID(id)
	if validationErr != nil {
		return p, validationErr
	}
	query := `
		SELECT
			id,
			apelido,
			nome,
			TO_CHAR(nascimento, 'YYYY-MM-DD') AS nascimento,
			stacks
		FROM
			pessoas p
		WHERE
			p.id = $1
	`

	err := db.Get(&p, query, id)
	if err != nil {
		return p, err
	}

	return p, nil
}

func Index(t string, db *sqlx.DB) ([]Pessoa, error) {
	pessoas := []Pessoa{}

	query := `
		SELECT
			id,
			apelido,
			nome,
			TO_CHAR(nascimento, 'YYYY-MM-DD') AS nascimento,
			stacks
		FROM
			pessoas p
		WHERE
			(LOWER(p.nome) || ' ' || LOWER(p.apelido) || ' ' || LOWER(p.stacks)) LIKE '%' || LOWER($1) || '%'
		LIMIT 50
	`

	err := db.Select(&pessoas, query, t)

	if err != nil {
		return nil, err
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
