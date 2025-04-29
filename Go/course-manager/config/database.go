package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"course-manager/models"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar com banco de dados")
	}

	database.AutoMigrate(&models.Curso{}, &models.Aluno{}, &models.Matricula{})
	DB = database
}