# REST API using Gin and Gorm

Установите goose для миграций  
```
$ go get github.com/pressly/goose/v3/cmd/goose@latest
$ goose -version
```
Как мигрировать
Создайте перед миграцией БД с параметрами ниже
```
$ export GOOSE_DRIVER=postgres
$ export GOOSE_DBSTRING="user=postgres password=12345 host=localhost dbname=gotest sslmode=disable"
$ cd migrations
$ goose up
```

How to use:

```
$ go mod tidy
$ go run .
```