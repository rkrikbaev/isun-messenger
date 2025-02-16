package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var DB *sql.DB

// InitDB инициализирует соединение с БД
func InitDB(dbType, dsn string) (*sql.DB, error) {
	DB, err := sql.Open(dbType, dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	if err = DB.Ping(); err != nil {
		DB.Close()
		return nil, fmt.Errorf("ошибка проверки соединения: %w", err)
	}

	log.Println("Подключение к БД успешно установлено")

	exists, err := tablesExist()
	if err != nil {
		DB.Close()
		return nil, fmt.Errorf("ошибка проверки существования таблиц: %w", err)
	}

	if !exists {
		if err := createTables(); err != nil {
			DB.Close()
			return nil, fmt.Errorf("ошибка создания таблиц: %w", err)
		}
	}

	return DB, nil
}

// tablesExist проверяет, существуют ли необходимые таблицы
func tablesExist() (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM information_schema.tables WHERE table_name = 'example'`
	err := DB.QueryRow(query).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// createTables создает таблицы, если их нет
func createTables() error {
	schema := `
	CREATE TABLE IF NOT EXISTS data (
		id SERIAL PRIMARY KEY,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		tag TEXT NOT NULL,
		value TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS data_history (
		id SERIAL PRIMARY KEY,
		data_id INTEGER NOT NULL,
		old_value TEXT NOT NULL,
		new_value TEXT NOT NULL,
		changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS settings (
		id SERIAL PRIMARY KEY,
		key TEXT UNIQUE NOT NULL,
		value TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		role TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS logs (
		id SERIAL PRIMARY KEY,
		message TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS license (
		id SERIAL PRIMARY KEY,
		license_key TEXT NOT NULL,
		verified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := DB.Exec(schema)
	if err != nil {
		return fmt.Errorf("ошибка создания таблиц: %w", err)
	}
	log.Println("Таблицы успешно созданы или уже существуют")
	return nil
}

// CloseDB закрывает соединение с БД
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Соединение с БД закрыто")
	}
}

// --- CRUD операции ---

// Добавление записи в data
func InsertData(tag, value string) (int64, error) {
	result, err := DB.Exec("INSERT INTO data (tag, value) VALUES ($1, $2)", tag, value)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// Чтение данных по тегу
func GetData(tag string) ([]map[string]string, error) {
	rows, err := DB.Query("SELECT id, timestamp, value FROM data WHERE tag = $1", tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []map[string]string
	for rows.Next() {
		var id int
		var timestamp time.Time
		var value string
		if err := rows.Scan(&id, &timestamp, &value); err != nil {
			return nil, err
		}
		result = append(result, map[string]string{
			"id":        fmt.Sprint(id),
			"timestamp": timestamp.String(),
			"value":     value,
		})
	}
	return result, nil
}

// Обновление данных
func UpdateData(id int, newValue string) error {
	var oldValue string
	err := DB.QueryRow("SELECT value FROM data WHERE id = $1", id).Scan(&oldValue)
	if err != nil {
		return err
	}

	_, err = DB.Exec("UPDATE data SET value = $1 WHERE id = $2", newValue, id)
	if err != nil {
		return err
	}

	_, err = DB.Exec("INSERT INTO data_history (data_id, old_value, new_value) VALUES ($1, $2, $3)", id, oldValue, newValue)
	return err
}

// Удаление данных
func DeleteData(id int) error {
	_, err := DB.Exec("DELETE FROM data WHERE id = $1", id)
	return err
}

// Добавление пользователя
func InsertUser(username, passwordHash, role string) error {
	_, err := DB.Exec("INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)", username, passwordHash, role)
	return err
}

// Получение пользователя
func GetUser(username string) (string, string, error) {
	var passwordHash, role string
	err := DB.QueryRow("SELECT password_hash, role FROM users WHERE username = $1", username).Scan(&passwordHash, &role)
	if err != nil {
		return "", "", err
	}
	return passwordHash, role, nil
}

// Обновление пароля пользователя
func UpdateUserPassword(username, newPasswordHash string) error {
	_, err := DB.Exec("UPDATE users SET password_hash = $1 WHERE username = $2", newPasswordHash, username)
	return err
}

// Удаление пользователя
func DeleteUser(username string) error {
	_, err := DB.Exec("DELETE FROM users WHERE username = $1", username)
	return err
}

// Установка настройки
func SetSetting(key, value string) error {
	_, err := DB.Exec("INSERT INTO settings (key, value) VALUES ($1, $2) ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value", key, value)
	return err
}

// Получение настройки
func GetSetting(key string) (string, error) {
	var value string
	err := DB.QueryRow("SELECT value FROM settings WHERE key = $1", key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

// Логирование событий
func InsertLog(message string) error {
	_, err := DB.Exec("INSERT INTO logs (message) VALUES ($1)", message)
	return err
}

// Проверка лицензии
func GetLicenseKey() (string, error) {
	var key string
	err := DB.QueryRow("SELECT license_key FROM license LIMIT 1").Scan(&key)
	if err != nil {
		return "", err
	}
	return key, nil
}

// Установка лицензии
func SetLicenseKey(key string) error {
	_, err := DB.Exec("INSERT INTO license (license_key) VALUES ($1) ON CONFLICT (id) DO UPDATE SET license_key = EXCLUDED.license_key", key)
	return err
}
