package repository

import (
	"fmt"
	"log"
	"todo/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func (d *DB) GetDB() *gorm.DB {
	return d.db
}

func (d *DB) PostgresConnect() {
	env, err := godotenv.Read("../../.env")
	if err != nil {
		log.Println("❌ Ошибка загрузки .env:", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		env["DB_HOST"],
		env["DB_USER"],
		env["DB_PASSWORD"],
		env["DB_NAME"],
		env["DB_PORT"],
	)

	log.Println("DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Не удалось подключиться к базе данных:", err)
	}

	d.db = db

	// Миграции
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("❌ Ошибка миграции базы данных:", err)
	}

	log.Println("✅ Успешно подключено к базе данных и выполнены миграции")
}
