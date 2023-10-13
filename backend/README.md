# REST API using Gin and Gorm

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