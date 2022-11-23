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
- http://127.0.0.1:8080/store/:sort

Получить один магазин:
- http://127.0.0.1:8080/store/1/:uuid

Получить промокоды магазина:
- http://127.0.0.1:8080/store/promo/:quantity/:uuid

Получить список магазинов в корзине:
- http://127.0.0.1:8080/cart

Получить один магазин в корзине:
- http://127.0.0.1:8080/cart/1/:store

Увеличить кол-во промокодов магазина в корзине:
- http://127.0.0.1:8080/cart/increase/:store

Уменьшить кол-во промокодов магазина в корзине:
- http://127.0.0.1:8080/cart/decrease/:store

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

Удалить магазин из корзины:
- http://127.0.0.1:8080/cart/delete/:store





