# Tasker API

**Tasker** — это REST API-сервис на Go, предоставляющий функциональность управления пользователями, заданиями и реферальной системой.

---

## 🚀 Запуск проекта

### 1. Создайте `.env`:

```bash
cd tasker
touch .env
```

Содержимое файла `.env`:

```env
ADMIN_LOGIN=admin
ADMIN_PASS=admin
JWT_SECRET=tasker_secret
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=tasker
PG_DSN=host=postgres port=5432 dbname=tasker user=postgres password=postgres sslmode=disable
HTTP_HOST=0.0.0.0
HTTP_PORT=8080
```

### 2. Запуск через Docker Compose

```bash
docker compose up -d --build
```

API будет доступен по адресу: [http://localhost:8081](http://localhost:8081)

---

## 📋 API

> Все эндпоинты доступны с префиксом `/api/v1`

---

### 🔓 Публичные

#### `POST /api/v1/login` — Авторизация

Request
```json

{
  "login": "login",
  "password": "password"
}
```

Response:
```json
{
  "token": "token"
}
```

#### `GET /api/v1/users/leaderboard` — Получение лидерборда

Response:
```json
{
  "users": []
}
```

---

### 🔐 Приватные (JWT)

#### `POST /api/v1/users/:id/task/complete` — Выполнение задания

Request:
```json
{
  "name": "join_telegram"
}
```

#### `POST /api/v1/users/:id/referrer` — Установка реферера

Request:
```json
{
  "referrer_id": "123"
}
```

#### `POST /api/v1/users` — Создание пользователя

Request:
```json
{
  "login": "login",
  "password": "password"
}
```

#### `PATCH /api/v1/users/:id` — Обновление пользователя

Request:
```json
{
  "name": "new name"
}
```

#### `GET /api/v1/users/:id/status` — Инфо о пользователе

Response:
```json
{
  "name": "name",
  "login": "login",
  "points": 100,
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-02T15:30:00Z"
}
```

---

## 📅 Примечания

* JWT-токен передаётся через заголовок:

```
Authorization: Bearer <token>
```

* База данных инициализируется при первом запуске Docker.
* Для тестирования подойдёт Postman, Insomnia или `curl`.

---

## ✅ TODO

* [x] REST API для пользователей и заданий
* [x] JWT-авторизация
* [x] Docker + PostgreSQL
* [ ] Swagger документация
* [ ] Тесты
* [ ] CI/CD
* [ ] Ролевой доступ

---

## 📄 Лицензия

MIT License
