package routes

import (
	"course-manager/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/cursos", controllers.CriarCurso)
	router.GET("/cursos", controllers.ListarCursos)
	router.GET("/cursos/:id", controllers.BuscarCursoPorID)
	router.PUT("/cursos/:id", controllers.AtualizarCurso)
	router.DELETE("/cursos/:id", controllers.DeletarCurso)
	router.GET("/cursos/:id/media", controllers.CalcularNotaMedia)

	router.POST("/alunos", controllers.CriarAluno)
	router.GET("/alunos", controllers.ListarAlunos)
	router.GET("/alunos/:id", controllers.BuscarAlunoPorID)
	router.PUT("/alunos/:id", controllers.AtualizarAluno)
	router.DELETE("/alunos/:id", controllers.DeletarAluno)
	router.GET("/alunos/:id/cursos", controllers.ListarCursosDoAluno)

	router.POST("/matriculas", controllers.CriarMatricula)
	router.GET("/matriculas", controllers.ListarMatriculas)
	router.GET("/matriculas/:id", controllers.BuscarMatriculaPorID)
	router.PUT("/matriculas/:id", controllers.AtualizarNota)
	router.DELETE("/matriculas/:id", controllers.DeletarMatricula)
}
