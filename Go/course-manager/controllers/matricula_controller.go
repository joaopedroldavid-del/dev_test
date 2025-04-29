package controllers

import (
	"course-manager/config"
	"course-manager/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CriarMatricula(c *gin.Context) {
	var matricula models.Matricula
	if err := c.ShouldBindJSON(&matricula); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&matricula)
	c.JSON(http.StatusCreated, matricula)
}

func ListarMatriculas(c *gin.Context) {
	var matriculas []models.Matricula
	config.DB.Preload("Aluno").Preload("Curso").Find(&matriculas)
	c.JSON(http.StatusOK, matriculas)
}

func BuscarMatriculaPorID(c *gin.Context) {
	id := c.Param("id")
	var matricula models.Matricula
	if err := config.DB.Preload("Aluno").Preload("Curso").First(&matricula, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Matrícula não encontrada"})
		return
	}
	c.JSON(http.StatusOK, matricula)
}

func AtualizarNota(c *gin.Context) {
	id := c.Param("id")
	var matricula models.Matricula
	if err := config.DB.First(&matricula, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Matrícula não encontrada"})
		return
	}

	var dadosAtualizados struct {
		Nota float64 `json:"nota"`
	}
	if err := c.ShouldBindJSON(&dadosAtualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	matricula.Nota = dadosAtualizados.Nota
	config.DB.Save(&matricula)

	c.JSON(http.StatusOK, matricula)
}

func DeletarMatricula(c *gin.Context) {
	id := c.Param("id")
	var matricula models.Matricula
	if err := config.DB.First(&matricula, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Matrícula não encontrada"})
		return
	}
	config.DB.Delete(&matricula)
	c.JSON(http.StatusNoContent, nil)
}