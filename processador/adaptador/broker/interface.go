package broker

type InterfaceProdutor interface {
	Publicar(mensagem interface{}, chave []byte, topico string) error
}
