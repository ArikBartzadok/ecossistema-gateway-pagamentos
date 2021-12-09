package kafka

import (
	"testing"

	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/adaptador/apresentacao/transacao"
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/dominio/entidade"
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/modelo/processo_transacao"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

func TesteProdutor_Publicar(t *testing.T) {
	saida_esperada := processo_transacao.SaidaTransacaoDTO{
		ID:           "1",
		Status:       entidade.REJEITADO,
		MensagemErro: "Você não possui limite para essa transação",
	}
	// saida_json, _ := json.Marshal(saida_esperada)

	config_map := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	produtor := NovoProdutorKafka(&config_map, transacao.NovaTransacaoApresentacaoKafka())

	err := produtor.Publicar(
		saida_esperada,
		[]byte("1"), // chave do kafka
		"teste",     // tópico do kafka
	)

	assert.Nil(t, err)
}
