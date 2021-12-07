package entidade

import "errors"

type Transacao struct {
	ID            string
	contaID       string
	valor         float64
	cartaoCredito CartaoCredito
	status        string
	mensagemErro  string
}

func NovaTransacao() *Transacao {
	return &Transacao{}
}

func (t *Transacao) Valida() error {
	if t.valor > 1000 {
		return errors.New("Você não possui limite para essa transação")
	}

	if t.valor < 1 {
		return errors.New("O valor da transação precisa ser maior que 1")
	}

	return nil
}

func (t *Transacao) adicionarCartaoCredito(cartao CartaoCredito) {
	t.cartaoCredito = cartao
}
