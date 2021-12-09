[<img src="../img/golang.svg" width="36"/>](Go)

# Processador de transações Go
Sistema desenvolvido para o processamento de transações desenvolvido com Golang e Apache Kafka

## Configurar /etc/hosts

A comunicação entre as aplicações se dá de forma direta através da rede da máquina.
Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu /etc/hosts (para Windows o caminho é C:\Windows\system32\drivers\etc\hosts):
```
127.0.0.1 host.docker.internal
```

## Rotina de execução e desenvolvimento da aplicação

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

# envio de um json com o contéudo da mensagem
```

```sh
# executando a aplicação Go
go run cmd/main.go
```