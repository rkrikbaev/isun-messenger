├── cmd
│   ├── main.go  // Точка входа
│
├── config
│   ├── config.go  // Загрузка и парсинг конфигурации
│
├── internal
│   ├── database
│   │   ├── db.go       // Подключение к БД
│   │   ├── migrations.go // Управление миграциями
│   │   ├── crud.go     // Операции CRUD
│   │   ├── migrations
│   │   │   ├── init.sql // Начальные миграции
│   │
│   ├── opcua
│   │   ├── client.go  // Подключение и опрос OPC UA
│   │   ├── handler.go // Обработка данных
│   │
│   ├── ui
│   │   ├── app.go     // Инициализация UI (Fyne)
│   │   ├── handlers.go // Обработчики UI
|   │   ├── layout.go    // Основной макет интерфейса
|   │   ├── components.go // Кастомные UI-элементы    
│   │
│   ├── soap
│   │   ├── client.go  // Клиент для отправки SOAP-запросов
│   │
│   ├── websocket
│   │   ├── server.go  // WebSocket-сервер
│   │
│   ├── auth
│   │   ├── auth.go    // Аутентификация и авторизация (RBAC)
│   │
│   ├── logging
│   │   ├── logger.go  // Логирование в файлы
│   │
│   ├── notifications
│   │   ├── email.go   // Отправка email-уведомлений
│   │
│   ├── licensing
│   │   ├── license.go // Проверка лицензии
│
├── pkg
│   ├── utils
│   │   ├── helpers.go // Утилиты и вспомогательные функции
│
├── scripts
│   ├── run_migrations.sh // Скрипт для запуска миграций
│
├── Dockerfile  // Сборка Docker-образа
├── docker-compose.yml  // Запуск с PostgreSQL
├── go.mod  // Go-модули
├── go.sum  // Зависимости
├── README.md  // Документация
