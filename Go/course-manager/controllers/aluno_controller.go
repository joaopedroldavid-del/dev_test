package controllers

import (
	"course-manager/config"
	"course-manager/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&aluno)
	c.JSON(http.StatusCreated, aluno)
}

func ListarAlunos(c *gin.Context) {
	var alunos []models.Aluno
	config.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func BuscarAlunoPorID(c *gin.Context) {
	id := c.Param("id")
	var aluno models.Aluno
	if err := config.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func AtualizarAluno(c *gin.Context) {
	id := c.Param("id")
	var aluno models.Aluno
	if err := config.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	var dadosAtualizados models.Aluno
	if err := c.ShouldBindJSON(&dadosAtualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aluno.Nome = dadosAtualizados.Nome
	config.DB.Save(&aluno)

	c.JSON(http.StatusOK, aluno)
}

func DeletarAluno(c *gin.Context) {
	id := c.Param("id")
	var aluno models.Aluno
	if err := config.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}
	config.DB.Delete(&aluno)
	c.JSON(http.StatusNoContent, nil)
}

func ListarCursosDoAluno(c *gin.Context) {
	id := c.Param("id")
	var matriculas []models.Matricula
	config.DB.Preload("Curso").Where("aluno_id = ?", id).Find(&matriculas)

	var cursos []models.Curso
	for _, m := range matriculas {
		cursos = append(cursos, m.Curso)
	}

	c.JSON(http.StatusOK, cursos)
}