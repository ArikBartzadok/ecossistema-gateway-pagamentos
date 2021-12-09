package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/adaptador/apresentacao/transacao"
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/adaptador/broker/kafka"
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/adaptador/fabrica"
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/modelo/processo_transacao"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("[INFO] Iniciando processador de pagamentos GO...")

	banco, err := sql.Open("sqlite3", "teste.db")
	if err != nil {
		log.Fatal(err)
	}

	fabricaRepositorioBanco := fabrica.NovaFabricaRepositorioBanco(banco)
	repositorio := fabricaRepositorioBanco.CriarRepositorioTransacao()

	// distribuindo mensagens
	produtorConfigMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	apresentadorKafka := transacao.NovaTransacaoApresentacaoKafka()
	produtorKafka := kafka.NovoProdutorKafka(produtorConfigMap, apresentadorKafka)

	// consumindo mensagens
	var mensagemChan = make(chan *ckafka.Message)
	consumidorConfigMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}
	topicos := []string{"transacoes"}
	consumidor := kafka.NovoConsumidor(consumidorConfigMap, topicos)

	// processando as mensagens em uma nova thread, sem travar a aplicação no broker/kafka/consumidor
	go consumidor.Consumir(mensagemChan)

	topico_transacoes := "resultado_transacoes"
	modelo := processo_transacao.NovoProcessoTransacao(repositorio, produtorKafka, topico_transacoes)

	for mensagem := range mensagemChan {
		var entrada processo_transacao.EntradaTransacaoDTO
		json.Unmarshal(mensagem.Value, &entrada)

		modelo.Executar(entrada)
	}

}
