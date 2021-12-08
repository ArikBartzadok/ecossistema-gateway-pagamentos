CREATE TABLE transacoes
(
    id            TEXT NOT NULL,
    conta_id      TEXT NOT NULL,
    valor         REAL NOT NULL,
    status        TEXT NOT NULL,
    mensagem_erro TEXT NOT NULL,
    criada_em     TEXT NOT NULL,
    atualizada_em TEXT NOT NULL
);