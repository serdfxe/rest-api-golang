# rest-api-golang

Скелет Rest API на golang с jwt.

# Установка

## Docker контейнер

```
docker pull postgres:14-alpine
```

```
docker run --name=db -e POSTGRES_PASSWORD=314159 -p 5432:5432 -d --rm postgres:14-alpine
```

## Миграция БД

```
migrate -path ./schema -database "postgres://postgres:314159@localhost:5432/postgres?sslmode=disable" up
```

Создать файл .env в корне

```
DB_PASSWORD=314159
```

## Запуск

```
go run cmd/main.go
```
