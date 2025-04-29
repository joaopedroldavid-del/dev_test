package services

import (
	"errors"
	"course-manager/models"
)

var matriculas []models.Matricula
var nextMatriculaID uint = 1

func CriarMatricula(matricula models.Matricula) models.Matricula {
	matricula.ID = nextMatriculaID
	nextMatriculaID++
	matriculas = append(matriculas, matricula)
	return matricula
}

func ListarMatriculas() []models.Matricula {
	return matriculas
}

func BuscarMatriculaPorID(id uint) (models.Matricula, error) {
	for _, m := range matriculas {
		if m.ID == id {
			return m, nil
		}
	}
	return models.Matricula{}, errors.New("matricula não encontrada")
}

func AtualizarNota(id uint, novaNota float64) (models.Matricula, error) {
	for i, m := range matriculas {
		if m.ID == id {
			matriculas[i].Nota = novaNota
			return matriculas[i], nil
		}
	}
	return models.Matricula{}, errors.New("matricula não encontrada")
}

func DeletarMatricula(id uint) error {
	for i, m := range matriculas {
		if m.ID == id {
			matriculas = append(matriculas[:i], matriculas[i+1:]...)
			return nil
		}
	}
	return errors.New("matricula não encontrada")
}