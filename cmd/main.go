package main

import (
	"messenger/internal/database"
	"messenger/internal/ui"
	"log"
)

func main() {
	// Инициализация БД
	dbType := "postgres"
	dsn := "user=your_user password=your_password dbname=your_db sslmode=disable"
	db, err := database.InitDB(dbType, dsn)
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer db.Close()

	// Запуск пользовательского интерфейса
	ui.InitUI()
}
