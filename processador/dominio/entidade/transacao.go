package entidade

import "errors"

const (
	REJEITADO = "rejeitado"
	APROVADO  = "aprovado"
)

type Transacao struct {
	ID            string
	ContaID       string
	Valor         float64
	CartaoCredito CartaoCredito
	Status        string
	MensagemErro  string
}

func NovaTransacao() *Transacao {
	return &Transacao{}
}

func (t *Transacao) Valida() error {
	if t.Valor > 1000 {
		return errors.New("Você não possui limite para essa transação")
	}

	if t.Valor < 1 {
		return errors.New("O valor da transação precisa ser maior que 1")
	}

	return nil
}

func (t *Transacao) adicionarCartaoCredito(cartao CartaoCredito) {
	t.CartaoCredito = cartao
}
