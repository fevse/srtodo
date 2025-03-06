
# TODO-list (REST API)

## Тестовое задание

### API-эндпоинты:

- POST /tasks – создание задачи;
- GET /tasks – получение списка всех задач;
- PUT /tasks/:id – обновление задачи;
- DELETE /tasks/:id – удаление задачи.

### Конфигурация базы данных и сервера с помощью .env

DBTYPE=

DBUSER=

DBPASS=

DBADDR=

DBNAME=


SERVHOST=

SERVPORT=

### Запуск через Makefile

make run -  запускает Docker контейнеры с приложением и базой данных (PostgreSQL) с помощью docker-compose
