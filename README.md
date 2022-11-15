# web-2022

## Описание
Лабораторная работа №3 по курсу "Разработка интернет-приложений" 2022 года

## Запуск миграций
Установка goose:
- `go install github.com/pressly/goose/v3/cmd/goose@latest`

Создание миграций в папке migrations:
- `goose create db_freebie_shop sql`

Запуск миграций:
- `make migrate`

## Добавить недостающие модули
- `go mod tidy`

## Создание .env файла
- `make local`

## Запуск сервиса
- `make run`

## Создание бинарника
- `make build`

## Генерация swagger
- `make swagger`

## Запуск Docker
- `make docker`

## Запросы
### Swagger
Открыть swagger в браузере:
- http://127.0.0.1:8080/swagger/index.html

### GET
Получить список всех магазинов:
- http://127.0.0.1:8080/store

Получить цену магазина:
- http://127.0.0.1:8080/store/price/:uuid

Получить промокод магазина:
- http://127.0.0.1:8080/store/promo/:uuid

### POST
Добавить магазин:
- http://127.0.0.1:8080/store

Добавить рандомные магазины:
- http://127.0.0.1:8080/store/random

### PUT
Изменить цену магазина:
- http://127.0.0.1:8080/store/:uuid

### DELETE
Удалить магазин:
- http://127.0.0.1:8080/store/:uuid





