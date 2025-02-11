package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// Connect осуществляет подключение к существующей базе данных.
func Connect() {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "avito-task")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// AutoMigrate добавит недостающие столбцы, но не удалит существующие данные.
	err = db.AutoMigrate(&User{}, &Inventory{}, &CoinHistory{})
	if err != nil {
		log.Fatalf("Ошибка при миграции моделей: %v", err)
	}

	log.Println("Соединение с базой данных установлено, модели синхронизированы.")
}

// getEnv возвращает значение переменной окружения или значение по умолчанию, если переменная не задана.
func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
