# Tasker API

Tasker — это REST API-сервис на Go, который предоставляет функциональность управления пользователями и заданиями.

---

## 🚀 Запуск проекта

### Создайте `.env`:

```bash
   cd tasker
   touch .env
```

Содержимое
```
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

### Запуск в docker-compose
```bash
docker compose up -d
```

## API

#### /api/v1/login
* `POST` : Логинизация
```json
Request: 
  {
    "login": "login",
    "password: "password"
  }
Response:
  {
    "token": "token"
  }
```

#### /api/v1/users/leaderboard
* `GET` : Получение лидерборда
```json
Response:
  {
    "users": []
  }
```

#### /api/v1/users/:id/task/complete
* `POST` : Выполение задания
```json
Request: 
  {
    "name": "name" // join_telegram, follow_twitter, invite_friend, fill_profile
  }
```

#### /api/v1/users/:id/referrer
* `POST` : Привязка одного пользователя к другом

#### /api/v1/users
* `POST` : Создание пользователя
```json
Request:
  {
      "login": "login",
      "password: "password"
  }
```

#### /api/v1/users/:id
* `PATCH` : Редактирование пользователя
```json
Request:
  {
      "name": "name"
  }
```

#### /api/v1/users/:id/status
* `GET` : Получение информации о пользователе
```json
Response:
  {
      "name": "name",
      "login": "login",
      "points": "points",
      "created_at": "created_at",
      "updated_at": "updated_at",
  }
```
