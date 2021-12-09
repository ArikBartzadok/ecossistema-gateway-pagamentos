package kafka

import (
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/adaptador/apresentacao"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Produtor struct {
	ConfigMap    *ckafka.ConfigMap
	Apresentacao apresentacao.Apresentacao
}

func NovoProdutorKafka(configMap *ckafka.ConfigMap, apresentacao apresentacao.Apresentacao) *Produtor {
	return &Produtor{
		ConfigMap:    configMap,
		Apresentacao: apresentacao,
	}
}

func (p *Produtor) Publicar(mensagem interface{}, chave []byte, topico string) error {
	produtor, err := ckafka.NewProducer(p.ConfigMap)

	if err != nil {
		return err
	}

	err = p.Apresentacao.Vincular(mensagem)

	if err != nil {
		return err
	}

	mensagemApresentacao, err := p.Apresentacao.Exibir()

	if err != nil {
		return err
	}

	mensagemPublicar := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic:     &topico,
			Partition: ckafka.PartitionAny,
		},
		Value: mensagemApresentacao,
		Key:   chave,
	}

	err = produtor.Produce(mensagemPublicar, nil) // a garantia de entrega está sendo enviada como nil, e por enquanto não haverá um delivery channel para garantia de entrega

	if err != nil {
		return err
	}

	return nil
}
