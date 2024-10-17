```
backend/
├── cmd/
│   ├── auth-service/
│   │   └── main.go              # Точка входа для запуска Auth Service (аутентификация и уведомления)
│   └── course-service/
│       └── main.go              # Точка входа для запуска Course Service (курсы, отзывы, кэш)
├── config/                      # Конфигурационные файлы для всех сервисов
│   ├── auth-config.yml
│   └── course-config.yml
├── internal/
│   ├── auth-service/
│   │   ├── controllers/         # Контроллеры для аутентификации и уведомлений
│   │   │   ├── auth_controller.go
│   │   │   └── notification_controller.go
│   │   ├── middleware/          # JWT Middleware для аутентификации
│   │   ├── models/              # Модели данных (User, Token и т.д.)
│   │   ├── repository/          # Репозиторий для взаимодействия с БД (User, Tokens)
│   │   ├── routers/             # Роуты для API
│   │   └── services/            # Логика для аутентификации и уведомлений (SMTP)
│   │       ├── auth_service.go
│   │       └── notification_service.go
│   ├── course-service/
│   │   ├── controllers/         # Контроллеры для курсов, отзывов
│   │   │   ├── course_controller.go
│   │   │   └── review_controller.go
│   │   ├── middleware/          # Middleware для кэширования
│   │   ├── models/              # Модели данных (Course, Review)
│   │   ├── repository/          # Репозиторий для работы с БД (Course, Review)
│   │   ├── routers/             # Роуты для API
│   │   └── services/            # Логика курсов, отзывов и кэширования (Redis)
│   │       ├── course_service.go
│   │       ├── review_service.go
│   │       └── cache_service.go # Логика работы с кэшом (Redis)
│   └── shared/
│       ├── utils/               # Общие утилиты для всех микросервисов (например, JWT или работа с Redis)
│       └── middleware/          # Общие middleware (например, CORS)
├── migrations/                  # SQL-скрипты для миграций БД
│   ├── auth-service/
│   │   ├── 0001_create_users_table.up.sql
│   │   └── 0001_create_users_table.down.sql
│   └── course-service/
│       ├── 0001_create_courses_table.up.sql
│       └── 0001_create_courses_table.down.sql
├── swagger/                     # Документация Swagger для каждого сервиса
│   ├── auth-service-swagger.yml
│   └── course-service-swagger.yml
└── Dockerfile                   # Основной Dockerfile
```