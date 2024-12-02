# effective_mobile_task

Тестовое задание для Effective Mobile

## Запуск проекта

### Предварительные требования

- Go 1.23.3 или выше
- PostgreSQL для работы с базой данных

## Установка и запуск

### 1. Клонирование репозитория:

```sh
git clone git@github.com:Serzh757/effective_mobile_task.git
cd effective_mobile_task
```

### 2. Настройка конфигурации:

Создайте .env файл, используя [.env](.env) как шаблон. Укажите параметры базы данных и URL внешнего API.

### 3. Запуск миграций:

Миграции автоматически запустятся при поднятии проекта.

### 4. Запуск основного сервера на порту 8080:

```sh
docker-compose up --build
```

### 5. Swagger-документация: Генерация и доступ к документации:

```sh
make docs
```

Документация будет доступна по адресу http://localhost:8080/api/v1/documenation.

## Основные маршруты API

- **GET /api/v1/songs** - Получение списка песен с возможностью фильтрации и пагинации.
- **GET /api/v1/song/:id** - Получение детальной информации о песни по ИД.
- **PUT /api/v1/song** - Обновление информации о песне.
- **POST /api/v1/song** - Добавление новой песни.
- **DELETE /api/v1/song/:id** - Удаление песни по ID.

## Makefile Команды

- `make docs` устанавливает swag и генерирует swagger документацию из комментариев в исходном коде;
- `make linter` устанавливает `golangci-lint` и запускает проверку кода;
- `make fmt` устанавливает `goimports` и форматирует исходный код.
- `make check` Выполняет (ре)генерацию Swagger и выполняет все проверки.