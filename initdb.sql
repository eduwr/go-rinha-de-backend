CREATE TABLE IF NOT EXISTS pessoas (
    id uuid PRIMARY KEY,
    apelido varchar(32) NOT NULL UNIQUE,
    nome varchar(100) NOT NULL,
    nascimento date,
    stacks text
);
