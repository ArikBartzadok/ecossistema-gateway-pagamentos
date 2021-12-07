package entidade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TesteTransacaoValida(t *testing.T) {
	transacao := NovaTransacao()
	transacao.ID = "1"
	transacao.contaID = "1"
	transacao.valor = 900

	assert.Nil(t, transacao.Valida())
}

func TesteTransacaoMaiorQue1000Invalida(t *testing.T) {
	transacao := NovaTransacao()
	transacao.ID = "1"
	transacao.contaID = "1"
	transacao.valor = 1001

	erro := transacao.Valida()

	assert.Error(t, erro)
	assert.Equal(t, "Você não possui limite para essa transação", erro.Error())
}

func TesteTransacaoMenorQue1Invalida(t *testing.T) {
	transacao := NovaTransacao()
	transacao.ID = "1"
	transacao.contaID = "1"
	transacao.valor = 0

	erro := transacao.Valida()

	assert.Error(t, erro)
	assert.Equal(t, "O valor da transação precisa ser maior que 1", erro.Error())
}
