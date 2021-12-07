package entidade

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TesteNumeroCartaoCredito(t *testing.T) {
	// inválido
	_, erro := NovoCartaoCredito("00000000", "Diogo Ferreira", 12, 2025, 123)
	assert.Equal(t, "Número do cartão de crédito inválido", erro.Error())

	// válido
	_, erro = NovoCartaoCredito("5555341244441115", "Diogo Ferreira", 12, 2025, 123)
	assert.Nil(t, erro)
}

func TesteMesExpiracaoCartaoCredito(t *testing.T) {
	// inválido
	_, erro := NovoCartaoCredito("5555341244441115", "Diogo Ferreira", 13, 2025, 123)
	assert.Equal(t, "Mês de expiração do cartão de crédito inválido", erro.Error())

	_, erro = NovoCartaoCredito("5555341244441115", "Diogo Ferreira", 0, 2025, 123)
	assert.Equal(t, "Mês de expiração do cartão de crédito inválido", erro.Error())

	// válido
	_, erro = NovoCartaoCredito("5555341244441115", "Diogo Ferreira", 1, 2025, 123)
	assert.Nil(t, erro)
}

func TesteAnoExpiracaoCartaoCredito(t *testing.T) {
	anoAnterior := time.Now().AddDate(-1, 0, 0)

	_, erro := NovoCartaoCredito("5555341244441115", "Diogo Ferreira", 12, anoAnterior.Year(), 123)
	assert.Equal(t, "Ano de expiração do cartão de crédito inválido", erro.Error())

	_, erro = NovoCartaoCredito("5555341244441115", "Diogo Ferreira", 12, 2025, 123)
	assert.Nil(t, erro)
}
