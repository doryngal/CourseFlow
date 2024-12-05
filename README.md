```
backend/
├── cmd/
│   ├── auth-service/
│   │   └── main.go                  # Точка входа для запуска Auth Service (аутентификация и уведомления)
│   ├── course-service/
│   │   └── main.go                  # Точка входа для запуска Course Service (курсы, кэширование, рейтинги)
│   └── user-service/
│       └── main.go                  # Точка входа для запуска User Service (профиль, записки)
├── config/                          # Конфигурационные файлы для всех сервисов
│   ├── auth-config.yml
│   ├── course-config.yml
│   └── user-config.yml
├── internal/
│   ├── auth-service/
│   │   ├── controllers/             # Контроллеры для аутентификации и уведомлений
│   │   │   ├── auth_controller.go
│   │   │   └── notification_controller.go
│   │   ├── middleware/              # JWT Middleware для аутентификации
│   │   ├── models/                  # Модели данных (User, Token)
│   │   ├── repository/              # Репозиторий для взаимодействия с БД (User, Tokens)
│   │   ├── routers/                 # Роуты для API
│   │   └── services/                # Логика для аутентификации и уведомлений (SMTP)
│   │       ├── auth_service.go
│   │       └── notification_service.go
│   ├── course-service/
│   │   ├── controllers/             # Контроллеры для курсов, модулей, комментариев и рейтингов
│   │   │   ├── course_controller.go
│   │   │   ├── module_controller.go
│   │   │   ├── comment_controller.go
│   │   │   └── rating_controller.go
│   │   ├── middleware/              # Middleware для кэширования
│   │   ├── models/                  # Модели данных (Course, Module, Comment, Rating)
│   │   ├── repository/              # Репозиторий для работы с БД (Course, Module, Comment, Rating)
│   │   ├── routers/                 # Роуты для API
│   │   └── services/                # Логика курсов, модулей, отзывов, рейтингов и кэширования (Redis)
│   │       ├── course_service.go
│   │       ├── module_service.go
│   │       ├── comment_service.go
│   │       ├── rating_service.go
│   │       └── cache_service.go     # Логика работы с кэшом (Redis)
│   ├── user-service/
│   │   ├── controllers/             # Контроллеры для профиля, записок, прогресса
│   │   │   ├── profile_controller.go
│   │   │   ├── enrollment_controller.go
│   │   │   └── progress_controller.go
│   │   ├── models/                  # Модели данных (UserProfile, Enrollment, Progress)
│   │   ├── repository/              # Репозиторий для работы с БД (Profile, Enrollment, Progress)
│   │   ├── routers/                 # Роуты для API
│   │   └── services/                # Логика профиля, записей на курсы, прогресса
│   │       ├── profile_service.go
│   │       ├── enrollment_service.go
│   │       └── progress_service.go
│   └── shared/
│       ├── utils/                   # Общие утилиты для всех микросервисов (например, JWT или работа с Redis)
│       └── middleware/              # Общие middleware (например, CORS)
├── migrations/                      # SQL-скрипты для миграций БД
│   ├── auth-service/
│   │   ├── 0001_create_users_table.up.sql
│   │   └── 0001_create_users_table.down.sql
│   ├── course-service/
│   │   ├── 0001_create_courses_table.up.sql
│   │   ├── 0002_create_modules_table.up.sql
│   │   ├── 0003_create_comments_table.up.sql
│   │   ├── 0004_create_ratings_table.up.sql
│   │   └── 0001_create_courses_table.down.sql
│   └── user-service/
│       ├── 0001_create_profiles_table.up.sql
│       ├── 0002_create_enrollments_table.up.sql
│       └── 0003_create_progress_table.up.sql
├── swagger/                         # Документация Swagger для каждого сервиса
│   ├── auth-service-swagger.yml
│   ├── course-service-swagger.yml
│   └── user-service-swagger.yml
└── Dockerfile                       # Основной Dockerfile
```