package services

import (
	"errors"
	"course-manager/models"
)

var alunos []models.Aluno
var nextAlunoID uint = 1

func CriarAluno(aluno models.Aluno) models.Aluno {
	aluno.ID = nextAlunoID
	nextAlunoID++
	alunos = append(alunos, aluno)
	return aluno
}

func ListarAlunos() []models.Aluno {
	return alunos
}

func BuscarAlunoPorID(id uint) (models.Aluno, error) {
	for _, a := range alunos {
		if a.ID == id {
			return a, nil
		}
	}
	return models.Aluno{}, errors.New("aluno não encontrado")
}

func AtualizarAluno(id uint, alunoAtualizado models.Aluno) (models.Aluno, error) {
	for i, a := range alunos {
		if a.ID == id {
			alunoAtualizado.ID = id
			alunos[i] = alunoAtualizado
			return alunoAtualizado, nil
		}
	}
	return models.Aluno{}, errors.New("aluno não encontrado")
}

func DeletarAluno(id uint) error {
	for i, a := range alunos {
		if a.ID == id {
			alunos = append(alunos[:i], alunos[i+1:]...)
			return nil
		}
	}
	return errors.New("aluno não encontrado")
}

func CursosDoAluno(id uint) ([]models.Matricula, error) {
	var cursos []models.Matricula
	for _, m := range matriculas {
		if m.AlunoID == id {
			cursos = append(cursos, m)
		}
	}
	return cursos, nil
} 