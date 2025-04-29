package controllers

import (
	"course-manager/config"
	"course-manager/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CriarCurso(c *gin.Context) {
	var curso models.Curso
	if err := c.ShouldBindJSON(&curso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&curso)
	c.JSON(http.StatusCreated, curso)
}

func ListarCursos(c *gin.Context) {
	var cursos []models.Curso
	config.DB.Find(&cursos)
	c.JSON(http.StatusOK, cursos)
}

func BuscarCursoPorID(c *gin.Context) {
	id := c.Param("id")
	var curso models.Curso
	if err := config.DB.First(&curso, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Curso não encontrado"})
		return
	}
	c.JSON(http.StatusOK, curso)
}

func AtualizarCurso(c *gin.Context) {
	id := c.Param("id")
	var curso models.Curso
	if err := config.DB.First(&curso, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Curso não encontrado"})
		return
	}

	var dadosAtualizados models.Curso
	if err := c.ShouldBindJSON(&dadosAtualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curso.Nome = dadosAtualizados.Nome
	curso.Descricao = dadosAtualizados.Descricao
	config.DB.Save(&curso)

	c.JSON(http.StatusOK, curso)
}

func DeletarCurso(c *gin.Context) {
	id := c.Param("id")
	var curso models.Curso
	if err := config.DB.First(&curso, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Curso não encontrado"})
		return
	}
	config.DB.Delete(&curso)
	c.JSON(http.StatusNoContent, nil)
}

func CalcularNotaMedia(c *gin.Context) {
	id := c.Param("id")
	var curso models.Curso
	if err := config.DB.First(&curso, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Curso não encontrado"})
		return
	}

	var matriculas []models.Matricula
	config.DB.Where("curso_id = ?", id).Find(&matriculas)

	var soma float64
	var contador int

	for _, m := range matriculas {
		soma += m.Nota
		contador++
	}

	var media float64
	if contador > 0 {
		media = soma / float64(contador)
	}

	c.JSON(http.StatusOK, gin.H{
		"curso_id": id,
		"media":    media,
	})
}