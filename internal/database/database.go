package database

import (
	"kura/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=kura_user password=kura_password dbname=kura_db port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco de dados:", err)
	}

	err = DB.AutoMigrate(&model.User{}, &model.Category{}, &model.Transaction{})
	if err != nil {
		log.Fatal("Falha ao rodar Migrations:", err)
	}

	log.Println("Conexão com o banco estabelecida e tabelas sincronizadas!")
}
