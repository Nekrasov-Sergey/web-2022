# web-2022

## Описание
Лабораторная работа №3 по курсу "Разработка интернет-приложений" 2022 года

## Запуск миграций
В консоли PostgreSQL написать команду:
- `CREATE EXTENSION "uuid-ossp";`

В консоли Goland написать команду:
- `go run .\cmd\migrate\main.go`

## Запуск сервиса через терминал

### Вариант №1
- `go run .\cmd\freebie-shop\main.go`

### Вариант №2
- `go build .\cmd\freebie-shop\main.go`

- `\main.exe`

## Сгенерировать swagger
- `swag init -g \cmd\freebie-shop\main.go`
## Эндпоинт
Получить список всех промокодов:
- http://127.0.0.1:8080/promos/get

Добавить промокод:
- http://127.0.0.1:8080/promos/create

Создать 5 рандомных промокодов:
- http://127.0.0.1:8080/promos/create/random






