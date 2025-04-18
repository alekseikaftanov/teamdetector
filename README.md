# Team Detector API

API для оценки и анализа командной культуры.

## Технологии

- Go 1.21
- Gin Web Framework
- PostgreSQL
- JWT аутентификация
- Docker & Docker Compose

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone https://github.com/alekseikaftanov/teamdetector.git
cd teamdetector
```

2. Создайте файл .env:
```bash
cp .env.example .env
```

3. Запустите приложение с помощью Docker Compose:
```bash
docker-compose up --build
```

## API Endpoints

### Аутентификация

#### Регистрация
```
POST /api/v1/auth/register
Body:
{
    "email": "user@example.com",
    "password": "password123",
    "name": "John Doe"
}
```

#### Авторизация
```
POST /api/v1/auth/login
Body:
{
    "email": "user@example.com",
    "password": "password123"
}
```

## Лицензия

MIT 