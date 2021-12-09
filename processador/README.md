```sh
# subindo containers
docker-compose up -d
```

```sh
# executando container
docker exec -it processador_go bash
```

```sh
# iniciando uma aplicação go
go mod init github.com/ArikBartzadok/ecossistema-gateway-pagamentos
```

```sh
# executando testes da aplicação
go test ./...
# ou
go test -count=1 ./...

# limpando cache dos testes
go clean -testcache
```

```sh
# gerando mocks para testes com o repositório
mockgen -destination=dominio/repositorio/mock/mock.go -source=dominio\/repositorio\/repositorio.go
```

```sh
# executando o kafka em outro container, para gerar mensagens
docker exec -t processador_kafka_1 bash

# conectando-se ao tópico
kafka-console-producer --bootstrap-server=localhost:9092 --topic=transacoes
```