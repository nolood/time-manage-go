## Для установки переменной окружения и загрузки конфига Windows

```sh
$env:CONFIG_PATH="./config/local.yaml"; go run .\cmd\rest-api\main.go
```

## Linux 

```sh
export export CONFIG_PATH=./config/local.yaml
```
```sh
go run ./cmd/rest-api/main.go
```