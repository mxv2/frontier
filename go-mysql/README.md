# MySQL abort connection

1. Вопроизведение ошибок
```
$ docker-compose up -d
$ docker-compose logs --follow
$ go run main.go
```
В логах должны быть 2 записи Abort connection

2. Воспроизведение без ошибок
```
# раскомментировать replace в go.mod
$ GOPRIVATE=github.com/mxv2 go mod tidy && go mod vendor
$ go run main.go
```
В логах не должно появиться новых записей Abort connection
