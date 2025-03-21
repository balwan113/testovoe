REST API для управления пользователями на Go
Этот проект представляет собой REST API для управления пользователями. Он написан на языке Go с использованием PostgreSQL в качестве базы данных, Docker для контейнеризации, а также включает миграции базы данных, модульные тесты и поддержку окружения через .env.

🛠️ Стек технологий
Язык программирования: Go
База данных: PostgreSQL
Контейнеризация: Docker, Docker Compose
Фреймворк: Gin
Миграции: golang-migrate
Тестирование:
Встроенный пакет testing Go
Библиотека testify для утверждений и моков
testcontainers-go для интеграционных тестов с PostgreSQL
🚀 Запуск проекта
1. Клонирование репозитория
Для начала клонируйте репозиторий:

bash
Копировать
Редактировать
git clone https://github.com/yourusername/your-repo-name.git
cd your-repo-name
2. Настройка окружения
Создайте файл .env на основе .env.example и заполните его своими данными:

bash
Копировать
Редактировать
cp .env.example .env
3. Запуск Docker
Для запуска проекта используйте Docker Compose. Это поднимет сервер на Go и PostgreSQL.

bash
Копировать
Редактировать
docker-compose up --build
После запуска проект будет доступен по адресу:

arduino
Копировать
Редактировать
http://localhost:8081
📚 Методы API
1. Создание пользователя
Метод: POST /users

Тело запроса:

json
Копировать
Редактировать
{
  "name": "Иван",
  "email": "ivan@example.com"
}
2. Получение информации о пользователе
Метод: GET /users/{id}

Ответ:

json
Копировать
Редактировать
{
  "id": 1,
  "name": "Иван",
  "email": "ivan@example.com"
}
3. Обновление данных пользователя
Метод: PUT /users/{id}

Тело запроса:

json
Копировать
Редактировать
{
  "name": "Иван Иванов",
  "email": "ivan.ivanov@example.com"
}
🧪 Тестирование
Для запуска модульных тестов выполните команду:

bash
Копировать
Редактировать
go test ./...
Интеграционные тесты
Для запуска интеграционных тестов с использованием Docker и PostgreSQL выполните команду:

bash
Копировать
Редактировать
go test -tags=integration ./...
🐳 Docker
Проект полностью контейнеризирован с помощью Docker. Для запуска достаточно выполнить:

bash
Копировать
Редактировать
docker-compose up --build
Это поднимет:

Сервер на Go (порт 8080)
PostgreSQL (порт 5432)
📁 Структура проекта
Проект имеет следующую структуру:

csharp
Копировать
Редактировать
.
├── cmd
│   └── main.go              # Точка входа
├── db
│   └── migrations           # Миграции базы данных
│       ├── 000001_create_users_table.up.sql
│       └── 000001_create_users_table.down.sql
├── internal
│   ├── config               # Конфигурация приложения
│   │   └── config.go
│   ├── database             # Подключение к базе данных
│   │   └── db.go
│   ├── handler              # Обработчики HTTP-запросов
│   │   ├── user_handler.go
│   │   └── user_handler_test.go
│   ├── models               # Модели данных
│   │   └── user.go
│   ├── repository           # Логика работы с базой данных
│   │   ├── user_repository.go
│   │   └── user_repository_test.go
│   ├── router               # Маршрутизация
│   │   └── router.go
│   └── service              # Бизнес-логика
│       ├── user_service.go
│       └── user_service_test.go
├── .env.example             # Пример файла окружения
├── docker-compose.yml       # Docker Compose конфигурация
├── Dockerfile               # Dockerfile для сборки образа
├── go.mod                   # Модули Go
└── README.md                # Документация
