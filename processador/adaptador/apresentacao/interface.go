package apresentacao

type Apresentacao interface {
	Exibir() ([]byte, error)
	Vincular(interface{}) error
}
