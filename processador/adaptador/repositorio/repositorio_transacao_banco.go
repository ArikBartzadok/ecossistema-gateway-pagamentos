package repositorio

import (
	"database/sql"
	"time"
)

type RepositorioTransacaoBanco struct {
	db *sql.DB
}

func NovoRepositorioTransacaoBanco(db *sql.DB) *RepositorioTransacaoBanco {
	return &RepositorioTransacaoBanco{
		db: db,
	}
}

func (r *RepositorioTransacaoBanco) Inserir(id string, conta string, valor float64, status string, mensagemErro string) error {
	stmt, err := r.db.Prepare(`
		insert into transacoes (id, conta, valor, status, mensagem_erro, criada_em, atualizada_em)
		values ($1, $2, $3, $4, $5, $6, $7)
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		id,
		conta,
		valor,
		status,
		mensagemErro,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return nil
	}

	return nil
}
