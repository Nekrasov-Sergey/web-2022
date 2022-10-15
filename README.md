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
### GET
Получить список всех промокодов:
- http://127.0.0.1:8080/promos/get

### POST
Добавить промокод:
- http://127.0.0.1:8080/promos/create?Store=&Discount=&Price=&Quantity=&Promo=

Добавить рандомные промокоды:
- http://127.0.0.1:8080/promos/create/random?Quantity=

### PUT
Изменить цену промокода:
- http://127.0.0.1:8080/promos/change/price?UUID=&Price=

### DELETE
Удалить промокод:
- http://127.0.0.1:8080/promos/delete?UUID=





