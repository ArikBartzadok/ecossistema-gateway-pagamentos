package processo_transacao

import (
	"testing"
	"time"

	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/dominio/entidade"
	mock_repositorio "github.com/ArikBartzadok/ecossistema-gateway-pagamentos/dominio/repositorio/mock"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TesteProcessoTransacao_CartaoCreditoInvalido(t *testing.T) {
	entrada := EntradaTransacaoDTO{
		ID:                        "1",
		ContaID:                   "1",
		NumeroCartaoCredito:       "0000000000000000",
		NomeCartaoCredito:         "Diogo Ferreira",
		MesExpiracaoCartaoCredito: 12,
		AnoExpiracaoCartaoCredito: time.Now().Year(),
		CvvCartaoCredito:          123,
		Valor:                     200,
	}
	saida_esperada := SaidaTransacaoDTO{
		ID:           "1",
		Status:       entidade.REJEITADO,
		MensagemErro: "Número do cartão de crédito inválido",
	}

	controlador := gomock.NewController(t)
	defer controlador.Finish()

	repositorio_mock := mock_repositorio.NewMockRepositorioTransacao(controlador)
	repositorio_mock.EXPECT().
		Inserir(
			entrada.ID,
			entrada.ContaID,
			entrada.Valor,
			saida_esperada.Status,
			saida_esperada.MensagemErro,
		).
		Return(nil)

	caso_uso := NovoProcessoTransacao(repositorio_mock)
	saida, err := caso_uso.Executar(entrada)

	assert.Nil(t, err)
	assert.Equal(t, saida_esperada, saida)
}
