CREATE TABLE IF NOT EXISTS users (
  ID   BIGSERIAL PRIMARY KEY,
  nome text      NOT NULL,
  email  text,
  cpf  text
);
