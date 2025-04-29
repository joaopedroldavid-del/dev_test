package services

import (
	"errors"
	"course-manager/models"
)

var cursos []models.Curso
var nextCursoID uint = 1

func CriarCurso(curso models.Curso) models.Curso {
	curso.ID = nextCursoID
	nextCursoID++
	cursos = append(cursos, curso)
	return curso
}

func ListarCursos() []models.Curso {
	return cursos
}

func BuscarCursoPorID(id uint) (models.Curso, error) {
	for _, c := range cursos {
		if c.ID == id {
			return c, nil
		}
	}
	return models.Curso{}, errors.New("curso não encontrado")
}

func AtualizarCurso(id uint, cursoAtualizado models.Curso) (models.Curso, error) {
	for i, c := range cursos {
		if c.ID == id {
			cursoAtualizado.ID = id
			cursos[i] = cursoAtualizado
			return cursoAtualizado, nil
		}
	}
	return models.Curso{}, errors.New("curso não encontrado")
}

func DeletarCurso(id uint) error {
	for i, c := range cursos {
		if c.ID == id {
			cursos = append(cursos[:i], cursos[i+1:]...)
			return nil
		}
	}
	return errors.New("curso não encontrado")
}

func MediaNotasCurso(cursoID uint) (float64, error) {
	soma := 0.0
	qtd := 0
	for _, m := range matriculas {
		if m.CursoID == cursoID {
			soma += m.Nota
			qtd++
		}
	}
	if qtd == 0 {
		return 0, nil
	}
	return soma / float64(qtd), nil
}