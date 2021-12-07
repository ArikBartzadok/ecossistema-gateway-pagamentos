```sh
docker-compose up -d

docker exec -it processador_go bash

go mod init github.com/ArikBartzadok/ecossistema-gateway-pagamentos

go test ./...
```