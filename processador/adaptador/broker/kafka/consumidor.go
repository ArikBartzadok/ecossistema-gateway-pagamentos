package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Consumidor struct {
	ConfigMap *ckafka.ConfigMap
	Topicos   []string
}

func NovoConsumidor(configMap *ckafka.ConfigMap, topicos []string) *Consumidor {
	return &Consumidor{
		ConfigMap: configMap,
		Topicos:   topicos,
	}
}

func (c *Consumidor) Consumir(mensagemChan chan *ckafka.Message) error {
	// chan -> ajuda a receber uma mensagem e a processála, ao passo que recebe e processa outras
	consumidor, err := ckafka.NewConsumer(c.ConfigMap)

	if err != nil {
		return err
	}

	err = consumidor.SubscribeTopics(c.Topicos, nil)

	if err != nil {
		return err
	}

	for {
		mensagem, err := consumidor.ReadMessage(-1)

		if err == nil {
			mensagemChan <- mensagem
		}
	}

}
