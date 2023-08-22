CREATE TABLE IF NOT EXISTS pessoas (
    id uuid PRIMARY KEY,
    apelido varchar(32) NOT NULL UNIQUE,
    nome varchar(100) NOT NULL,
    nascimento date,
    stacks text
);


CREATE INDEX idx_pessoas_nome ON pessoas (LOWER(nome));
CREATE INDEX idx_pessoas_apelido ON pessoas (LOWER(apelido));
CREATE INDEX idx_pessoas_stacks ON pessoas (LOWER(stacks));