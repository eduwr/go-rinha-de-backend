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